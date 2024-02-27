package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerror "emailn/internal/internalError"
	internalMock "emailn/internal/test/internal-mock"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Test Y",
		Content:   "teste",
		Emails:    []string{"test@example.com"},
		CreatedBy: "teste@teste.com.br",
	}
	campaignPendenting *campaign.Campaign
	repositoryMock     *internalMock.CampaingRepositoryMock
	service            = campaign.ServiceImp{
		Repository: repositoryMock,
	}
	campaignStated *campaign.Campaign
)

func SetUp() {

	repositoryMock = new(internalMock.CampaingRepositoryMock)
	service = campaign.ServiceImp{
		Repository: repositoryMock,
	}

	campaignPendenting, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	campaignStated = &campaign.Campaign{
		ID:     "1",
		Status: campaign.Started,
	}

}

func setUpUpdateRepository(err error) {
	repositoryMock.On("Update", mock.Anything).Return(err)
}

func setUpSendEmail(err error) {
	sendMail := func(campaign *campaign.Campaign) error {
		return err
	}

	service.SendMail = sendMail
}

func Test_Create_Campaign(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

}
func Test_Create_ValidateDomainError(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerror.ErrInternal, err))
}
func Test_Save_Campaign(t *testing.T) {
	SetUp()
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerror.ErrInternal, err))

}

func Test_Get_By_ID(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaignPendenting.ID
	})).
		Return(campaignPendenting, nil)

	service.Repository = repositoryMock

	campaingReturned, _ := service.GetById(campaignPendenting.ID)

	assert.Equal(campaignPendenting.Name, campaingReturned.Name)
	assert.Equal(campaignPendenting.Status, campaingReturned.Status)
	assert.Equal(campaignPendenting.CreatedBy, campaingReturned.CreatedBy)
}

func Test_Get_By_ID_Error(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetById", mock.Anything).
		Return(nil, errors.New("Something"))

	_, err := service.GetById("invalid_campaign")

	assert.Equal(internalerror.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnRecordNotFound_when_campaign_not_Exist(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete("invalid_campaign")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())

}

func Test_Delete_ReturnStatusInvalid(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(campaignStated, nil)

	err := service.Delete(campaignStated.ID)

	assert.Equal(err.Error(), "Campaign status invalid")

}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(campaignPendenting, nil)
	repositoryMock.On("Delete", mock.Anything).Return(errors.New("error to delete campaing"))

	err := service.Delete(campaignPendenting.ID)

	assert.Equal(internalerror.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnNilwhen_delete_has_success(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(campaignPendenting, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaing *campaign.Campaign) bool {

		return campaignPendenting == campaing

	})).Return(nil)

	err := service.Delete(campaignPendenting.ID)

	assert.Nil(err)

}

func Test_Start_ReturnRecordNotFound_when_campaign_not_Exist(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start("invalid")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())

}

func Test_Start_ReturnRecordNotFound_when_campaign_is_pending(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetById", mock.Anything).Return(campaignStated, nil)

	err := service.Start(campaignStated.ID)

	assert.Equal("Campaign status invalid", err.Error())

}

func Test_Start_ReturnNil_when_updated_campaign_to_started(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(campaignStated, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(campaignUpdated *campaign.Campaign) bool {
		return campaignStated.ID == campaignUpdated.ID && campaignUpdated.Status == campaign.Started

	})).Return(nil)

	sendMail := func(campaing *campaign.Campaign) error {
		return nil
	}

	service.SendMail = sendMail

	service.Start(campaignStated.ID)

	assert.Equal(campaign.Started, campaignStated.Status)

}
func Test_Start_SendEmailUpdateStarted_StatusIsStarted_WhenFail(t *testing.T) {
	SetUp()
	setUpSendEmail(errors.New("error to send email"))

	repositoryMock.On("Update", mock.MatchedBy(func(campaignUpdated *campaign.Campaign) bool {
		return campaignStated.ID == campaignUpdated.ID && campaignUpdated.Status == campaign.Fail

	})).Return(nil)

	service.SendEmailAndUpdateStatus(campaignStated)

	repositoryMock.AssertExpectations(t)

}

func Test_Start_SendEmailUpdateStarted_StatusIsDone(t *testing.T) {
	SetUp()
	setUpSendEmail(nil)

	repositoryMock.On("Update", mock.MatchedBy(func(campaignUpdated *campaign.Campaign) bool {
		return campaignPendenting.ID == campaignUpdated.ID && campaignUpdated.Status == campaign.Done

	})).Return(nil)

	service.SendEmailAndUpdateStatus(campaignPendenting)

	repositoryMock.AssertExpectations(t)

}
