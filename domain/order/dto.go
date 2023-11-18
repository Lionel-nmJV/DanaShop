package order

type paginateOrdersResponse struct {
	Orders     []orderResponse `json:"orders"`
	Pagination pagination      `json:"pagination"`
}

type pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type orderResponse struct {
	ID              string  `json:"id"`
	ProductName     string  `json:"product_name"`
	Quantity        int     `json:"quantity"`
	ProductPrice    float64 `json:"price"`
	UserName        string  `json:"user"`
	UserPhoneNumber string  `json:"phone_number"`
	UserAddress     string  `json:"address"`
	TotalPrice      float64 `json:"total_price"`
	Status          string  `json:"status"`
}
