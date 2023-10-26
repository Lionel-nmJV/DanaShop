package campaign

import (
	"time"

	"github.com/google/uuid"
)

type createRequest struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Products    []product `json:"products" validate:"required"`
	VideoUrl    string    `json:"video_url" validate:"required"`
}

type getCampaignsResponse struct {
	Campaigns []campaign `json:"campaigns"`
}

type campaign struct {
	Id           uuid.UUID `json:"campaign_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	TotalProduct int       `json:"total_product"`
	VideoUrl     string    `json:"video_url"`
}
