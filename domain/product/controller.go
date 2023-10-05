package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
	service productService
}

func newController(service productService) productController {
	return productController{service: service}
}

func (controller productController) findAllByMerchantID(ctx *gin.Context) {

	products, err := controller.service.findAllByMerchantID(ctx)
	if err != nil {
		writeError(ctx, err, 40401, http.StatusNotFound)
		return
	}

	writeSuccess(ctx, products, http.StatusOK)
}
