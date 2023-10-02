package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"starfish/config"
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

	_, err = config.NewDB(DBConfig)
	if err != nil {
		panic(err)
	}

	server.Run(":" + AppConfig.Port)
}
