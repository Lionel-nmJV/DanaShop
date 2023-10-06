package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type productController struct {
	service  productService
	validate *validator.Validate
}

func newController(service productService, validate *validator.Validate) productController {
	return productController{service: service, validate: validate}
}

func (c productController) findAllByMerchantID(ctx *gin.Context) {

	products, err := c.service.findAllByMerchantID(ctx)
	if err != nil {
		writeError(ctx, err, 40401, http.StatusNotFound)
		return
	}

	writeSuccess(ctx, products, http.StatusOK)
}

func (c productController) addProduct(ctx *gin.Context) {
	var request createRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	product, err := NewProduct().formAddProduct(request, c.validate)
	if err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	err = c.service.addProduct(ctx, product)
	if err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "create success",
	})
}
