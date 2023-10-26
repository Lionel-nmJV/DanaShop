package product

import (
	"errors"
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
		writeError(ctx, errors.New("invalid request"), 40001, http.StatusBadRequest)
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

func (c productController) findByID(ctx *gin.Context) {
	productID := ctx.Param("productID")

	product, err := c.service.findByID(ctx, productID)
	if err != nil {
		writeError(ctx, errors.New("not found"), 40401, http.StatusNotFound)
		return
	}

	writeSuccess(ctx, product, http.StatusOK)
}

func (c productController) UpdateProduct(ctx *gin.Context) {
	var request updateRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		writeError(ctx, errors.New("invalid request"), 40002, http.StatusBadRequest)
		return
	}

	product, err := NewProduct().formUpdateProduct(request, c.validate)
	if err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	// You can get the product ID from the request or URL parameters.
	productID := ctx.Param("productID")

	// Add logic to update the product using the service's updateProduct method.
	err = c.service.updateProduct(ctx, productID, product)
	if err != nil {
		writeError(ctx, err, 40401, http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "update success",
	})
}

func (c productController) DeleteProduct(ctx *gin.Context) {
	// You can get the product ID from the request or URL parameters.
	productID := ctx.Param("productID")

	// Add logic to delete the product using the service's deleteProduct method.
	err := c.service.deleteProduct(ctx, productID)
	if err != nil {
		writeError(ctx, err, 40401, http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "delete success",
	})
}
