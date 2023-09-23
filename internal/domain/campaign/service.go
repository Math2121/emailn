package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaging contract.NewCampaign) (string, error) {

	campaign, _ := NewCampaign(newCampaging.Name, newCampaging.Content, newCampaging.Emails)

	s.Repository.Save(campaign)
	return campaign.ID, nil
}
