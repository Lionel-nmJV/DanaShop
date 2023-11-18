package order

import (
	"time"
)

type Order struct {
	ID              string    `db:"id"`
	MerchantID      string    `db:"merchant_id"`
	UserID          string    `db:"user_id"`
	UserName        string    `db:"user_name"`
	UserAddress     string    `db:"user_address"`
	UserPhoneNumber string    `db:"user_phone_number"`
	CreatedAt       time.Time `db:"created_at"`
	ProductID       string    `db:"product_id"`
	ProductName     string    `db:"product_name"`
	ProductImageURL string    `db:"product_image_url"`
	Quantity        int       `db:"quantity"`
	ProductPrice    float64   `db:"product_price"`
	SubTotalPrice   float64   `db:"sub_total_price"`
	AdditionalFee   float64   `db:"additional_fee"`
	ShippingCost    float64   `db:"shipping_cost"`
	TotalDiscount   float64   `db:"total_discount"`
	TotalPrice      float64   `db:"total_price"`
	Status          string    `db:"status"`
	InvoiceID       string    `db:"invoice_id"`
}

func newOrder() Order {
	return Order{}
}
