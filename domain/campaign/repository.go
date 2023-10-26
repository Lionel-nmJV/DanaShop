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

func (postgres postgres) findAllCampaigns(ctx context.Context, db *sqlx.DB, merchantID string, query string, limit int, offset int) ([]campaign, error) {
	querySQL := `SELECT
					campaigns.id,
					campaigns.name,
					campaigns.description,
					campaigns.start_date,
					campaigns.end_date,
					campaigns.video_url,
					COUNT(campaigns_products.product_id) AS total_product
				FROM
					campaigns
				JOIN
					campaigns_products
				ON
					campaigns.id = campaigns_products.campaign_id
				WHERE 
					campaigns.merchant_id = $1
				AND 
					campaigns.name  ILIKE CONCAT('%', $2::text, '%')
				GROUP BY
					campaigns.id
				LIMIT $3 OFFSET $4	
				`

	rows, err := db.QueryContext(ctx, querySQL, merchantID, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var campaigns []campaign
	for rows.Next() {
		campaign := campaign{}
		err := rows.Scan(
			&campaign.Id,
			&campaign.Name,
			&campaign.Description,
			&campaign.StartDate,
			&campaign.EndDate,
			&campaign.VideoUrl,
			&campaign.TotalProduct)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
