package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	svc UserService
}

func newController(svc UserService) userController {
	return userController{
		svc: svc,
	}
}

func (u userController) register(c *gin.Context) {
	var request register

	if err := c.ShouldBindJSON(&request); err != nil {
		writeError(c, "invalid request", 40004, 400)
		return
	}

	user, err := newUser().fromRegisterToUser(request)
	if err != nil {
		customErr, ok := err.(*customError)
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
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	if err := u.svc.Register(context.Background(), user, merchant); err != nil {
		customErr, ok := err.(*customError)
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

func (u userController) login(c *gin.Context) {
	var request login

	if err := c.ShouldBindJSON(&request); err != nil {
		writeError(c, "invalid request", 40004, 400)
		return
	}

	userLogin, err := newUser().FromLogin(request)
	if err != nil {
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	data, err := u.svc.login(context.Background(), userLogin)
	if err != nil {
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return
	}
	writeSuccess(c, data, http.StatusOK)
}
