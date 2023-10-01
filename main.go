package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"starfish/config"
)

func main() {
	server := gin.Default()
	config.NewDB()

	server.Run(":3000")
}
