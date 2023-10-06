package product

import (
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

func (r repoProduct) saveProduct(ctx *gin.Context, tx *sqlx.Tx, product Product) (string, error) {
	SQL := `INSERT INTO "products"("merchant_id", "name", "category", "price", "stock", "image_url", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var lastInsertID string
	result := tx.QueryRowContext(ctx, SQL, product.MerchantID, product.Name, product.Category, product.Price, product.Stock, product.ImageURL, product.CreatedAt, product.UpdatedAt)

	err := result.Scan(&lastInsertID)
	if err != nil {
		return lastInsertID, err
	}

	product.ID = lastInsertID

	return lastInsertID, nil
}
