package product

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
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	ImageURL string  `json:"image_url"`
}

type createRequest struct {
	Name     string  `json:"name"validate:"required"`
	Category string  `json:"category"validate:"required"`
	Stock    int     `json:"stock"validate:"required"`
	ImageURL string  `json:"image_url"validate:"required"`
	Price    float64 `json:"price"validate:"required"`
}

// Request structure for updating a product
type updateRequest struct {
	Name     string  `json:"name" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Stock    int     `json:"stock" validate:"required"`
	ImageURL string  `json:"image_url" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

// Request structure for deleting a product
type deleteRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}
