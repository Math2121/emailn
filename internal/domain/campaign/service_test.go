package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerror "emailn/internal/internalError"
	internalMock "emailn/internal/test/internal-mock"
	internalmock "emailn/internal/test/internal-mock"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Test Y",
		Content:   "teste",
		Emails:    []string{"test@example.com"},
		CreatedBy: "teste@teste.com.br",
	}
	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalMock.CampaingRepositoryMock)
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

	repositoryMock := new(internalMock.CampaingRepositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaingRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerror.ErrInternal, err))

}

func Test_Get_By_ID(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).
		Return(campaign, nil)

	service.Repository = repositoryMock

	campaingReturned, _ := service.GetById(campaign.ID)

	assert.Equal(campaign.Name, campaingReturned.Name)
	assert.Equal(campaign.Status, campaingReturned.Status)
	assert.Equal(campaign.CreatedBy, campaingReturned.CreatedBy)
}

func Test_Get_By_ID_Error(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	repositoryMock.On("GetById", mock.Anything).
		Return(nil, errors.New("Something"))

	service.Repository = repositoryMock

	_, err := service.GetById(campaign.ID)

	assert.Equal(internalerror.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnRecordNotFound_when_campaign_not_Exist(t *testing.T) {
	assert := assert.New(t)
	campaingIdInvalid := "invalid"
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	service.Repository = repositoryMock

	err := service.Delete(campaingIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())

}

func Test_Delete_ReturnStatusInvalid(t *testing.T) {
	assert := assert.New(t)
	campaing := campaign.Campaign{
		ID:     "1",
		Status: campaign.Done,
	}
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	repositoryMock.On("GetById", mock.Anything).Return(&campaing, nil)

	service.Repository = repositoryMock

	err := service.Delete(campaing.ID)

	assert.Equal(err.Error(), "Campaign status invalid")

}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign(
		"Test12222",
		"bodsssssy",
		[]string{"teste@example.com"},
		newCampaign.CreatedBy,
	)
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	repositoryMock.On("GetById", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaing *campaign.Campaign) bool {
		return campaignFound == campaing

	})).Return(errors.New("error to delete campaing"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerror.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnNilwhen_delete_has_success(t *testing.T) {
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign(
		"Test12222",
		"bodsssy",
		[]string{"teste@example.com"},
		newCampaign.CreatedBy,
	)
	repositoryMock := new(internalmock.CampaingRepositoryMock)

	repositoryMock.On("GetById", mock.Anything).Return(campaignFound, nil)

	repositoryMock.On("Delete", mock.MatchedBy(func(campaing *campaign.Campaign) bool {

		return campaignFound == campaing

	})).Return(nil)

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)

}
