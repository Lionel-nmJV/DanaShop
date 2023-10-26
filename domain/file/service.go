package file

import (
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"path/filepath"
	"starfish/domain/merchant"
	"strings"
)

type MerchantRepository interface {
	ReadMerchantRepo
}

type ReadMerchantRepo interface {
	FindByUserID(ctx *gin.Context, tx *sqlx.Tx, merchantID string) (merchant.Merchant, error)
}

type fileService struct {
	cloudinaryService *cloudinary.Cloudinary
	RepoMerchant      MerchantRepository
	db                *sqlx.DB
}

func newService(cloudinary *cloudinary.Cloudinary, repoMerchant MerchantRepository, db *sqlx.DB) fileService {
	return fileService{cloudinaryService: cloudinary, RepoMerchant: repoMerchant, db: db}
}

func (s fileService) uploadImage(ctx *gin.Context, request File) (string, error) {
	userClaims := ctx.MustGet("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(string)

	tx, err := s.db.Beginx()
	if err != nil {
		return "", errors.New("invalid request")
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	merchant, err := s.RepoMerchant.FindByUserID(ctx, tx, userID)
	if err != nil {
		return "", err
	}

	filename := request.File.Filename
	filenameWithoutExt := filepath.Base(filename)
	filenameWithoutExt = strings.TrimSuffix(filename, filepath.Ext(filenameWithoutExt))

	publicID := fmt.Sprintf("merchants/%s/%s/%s", merchant.Name, request.Type, filenameWithoutExt)

	fileData, err := request.File.Open()
	if err != nil {
		return "", errors.New("invalid request")
	}
	response, err := s.cloudinaryService.Upload.Upload(ctx, fileData, uploader.UploadParams{
		PublicID: publicID,
	})

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("invalid request")
	}

	if err := tx.Commit(); err != nil {
		return "", errors.New("invalid request")
	}

	return response.SecureURL, nil
}
