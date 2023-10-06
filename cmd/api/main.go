package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"starfish/config"
	"starfish/domain/product"
	"starfish/domain/user"
)

func main() {
	server := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
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

	api := server.Group("/api")

	// product routes
	product.Run(api, db)

	// Set up user routes
	user.SetupRoutes(server, db)

	server.Run(":" + AppConfig.Port)
}
