package campaign

import "time"

type createRequest struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Products    []product `json:"products" validate:"required"`
	VideoUrl    string    `json:"video_url" validate:"required"`
}
