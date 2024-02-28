package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaingRepositoryMock struct {
	mock.Mock
}

var (
	newCampaign = campaign.NewCampaignRequest{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = campaign.ServiceImp{}
)

func (r *CampaingRepositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)

	return args.Error(0)
}
func (r *CampaingRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)

	return args.Error(0)
}

func (r *CampaingRepositoryMock) Get() ([]campaign.Campaign, error) {
	return nil, nil
}
func (r *CampaingRepositoryMock) GetById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (r *CampaingRepositoryMock) Delete(campaing *campaign.Campaign) error {
	args := r.Called(campaing)

	return args.Error(0)
}

func (r *CampaingRepositoryMock) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	args := r.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]campaign.Campaign), nil
}
