package merchant

import (
	"time"
)

type Merchant struct {
	ID        string    `json:"id" db:"id" validate:"required"`
	UserID    string    `json:"user_id" db:"user_id" validate:"required"`
	Name      string    `json:"name" db:"name" validate:"required"`
	ImageURL  string    `json:"image_url" db:"image_url" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validate:"required"`
}

func NewMerchant() Merchant {
	return Merchant{}
}
