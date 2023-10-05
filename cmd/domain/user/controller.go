package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	svc UserService
}

func NewController(svc UserService) userController {
	return userController{
		svc: svc,
	}
}

func (u userController) SignUp(c *gin.Context) {
	// Parse the registration request JSON from the request body
	var request Register

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":    false,
			"error_code": 40004,
			"message":    "invalid request"})
		return
	}

	user, err := NewUser().FromRegisterToUser(request)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":    false,
				"error_code": 50003,
				"message":    "internal error"})
			return
		}
		c.JSON(customErr.StatusCode, gin.H{
			"success":    false,
			"error_code": customErr.ErrorCode,
			"message":    customErr.Message})
		return

	}

	var merchant Merchant

	merchant, err = merchant.FromRegisterToMerchant(request)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":    false,
				"error_code": 50003,
				"message":    "internal error"})
			return
		}
		c.JSON(customErr.StatusCode, gin.H{
			"success":    false,
			"error_code": customErr.ErrorCode,
			"message":    customErr.Message})
		return

	}

	// Call the user registration service
	if err := u.svc.Register(context.Background(), user, merchant); err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":    false,
				"error_code": 50003,
				"message":    "internal error"})
			return
		}
		c.JSON(customErr.StatusCode, gin.H{
			"success":    false,
			"error_code": customErr.ErrorCode,
			"message":    customErr.Message})
		return
	}

	// Return a success response with the user ID
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "create success",
	})
}
