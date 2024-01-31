package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)

	return tx.Error
}
func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := c.Db.Save(campaign)

	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaignFounded campaign.Campaign
	tx := c.Db.Preload("Contacts").First(&campaignFounded,"id = ?", id)
	return &campaignFounded, tx.Error
}


func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {

	// for i, _ := range campaign.Contacts {

	// 	c.Db.Delete(campaign.Contacts[i])
	// }

	tx := c.Db.Select("Contacts").Delete(campaign)

	return tx.Error
}
