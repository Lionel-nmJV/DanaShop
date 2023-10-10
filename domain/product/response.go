package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseError struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type responseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func writeSuccess(ctx *gin.Context, data interface{}, statusCode int) {
	resp := responseSuccess{
		Success: true,
		Data:    data,
	}

	ctx.JSON(statusCode, resp)
}

func writeError(ctx *gin.Context, message error, errorCode int, statusCode int) {
	resp := responseError{
		Success:   false,
		ErrorCode: errorCode,
		Message:   message.Error(),
	}

	ctx.JSON(statusCode, resp)
}

func (controller productController) updateProduct(ctx *gin.Context) {
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

func (controller productController) deleteProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	if err := controller.service.deleteProduct(ctx, productID); err != nil {
		writeError(ctx, err, 50002, http.StatusInternalServerError)
		return
	}

	writeSuccess(ctx, "Product deleted successfully", http.StatusOK)
}

func writeError2(ctx *gin.Context, err error, errorCode int, statusCode int) {
	response := responseError{
		Success:   false,
		ErrorCode: errorCode,
		Message:   err.Error(),
	}
	ctx.JSON(statusCode, response)
}

func writeSuccess2(ctx *gin.Context, data interface{}, statusCode int) {
	response := responseSuccess{
		Success: true,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}
