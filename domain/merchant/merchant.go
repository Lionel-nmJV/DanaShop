package merchant

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Merchant struct {
	ID        string    `json:"id" db:"id" validate:"required"`
	UserID    string    `json:"user_id" db:"user_id" validate:"required"`
	Name      string    `json:"name" db:"name" validate:"required"`
	ImageURL  string    `json:"image_url" db:"image_url" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validate:"required"`
	FocusOn   string    `json:"focus_on" db:"focused_on" validate:"required"`
	Address   string    `json:"address" db:"address" validate:"required"`
}

func newMerchant() Merchant {
	return Merchant{}
}

func (merchant Merchant) fromRequest(req updateRequest, validate *validator.Validate) (Merchant, error) {
	if err := validate.Struct(req); err != nil {
		return merchant, err
	}

	merchant.Name = req.Name
	merchant.Address = req.Address
	merchant.ImageURL = req.ImageURL
	merchant.FocusOn = req.FocusOn

	return merchant, nil

}
