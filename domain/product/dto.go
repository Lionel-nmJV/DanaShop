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
