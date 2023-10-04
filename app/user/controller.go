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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := NewUser().FromRegisterToUser(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var merchant Merchant

	merchant, err = merchant.FromRegisterToMerchant(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the user registration service
	if err := u.svc.Register(context.Background(), user, merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response with the user ID
	c.JSON(http.StatusCreated, gin.H{"messages": "success"})
}
