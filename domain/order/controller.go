package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type orderController struct {
	service orderService
}

func newController(service orderService) orderController {
	return orderController{service: service}
}

func (c orderController) findAllByMerchantID(ctx *gin.Context) {
	orders, err := c.service.findAllByMerchantID(ctx)
	if err != nil {
		writeError(ctx, err, 40401, http.StatusNotFound)
		return
	}

	writeSuccess(ctx, orders, http.StatusOK)
}
