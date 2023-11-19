package merchant

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type merchantController struct {
	svc      merchantService
	validate *validator.Validate
}

func newController(svc merchantService, validate *validator.Validate) merchantController {
	return merchantController{
		svc:      svc,
		validate: validate,
	}
}

func (ctl merchantController) getMerchantProfileById(c *gin.Context) {
	userClaims := c.MustGet("user").(jwt.MapClaims)
	userId := userClaims["user_id"].(string)

	data, err := ctl.svc.getMerchantProfileById(c, userId)
	if err != nil {
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, customErr, 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr, customErr.ErrorCode, customErr.StatusCode)
		return
	}
	writeSuccess(c, data, http.StatusOK)

}

func (ctl merchantController) updateMerchantProfileById(c *gin.Context) {
	userClaims := c.MustGet("user").(jwt.MapClaims)
	userId := userClaims["user_id"].(string)

	var req updateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, errors.New("invalid request"), 40001, 400)
		return
	}

	merchant, err := newMerchant().fromRequest(req, ctl.validate)
	if err != nil {
		fmt.Println(err)
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, errors.New("invalid request"), 40001, 400)
			return
		}
		writeError(c, customErr, customErr.ErrorCode, customErr.StatusCode)
		return

	}

	fmt.Println("merchant ctl after valid :", merchant)

	ok, err := ctl.svc.updateMerchantProfileById(c, userId, merchant)
	if err != nil {
		customErr, ok := err.(*customError)
		if !ok {
			writeError(c, errors.New("internal error"), 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr, customErr.ErrorCode, customErr.StatusCode)
		return
	}

	if ok {
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"messages": "merchants profile updated",
		})
	}

}
