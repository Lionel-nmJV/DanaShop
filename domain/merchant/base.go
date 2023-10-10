package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	repoMerchant := NewRepoMerchant()

	service := newService(repoMerchant, db)
	controller := newController(service)

	router.GET("/merchants/profile", controller.getMerchantProfileById)
}
