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
	var request Register

	if err := c.ShouldBindJSON(&request); err != nil {
		writeError(c, "invalid request", 40004, 400)
		return
	}

	user, err := NewUser().FromRegisterToUser(request)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	var merchant Merchant

	merchant, err = merchant.FromRegisterToMerchant(request)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	if err := u.svc.Register(context.Background(), user, merchant); err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "create success",
	})
}

func (u userController) SignIn(c *gin.Context) {
	var request Login

	if err := c.ShouldBindJSON(&request); err != nil {
		writeError(c, "invalid request", 40004, 400)
		return
	}

	userLogin, err := NewUser().FromLogin(request)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	data, err := u.svc.login(context.Background(), userLogin)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return
	}
	writeSuccess(c, data, http.StatusOK)
}
