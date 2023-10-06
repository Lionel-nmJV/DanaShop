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

func (controller productController) updateProduct1(ctx *gin.Context) {
	var product Product
	if err := ctx.BindJSON(&product); err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	if err := controller.service.updateProduct(ctx, &product); err != nil {
		writeError(ctx, err, 50001, http.StatusInternalServerError)
		return
	}

	writeSuccess(ctx, "Product updated successfully", http.StatusOK)
}

func (controller productController) deleteProduct1(ctx *gin.Context) {
	productID := ctx.Param("id")

	if err := controller.service.deleteProduct(ctx, productID); err != nil {
		writeError(ctx, err, 50002, http.StatusInternalServerError)
		return
	}

	writeSuccess(ctx, "Product deleted successfully", http.StatusOK)
}

func writeError1(ctx *gin.Context, err error, errorCode int, statusCode int) {
	// Implement error response handling here
	// You can format the error message and status code as needed
	ctx.JSON(statusCode, gin.H{"error": err.Error(), "code": errorCode})
}

func writeSuccess1(ctx *gin.Context, data interface{}, statusCode int) {
	// Implement success response handling here
	// You can format the response data and status code as needed
	ctx.JSON(statusCode, data)
}
