package merchant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type merchantController struct {
	svc MerchantService
}

func newController(svc MerchantService) merchantController {
	return merchantController{
		svc: svc,
	}
}

func (u merchantController) getMerchantProfileById(c *gin.Context) {
	userClaims := c.MustGet("user").(jwt.MapClaims)
	userId := userClaims["user_id"].(string)

	data, err := u.svc.GetMerchantProfileById(c, userId)
	if err != nil {
		customErr, ok := err.(*CustomError)
		if !ok {
			writeError(c, customErr, 50003, http.StatusInternalServerError)
			return
		}
		writeError(c, customErr, customErr.ErrorCode, customErr.StatusCode)
		return
	}
	writeSuccess(c, data, http.StatusOK)

}
