package order

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"starfish/infra/middleware"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoOrder := newRepoOrder()
	service := newService(repoOrder, db)
	controller := newController(service)

	// protected route
	orderRouter := router.Group("/merchants")
	orderRouter.Use((middleware.JWTMiddleware()))
	orderRouter.GET("/orders", controller.findAllByMerchantID)
}
