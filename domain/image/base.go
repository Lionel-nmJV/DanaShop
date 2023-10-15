package image

import (
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"starfish/config"
	"starfish/domain/merchant"
	"starfish/infra/middleware"

	"github.com/gin-gonic/gin"
)

func Run(router *gin.RouterGroup, db *sqlx.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		// panic(err)
		log.Println("no env provided")
	}

	cloudinaryConfig := config.CloudinaryConfig{
		Name:      os.Getenv("CLOUDY_NAME"),
		APIKey:    os.Getenv("CLOUDY_API_KEY"),
		APISecret: os.Getenv("CLOUDY_API_SECRET"),
	}

	CLOUDINARY_URL := fmt.Sprintf("cloudinary://%s:%s@%s", cloudinaryConfig.APIKey, cloudinaryConfig.APISecret, cloudinaryConfig.Name)

	cloudinaryService, err := cloudinary.NewFromURL(CLOUDINARY_URL)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	validate := validator.New()
	repoMerchant := merchant.NewRepoMerchant()
	service := newService(cloudinaryService, repoMerchant, db)
	controller := newController(service, validate)

	// protected route
	imageRouter := router.Group("/files")
	imageRouter.Use((middleware.JWTMiddleware()))
	imageRouter.POST("/upload", controller.uploadFile)
}
