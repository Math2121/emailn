package campaign

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CampaingResponse, error)
}

type campaingAttributes struct {
	Name string
	Status string
	ID string
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerror.ErrInternal
	}
	return campaign.ID, nil
}

func (s *ServiceImp) GetById(id string) (*contract.CampaingResponse, error) {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return &contract.CampaingResponse{}, internalerror.ErrInternal
	}
	return &contract.CampaingResponse{
		Name: campaign.Name,
		Status: campaign.Status,
		ID: campaign.ID,
		Content: campaign.Content,
	}, nil
}

