package main

import (
	"log"
	"os"
	"starfish/config"
	"starfish/domain/auth"
	"starfish/domain/campaign"
	"starfish/domain/file"
	"starfish/domain/merchant"
	"starfish/domain/order"
	"starfish/domain/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	// CORS
	server.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Length", "Content-Type"},
		AllowAllOrigins: true,
	}))

	err := godotenv.Load(".env")
	if err != nil {
		// panic(err)
		log.Println("no env provided")
	}

	AppConfig := config.AppConfig{Port: os.Getenv("APP_PORT")}
	DBConfig := config.DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}

	db, err := config.NewDB(DBConfig)
	if err != nil {
		panic(err)
	}

	api := server.Group("/api/v1")

	// product routes
	product.Run(api, db)

	// merchant routes
	merchant.Run(api, db)

	// user routes
	auth.Run(api, db)

	file.Run(api, db)

	// campaign routes
	campaign.Run(api, db)

	// order routes
	order.Run(api, db)

	server.Run(":" + AppConfig.Port)
}
