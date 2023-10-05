package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type repoProduct struct {
}

func newRepoProduct() repoProduct {
	return repoProduct{}
}

func (r repoProduct) findAllByMerchantID(ctx *gin.Context, tx *sqlx.Tx, merchantID string, query string, limit interface{}, offset int) ([]productResponses, error) {
	SQL := `SELECT "name", "category", "price", "stock", "image_url" FROM "products" WHERE "merchant_id"=$1 AND "name"  ILIKE CONCAT('%', $2::text, '%') LIMIT $3 OFFSET $4`
	rows, err := tx.QueryContext(ctx, SQL, merchantID, query, limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []productResponses
	for rows.Next() {
		product := productResponses{}
		err := rows.Scan(&product.Name, &product.Category, &product.Price, &product.Stock, &product.ImageURL)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
