package campaign

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	writeRepository
}

type writeRepository interface {
	insertCampaign(ctx context.Context, tx *sqlx.Tx, campaign Campaign) (uuid.UUID, error)
	insertCampaignProducts(ctx context.Context, tx *sqlx.Tx, campaign Campaign) error
}

type CampaignService struct {
	repo repository
	db   *sqlx.DB
}

func newService(repo repository, db *sqlx.DB) CampaignService {
	return CampaignService{
		repo: repo,
		db:   db,
	}
}

func (svc CampaignService) createCampaign(ctx context.Context, campaign Campaign) error {
	tx, err := svc.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	campaginId, err := svc.repo.insertCampaign(ctx, tx, campaign)
	if err != nil {

		return newCustomError(50001, 500, "repository error")
	}

	for i := range campaign.Products {
		campaign.Products[i].CampaignId = campaginId
	}

	if err := svc.repo.insertCampaignProducts(ctx, tx, campaign); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}
