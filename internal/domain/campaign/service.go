package campaign

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaging contract.NewCampaign) (string, error) {

	campaign, _ := NewCampaign(newCampaging.Name, newCampaging.Content, newCampaging.Emails)

	err := s.Repository.Save(campaign)

	if err != nil {
		return "", internalerror.ErrInternal
	}
	return campaign.ID, nil
}
