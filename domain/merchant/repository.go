package merchant

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type repoMerchant struct {
}

func NewRepoMerchant() repoMerchant {
	return repoMerchant{}
}

func (r repoMerchant) FindByUserID(ctx *gin.Context, tx *sqlx.Tx, userID string) (Merchant, error) {
	SQL := `SELECT "id", "name", "image_url" FROM "merchants" WHERE "user_id"=$1`
	rows, err := tx.QueryContext(ctx, SQL, userID)
	if err != nil {
		return Merchant{}, err
	}
	defer rows.Close()

	merchant := Merchant{}
	if rows.Next() {
		err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.ImageURL)
		if err != nil {
			panic(err)
		}
		return merchant, nil
	} else {
		return merchant, errors.New("not found")
	}
}

func (r repoMerchant) getMerchantByUserId(ctx *gin.Context, db *sqlx.DB, userID string) (merchantResponse, error) {
	SQL := `SELECT id,name,created_at,updated_at,image_url,focused_on,address FROM "merchants" WHERE "user_id"=$1`

	merchant := merchantResponse{}
	err := db.QueryRowContext(ctx, SQL, userID).
		Scan(&merchant.ID, &merchant.Name, &merchant.CreatedAt, &merchant.UpdatedAt, &merchant.ImageURL, &merchant.FocusOn, &merchant.Address)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return merchant, newCustomError(40401, 404, "not found")
		case err != nil:
			return merchant, newCustomError(50001, 500, "repository error")
		}
	}

	return merchant, nil
}

func (repo repoMerchant) updateMerchantByUserId(ctx *gin.Context, db *sqlx.DB, userID string, merchant Merchant) (bool, error) {
	SQL := `update merchants set
			name=$1,
			focused_on=$2,
			address=$3,
			image_url=$4,
			updated_at=NOW()
			where user_id=$5
			`
	_, err := db.ExecContext(ctx, SQL, merchant.Name, merchant.FocusOn, merchant.Address, merchant.ImageURL, userID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return false, newCustomError(40902, 409, "merchant name already taken")
			} else {
				return false, newCustomError(50001, 500, "repository error")
			}
		}
	}

	return true, nil
}

func (r repoMerchant) GetMerchantAnalytics(ctx *gin.Context, db *sqlx.DB, merchantID string) (Analytics, error) {
	var analytics Analytics

	query := `
        SELECT 
            COALESCE(SUM(total_price), 0) AS total_price,
            COALESCE(SUM(quantity), 0) AS total_product_sold,
            COALESCE(SUM(total_discount), 0) AS total_discount,
            COALESCE(COUNT(*) FILTER (WHERE status = 'pending'), 0) AS total_transaction_pending,
            COALESCE(COUNT(*) FILTER (WHERE status = 'success'), 0) AS total_transaction_success,
            COALESCE(COUNT(*) FILTER (WHERE status = 'failed'), 0) AS total_transaction_fail
        FROM orders
        WHERE merchant_id = $1
    `

	err := db.QueryRowContext(ctx, query, merchantID).Scan(
		&analytics.TotalPrice,
		&analytics.TotalProductSold,
		&analytics.TotalDiscount,
		&analytics.TotalTransactionPending,
		&analytics.TotalTransactionSuccess,
		&analytics.TotalTransactionFail,
	)

	if err != nil {
		fmt.Println(err.Error())
		return Analytics{}, err
	}

	return analytics, nil
}
