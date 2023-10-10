package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {

	repo := newPostgres()
	svc := newService(repo, db)
	ctl := newController(svc)

	router.POST("/signup", ctl.register)
	router.POST("/signin", ctl.login)
}
