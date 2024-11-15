package ports

import "backend-ad-campaign/internal/core"

type Repository interface {
	CreateCampaign(campaign *core.Campaign) (*core.Campaign, error)
	GetCampaignByID(id string) (*core.Campaign, error)
	GetAllCampaigns() ([]*core.Campaign, error)
	DeleteCampaign(id string) error
	UpdateCampaign(campaign *core.Campaign) (*core.Campaign, error)
	SearchCampaignByName(name string) ([]*core.Campaign, error)
	GetActiveCampaigns() ([]*core.Campaign, error)
}
