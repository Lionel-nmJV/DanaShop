package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(router *gin.Engine, db *sqlx.DB) {

	repo := NewPostgres()
	svc := NewService(repo, db)
	ctl := NewController(svc)

	router.POST("/signup", ctl.SignUp)
}
