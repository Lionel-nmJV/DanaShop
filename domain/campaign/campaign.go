package campaign

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Campaign struct {
	Id           uuid.UUID `db:"id"`
	MerchantId   uuid.UUID `db:"merchant_id"`
	Name         string    `db:"name"`
	StartDate    time.Time `db:"start_date"`
	EndDate      time.Time `db:"end_date"`
	IsActive     bool      `db:"is_active"`
	VideoUrl     string    `db:"video_url"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	ThumbnailUrl string    `db:"thubmnail_url"`
	Products     []product
}

type product struct {
	CampaignId    uuid.UUID `db:"campaign_id" json:"campaign_id"`
	ProductId     uuid.UUID `db:"product_id" json:"product_id" `
	Campaignprice float64   `db:"campaign_price" json:"campaign_price"`
}

func newCampaign() Campaign {
	return Campaign{}
}

func (campaign Campaign) fromRequest(request createRequest, validate *validator.Validate) (Campaign, error) {
	if err := validate.Struct(request); err != nil {
		return campaign, errors.New("invalid request")
	}

	campaign.Name = request.Name
	campaign.Description = request.Description
	campaign.StartDate = request.StartDate
	campaign.EndDate = request.EndDate
	campaign.Products = request.Products
	campaign.VideoUrl = request.VideoUrl
	campaign.ThumbnailUrl = request.ThumbnailUrl

	return campaign, nil

}
