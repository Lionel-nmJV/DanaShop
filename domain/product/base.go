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

	// protected route
	productRouter := router.Group("/merchants")
	productRouter.Use((middleware.JWTMiddleware()))
	productRouter.GET("/products", controller.findAllByMerchantID)
	productRouter.POST("/products", controller.addProduct)
}
