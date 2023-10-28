package campaign

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	writeRepository
	readRepository
}

type writeRepository interface {
	insertCampaign(ctx context.Context, tx *sqlx.Tx, campaign Campaign) (uuid.UUID, error)
	insertCampaignProducts(ctx context.Context, tx *sqlx.Tx, campaign Campaign) error
	updateCampaignStatus(ctx context.Context, db *sqlx.DB, isActive bool, campaginId uuid.UUID, merchantId uuid.UUID) (int64, error)
}

type readRepository interface {
	findAllCampaigns(ctx context.Context, db *sqlx.DB, merchantID string, query string, limit interface{}, offset int) ([]campaign, error)
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

func (svc CampaignService) deactivateCampaign(ctx context.Context, campaginId uuid.UUID, merchantId uuid.UUID) (int64, error) {

	rowsUpdated, err := svc.repo.updateCampaignStatus(ctx, svc.db, false, campaginId, merchantId)
	if err != nil {
		return rowsUpdated, newCustomError(50001, 500, "repository error")
	}
	return rowsUpdated, nil
}

func (svc CampaignService) findAllCampaigns(ctx context.Context, query string, merchantId string, limit int, offset int) ([]campaign, int, error) {
	campaigns, err := svc.repo.findAllCampaigns(ctx, svc.db, merchantId, query, limit, offset)

	if err != nil {
		return []campaign{}, 0, newCustomError(50001, 500, "repository error")
	}

	if len(campaigns) == 0 {
		return []campaign{}, 0, newCustomError(40401, 404, "campaigns not found")
	}

	allCampaigns, err := svc.repo.findAllCampaigns(ctx, svc.db, merchantId, query, nil, 0)

	if err != nil {
		return []campaign{}, 0, newCustomError(50001, 500, "repository error")
	}

	totalCampaign := len(allCampaigns)

	return campaigns, totalCampaign, nil
}
