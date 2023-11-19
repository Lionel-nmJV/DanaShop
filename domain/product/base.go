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
	productRouter := router.Group("/products")
	productRouter.Use((middleware.JWTMiddleware()))
	productRouter.GET("/", controller.findAllByMerchantID)
	productRouter.POST("/", controller.addProduct)
	productRouter.GET("/:productID", controller.findByID)
	productRouter.PUT("/:productID", controller.UpdateProduct)
	productRouter.DELETE("/:productID", controller.DeleteProduct)
}
