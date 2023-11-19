package merchant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type merchantController struct {
	svc merchantService
}

func newController(svc merchantService) merchantController {
	return merchantController{
		svc: svc,
	}
}

func (u merchantController) getMerchantProfileById(c *gin.Context) {
	userClaims := c.MustGet("user").(jwt.MapClaims)
	userId := userClaims["user_id"].(string)

	data, err := u.svc.GetMerchantProfileById(c, userId)
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

func (u merchantController) getMerchantAnalytics(c *gin.Context) {
	merchantID := "merchant_id_example" // Ganti aja ID merchant
	analytics, err := u.svc.GetMerchantAnalytics(c, merchantID)
	if err != nil {
		writeError(c, err, 50001, http.StatusInternalServerError) // ganti aja cuma contoh
		return
	}

	writeSuccess(c, analytics, http.StatusOK)
}
