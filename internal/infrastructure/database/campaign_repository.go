package database

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaings []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaings = append(c.campaings, *campaign)

	return nil


}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaings, nil
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaignFounded campaign.Campaign

	for _, campaign := range c.campaings {
		if campaign.ID == id {
			campaignFounded = campaign
			break
		}
	}
	

	return &campaignFounded, nil
}


