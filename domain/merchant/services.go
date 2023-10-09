package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	readRepository
}

type readRepository interface {
	GetMerchantByUserId(ctx *gin.Context, db *sqlx.DB, merchantID string) (merchantResponse, error)
}

type MerchantService struct {
	repo Repository
	db   *sqlx.DB
}

func newService(repo Repository, db *sqlx.DB) MerchantService {
	return MerchantService{
		repo: repo,
		db:   db,
	}
}

func (m MerchantService) GetMerchantProfileById(ctx *gin.Context, merchantId string) (merchant merchantResponse, err error) {

	merchant, err = m.repo.GetMerchantByUserId(ctx, m.db, merchantId)
	if err != nil {
		return merchantResponse{}, err
	}

	return merchant, nil
}
