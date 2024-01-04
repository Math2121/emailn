package campaign

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = ServiceImp{}
)

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)

	return args.Error(0)
}
func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}
func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}


func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

}
func Test_Create_ValidateDomainError(t *testing.T) {

	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerror.ErrInternal, err))
}
func Test_Save_Campaign(t *testing.T) {

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := ServiceImp{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service := ServiceImp{Repository: repositoryMock}

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerror.ErrInternal, err))

}

func Test_Get_By_ID(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)

	campaign,_ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).
	Return(campaign, nil)

	service := ServiceImp{Repository: repositoryMock}

	campaingReturned, _ := service.GetById(campaign.ID)

	assert.Equal(campaign.Name, campaingReturned.Name)
	assert.Equal(campaign.Status, campaingReturned.Status)
}


func Test_Get_By_ID_Error(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)

	campaign,_ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock.On("GetById", mock.Anything).
	Return(nil, errors.New("Something"))

	service := ServiceImp{Repository: repositoryMock}

	_, err := service.GetById(campaign.ID)

	assert.Equal(internalerror.ErrInternal.Error(), err.Error())

}
