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

func (r repoProduct) updateProduct(ctx *gin.Context, tx *sqlx.Tx, productID string, product Product) error {
	SQL := `UPDATE "products" SET "name"=$1, "category"=$2, "price"=$3, 
                      "stock"=$4, "image_url"=$5, "updated_at"=$6 WHERE "id"=$7`

	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Category, product.Price, product.Stock, product.ImageURL, product.UpdatedAt, productID)

	if err != nil {
		return err
	}

	return nil
}

func (r repoProduct) deleteProduct(ctx *gin.Context, tx *sqlx.Tx, productID string) error {
	SQL := `DELETE FROM "products" WHERE "id" = $1`

	_, err := tx.ExecContext(ctx, SQL, productID)

	if err != nil {
		return err
	}

	return nil
}

func (r repoProduct) findProductByID(ctx *gin.Context, tx *sqlx.Tx, productID string) (Product, error) {
	SQL := `SELECT "id", "merchant_id", "name", "category", "price", "stock", "image_url", "created_at", "updated_at" FROM "products" WHERE "id" = $1`
	var product Product
	err := tx.GetContext(ctx, &product, SQL, productID)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
