package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	readRepository
}

type readRepository interface {
	GetMerchantByUserId(ctx *gin.Context, db *sqlx.DB, merchantID string) (merchantResponse, error)
}

type merchantService struct {
	repo repository
	db   *sqlx.DB
}

func newService(repo repository, db *sqlx.DB) merchantService {
	return merchantService{
		repo: repo,
		db:   db,
	}
}

func (m merchantService) GetMerchantProfileById(ctx *gin.Context, merchantId string) (merchant merchantResponse, err error) {

	merchant, err = m.repo.GetMerchantByUserId(ctx, m.db, merchantId)
	if err != nil {
		return merchantResponse{}, err
	}

	return merchant, nil
}
