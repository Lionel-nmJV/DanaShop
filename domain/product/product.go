package product

import (
	"time"
)

// Product represents the product entity.
type Product struct {
	ID         string    `json:"id"`
	MerchantID string    `json:"merchant_id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	ImageURL   string    `json:"image_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ProductRepository defines the repository interface for products.
type ProductRepository interface {
	FindByID(id string) (*Product, error)
	update(product *Product) error
	delete(id string) error
}

// ProductService provides business logic for managing products.
type ProductService struct {
	repository ProductRepository
}

// NewProductService creates a new ProductService.
func NewProductService(repository ProductRepository) *ProductService {
	return &ProductService{repository}
}

// UpdateProduct updates a product's details.
func (s *ProductService) UpdateProduct(product *Product) error {
	// Implement validation or business rules here if needed
	return s.repository.update(product)
}

// DeleteProduct deletes a product by its ID.
func (s *ProductService) DeleteProduct(id string) error {
	// Implement validation or business rules here if needed
	return s.repository.delete(id)
}
