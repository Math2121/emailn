package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaig campaign.NewCampaignRequest) (string, error) {
	args := r.Called(newCampaig)

	return args.String(0), args.Error(1)
}
func (s *CampaignServiceMock) GetById(id string) (*campaign.CampaingResponse, error) {
	args := s.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.CampaingResponse), args.Error(1)
}

func (s *CampaignServiceMock) Cancel(id string) error {
	return nil
}
func (s *CampaignServiceMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CampaignServiceMock) Start(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
