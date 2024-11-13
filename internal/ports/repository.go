package ports

import "backend-ad-campaign/internal/core"

type Repository interface {
	CreateCampaign(campaign *core.Campaign) (*core.Campaign, error)
}
