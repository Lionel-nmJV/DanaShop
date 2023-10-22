package main

import (
	"log"
	"os"
	"starfish/config"
	"starfish/domain/image"
	"starfish/domain/merchant"
	"starfish/domain/product"
	"starfish/domain/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	// CORS
	server.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE","OPTION"},
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

	// Set up user routes
	user.Run(api, db)

	image.Run(api, db)

	server.Run(":" + AppConfig.Port)
}
