package campaign

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type campaignController struct {
	svc      CampaignService
	validate *validator.Validate
}

func newController(svc CampaignService, validate *validator.Validate) campaignController {
	return campaignController{
		svc:      svc,
		validate: validate,
	}
}

func (ctl campaignController) createCampaign(ctx *gin.Context) {
	var request createRequest

	userClaims := ctx.MustGet("user").(jwt.MapClaims)
	merchantId := userClaims["merchant_id"].(string)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		writeError(ctx, "invalid request", 40002, 400)
		return
	}

	campaign, err := newCampaign().fromRequest(request, ctl.validate)
	if err != nil {
		writeError(ctx, "invalid request", 40001, http.StatusBadRequest)
		return
	}

	campaign.MerchantId = uuid.MustParse(merchantId)

	if err := ctl.svc.createCampaign(context.Background(), campaign); err != nil {
		customErr, ok := err.(*customError)
		if !ok {
			writeError(ctx, "internal error", 50003, http.StatusInternalServerError)
			return
		}
		writeError(ctx, customErr.Message, customErr.ErrorCode, customErr.StatusCode)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "create campaign success",
	})

}