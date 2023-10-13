package product

import (
	"starfish/domain/merchant"
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoProduct := newRepoProduct()
	repoMerchant := merchant.NewRepoMerchant()

	validate := validator.New()
	service := newService(repoProduct, repoMerchant, db)
	controller := newController(service, validate)

	router.Use(middleware.JWTMiddleware())
	router.GET("/merchants/products", controller.findAllByMerchantID)
	router.POST("/merchants/products", controller.addProduct)
	router.PUT("/merchants/products/:id", controller.UpdateProduct)
	router.DELETE("/merchants/products/:id", controller.DeleteProduct)
}
