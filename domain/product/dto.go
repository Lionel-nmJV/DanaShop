package product

import "time"

type paginateProductsResponse struct {
	Products   []productResponses `json:"products"`
	Pagination pagination         `json:"pagination"`
}

type pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type productResponses struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	Weight      int       `json:"weight"`
	Threshold   int       `json:"threshold"`
	IsNew       bool      `json:"is_new"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type createRequest struct {
	Name        string  `json:"name"validate:"required"`
	Category    string  `json:"category"validate:"required"`
	Stock       int     `json:"stock"validate:"required"`
	ImageURL    string  `json:"image_url"validate:"required"`
	Price       float64 `json:"price"validate:"required"`
	Weight      int     `json:"weight"validate:"required"`
	Threshold   int     `json:"threshold"validate:"required"`
	IsNew       bool    `json:"is_new"validate:"required"`
	Description string  `json:"description"validate:"required"`
}

// Request structure for updating a product
type updateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	ImageURL    string  `json:"image_url" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Weight      int     `json:"weight"validate:"required"`
	Threshold   int     `json:"threshold"validate:"required"`
	IsNew       bool    `json:"is_new"validate:"required"`
	Description string  `json:"description"validate:"required"`
}

// Request structure for deleting a product
type deleteRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}
