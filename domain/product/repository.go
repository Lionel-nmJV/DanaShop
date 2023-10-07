package product

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ProductRepositoryImpl implements the ProductRepository interface.
type ProductRepositoryImpl struct {
	db *sqlx.DB // Adapt to your PostgreSQL database connection
}

// NewProductRepository creates a new ProductRepositoryImpl.
func NewProductRepository(db *sqlx.DB) *ProductRepositoryImpl {
	// Initialize your PostgreSQL database connection here
	return &ProductRepositoryImpl{db}
}

// FindAllByMerchantID retrieves products by merchant ID and query string with pagination.
func (r *ProductRepositoryImpl) FindAllByMerchantID(ctx *gin.Context, merchantID string, query string, limit int, offset int) ([]*Product, error) {
	SQL := `SELECT "id", "name", "category", "price", "stock", "image_url", "created_at", "updated_at" FROM "products" WHERE "merchant_id"=$1 AND "name" ILIKE '%' || $2 || '%' LIMIT $3 OFFSET $4`
	rows, err := r.db.QueryContext(ctx, SQL, merchantID, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Stock, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt)
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
