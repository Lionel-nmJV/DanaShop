package campaign

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type postgres struct {
}

func newPostgres() postgres {
	return postgres{}
}

func (postgres postgres) insertCampaign(ctx context.Context, tx *sqlx.Tx, campaign Campaign) (uuid.UUID, error) {
	SQL := `INSERT INTO "campaigns" (
				"merchant_id", 
				"name" , 
				"start_date" ,
				"end_date",
				"video_url",
				"description",
				"created_at") 
				VALUES ($1, $2, $3, $4, $5, $6, NOW()) returning id `

	var campaignId uuid.UUID
	if err := tx.QueryRowContext(
		ctx,
		SQL,
		campaign.MerchantId,
		campaign.Name,
		campaign.StartDate,
		campaign.EndDate,
		campaign.VideoUrl,
		campaign.Description).
		Scan(&campaignId); err != nil {
		return uuid.UUID{}, err
	}
	return campaignId, nil
}

func (postgres postgres) insertCampaignProducts(ctx context.Context, tx *sqlx.Tx, campaign Campaign) error {

	SQL := `INSERT INTO "campaigns_products" ("campaign_id", "product_id", "campaign_price")
			VALUES (:campaign_id, :product_id, :campaign_price)`

	if _, err := tx.NamedExecContext(ctx, SQL, campaign.Products); err != nil {

		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23503" {
				return newCustomError(40401, 404, "product id not found")
			} else {
				return newCustomError(50001, 500, "repository error")
			}
		}
	}

	return nil
}

func (postgres postgres) updateCampaignStatus(
	ctx context.Context,
	db *sqlx.DB,
	isActive bool,
	campaignId uuid.UUID,
	merchantId uuid.UUID) (int64, error) {
	query := `UPDATE campaigns SET is_active=$1 where id = $2 AND merchant_id=$3 AND is_active <> false`

	result, err := db.ExecContext(ctx, query, isActive, campaignId, merchantId)
	if err != nil {
		return 0, err
	}

	affected, _ := result.RowsAffected()
	return affected, nil
}
