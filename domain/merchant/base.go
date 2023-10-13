package merchant

import (
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoMerchant := NewRepoMerchant()

	service := newService(repoMerchant, db)
	controller := newController(service)

	merchantRouter := router.Group("/merchants")
	merchantRouter.Use(middleware.JWTMiddleware())
	merchantRouter.GET("/profile", controller.getMerchantProfileById)
}
