package campaign

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CampaingResponse, error)
	Cancel(id string) (error)
	Delete(id string) error
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

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, "")
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
		if !errors.Is(err, gorm.ErrRecordNotFound){
			return nil, internalerror.ErrInternal
		}
		return nil, internalerror.ProcessErrorToReturn(err)
	}

	if campaign == nil {
		return nil, nil
	}

	return &contract.CampaingResponse{
		Name: campaign.Name,
		Status: campaign.Status,
		ID: campaign.ID,
		Content: campaign.Content,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImp) Cancel(id string) ( error) {

	campaing, err := s.Repository.GetById(id)

	if err != nil {
		return internalerror.ProcessErrorToReturn(err)
	}

	if campaing.Status != Pending {

		return errors.New("Campaign already started")
	}

	campaing.Cancel()
	err = s.Repository.Update(campaing)

	if err != nil {
		return internalerror.ErrInternal
	}


	return nil
}
func (s *ServiceImp) Delete(id string) ( error) {

	campaing, err := s.Repository.GetById(id)

	if err != nil {
		return internalerror.ProcessErrorToReturn(err)
	}

	if campaing.Status != Pending {

		return errors.New("Campaign status invalid")
	}

	campaing.Delete()
	err = s.Repository.Delete(campaing)

	if err != nil {
		return internalerror.ErrInternal
	}


	return nil
}

