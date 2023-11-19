package merchant

import "time"

type merchantResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Analytics struct {
	TotalPrice              int `json:"total_price"`
	TotalProductSold        int `json:"total_product_sold"`
	TotalDiscount           int `json:"total_discount"`
	TotalTransactionPending int `json:"total_transaction_pending"`
	TotalTransactionSuccess int `json:"total_transaction_success"`
	TotalTransactionFail    int `json:"total_transaction_fail"`
}
