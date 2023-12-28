package database

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaings []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaings = append(c.campaings, *campaign)

	return nil


}

func (c *CampaignRepository) Get() []campaign.Campaign {
	return c.campaings
}

