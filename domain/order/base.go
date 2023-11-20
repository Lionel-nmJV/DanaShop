package order

import (
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoOrder := newRepoOrder()
	service := newService(repoOrder, db)
	controller := newController(service)

	// protected route
	orderRouter := router.Group("/orders")
	orderRouter.Use((middleware.JWTMiddleware()))
	orderRouter.GET("", controller.findAllByMerchantID)
}
