package campaign

import (
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repo := newPostgres()
	validate := validator.New()
	service := newService(repo, db)
	controller := newController(service, validate)

	campaignRouter := router.Group("/merchants")
	campaignRouter.Use(middleware.JWTMiddleware())
	campaignRouter.POST("/campaigns", controller.createCampaign)
}
