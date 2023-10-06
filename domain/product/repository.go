package product

import (
	"fmt"
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
		fmt.Println(err.Error())
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

// Update updates a product's details in the database.
func (r *ProductRepositoryImpl) update(product *Product) error {
	// Implement database update operation for the product
	// Use the product.ID to identify the product to update
	_, err := r.db.Exec("UPDATE products SET name=$1, category=$2, price=$3, stock=$4, image_url=$5, updated_at=$6 WHERE id=$7", product.Name, product.Category, product.Price, product.Stock, product.ImageURL, product.UpdatedAt, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) delete(id string) error {
	// Implement database delete operation for the product
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
