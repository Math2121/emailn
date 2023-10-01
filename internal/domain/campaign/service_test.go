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

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)

	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "test",
		Content: "Teste 2",
		Emails:  []string{"teste@gmail.co"},
	}

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Save_Campaign(t *testing.T) {

	repositoryMock := new(repositoryMock)
	newCampaign := contract.NewCampaign{
		Name:    "test",
		Content: "Teste 2",
		Emails:  []string{"teste@gmail.co"},
	}
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content || 
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)

}


func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	newCampaign := contract.NewCampaign{
		Name:    "test",
		Content: "Teste 2",
		Emails:  []string{"teste@gmail.co"},
	}
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service := Service{Repository: repositoryMock}

	_, err := service.Create(newCampaign)
	assert.True(errors.Is(internalerror.ErrInternal, err))

}