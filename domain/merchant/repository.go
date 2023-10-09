package merchant

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RepoMerchant struct {
}

func NewRepoMerchant() RepoMerchant {
	return RepoMerchant{}
}

func (r RepoMerchant) FindByUserID(ctx *gin.Context, tx *sqlx.Tx, userID string) (Merchant, error) {
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

func (r RepoMerchant) GetMerchantByUserId(ctx *gin.Context, db *sqlx.DB, userID string) (merchantResponse, error) {
	SQL := `SELECT id,name,created_at,updated_at,image_url FROM "merchants" WHERE "user_id"=$1`

	merchant := merchantResponse{}
	err := db.QueryRowContext(ctx, SQL, userID).
		Scan(&merchant.ID, &merchant.Name, &merchant.CreatedAt, &merchant.UpdatedAt, &merchant.ImageURL)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return merchant, NewCustomError(40401, 404, "not found")
		case err != nil:
			return merchant, NewCustomError(50001, 500, "repository error")
		}
	}
	return merchant, nil
}
