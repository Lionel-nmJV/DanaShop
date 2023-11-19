package merchant

import (
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoMerchant := NewRepoMerchant()
	validate := validator.New()

	repoOrder := NewOrderRepository()
	repoAnalytics := NewAnalyticsRepository()

	service := newService(repoMerchant, repoOrder, db, repoAnalytics)
	controller := newController(service, validate)

	merchantRouter := router.Group("/merchants")
	merchantRouter.Use(middleware.JWTMiddleware())
	merchantRouter.GET("/profile", controller.getMerchantProfileById)
	merchantRouter.PUT("/profile", controller.updateMerchantProfileById)
	merchantRouter.GET("/analytics", controller.getMerchantAnalytics)
}
