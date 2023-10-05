package product

import (
	"time"
)

type Product struct {
	ID         string    `json:"id"db:"id"validate:"required"`
	MerchantID string    `json:"merchant_id"db:"merchant_id"validate:"required"`
	Name       string    `json:"name"db:"name"validate:"required"`
	Category   string    `json:"category"db:"category"validate:"required"`
	Price      float64   `json:"price"db:"price"validate:"required"`
	Stock      int       `json:"stock"db:"stock"validate:"required"`
	ImageURL   string    `json:"image_url"db:"image_url"validate:"required"`
	CreatedAt  time.Time `json:"created_at"db:"created_at"validate:"required"`
	UpdatedAt  time.Time `json:"updated_at"db:"updated_at"validate:"required"`
}
