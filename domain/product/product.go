package product

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type Product struct {
	ID          string    `json:"id"db:"id"`
	MerchantID  string    `json:"merchant_id"db:"merchant_id"`
	Name        string    `json:"name"db:"name"`
	Category    string    `json:"category"db:"category"`
	Price       float64   `json:"price"db:"price"`
	Stock       int       `json:"stock"db:"stock"`
	ImageURL    string    `json:"image_url"db:"image_url"`
	Weight      int       `json:"weight"db:"weight"`
	Threshold   int       `json:"threshold"db:"threshold"`
	IsNew       bool      `json:"is_new"db:"is_new"`
	Description string    `json:"description"db:"description"`
	CreatedAt   time.Time `json:"created_at"db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"db:"updated_at"`
}

func NewProduct() Product {
	return Product{}
}

func (p Product) formAddProduct(request createRequest, validate *validator.Validate) (Product, error) {
	err := validate.Struct(request)
	if err != nil {
		return p, errors.New("invalid request")
	}
	p.Name = request.Name
	p.Category = request.Category
	p.Stock = request.Stock
	p.ImageURL = request.ImageURL
	p.Price = request.Price
	p.Weight = request.Weight
	p.Threshold = request.Threshold
	p.IsNew = request.IsNew
	p.Description = request.Description
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return p, nil
}

func (p *Product) UpdateProduct(request updateRequest, validate *validator.Validate) error {
	err := validate.Struct(request)
	if err != nil {
		return errors.New("invalid update request")
	}

	p.Name = request.Name
	p.Category = request.Category
	p.Stock = request.Stock
	p.ImageURL = request.ImageURL
	p.Price = request.Price
	p.UpdatedAt = time.Now()

	return nil
}

func (p *Product) DeleteProduct() {

}
