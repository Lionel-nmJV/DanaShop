package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	readRepository
	writeRepository
}

type readRepository interface {
	getMerchantByUserId(ctx *gin.Context, db *sqlx.DB, merchantID string) (merchantResponse, error)
}

type writeRepository interface {
	updateMerchantByUserId(ctx *gin.Context, db *sqlx.DB, userID string, merchant Merchant) (bool, error)
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

func (svc merchantService) getMerchantProfileById(ctx *gin.Context, userId string) (merchant merchantResponse, err error) {

	merchant, err = svc.repo.getMerchantByUserId(ctx, svc.db, userId)
	if err != nil {
		return merchantResponse{}, err
	}
	return merchant, nil
}

func (svc merchantService) updateMerchantProfileById(ctx *gin.Context, userId string, merchant Merchant) (bool, error) {

	ok, err := svc.repo.updateMerchantByUserId(ctx, svc.db, userId, merchant)
	if err != nil {
		return false, err
	}
	return ok, nil
}
