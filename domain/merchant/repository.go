package merchant

import (
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
