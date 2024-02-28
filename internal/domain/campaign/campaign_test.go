package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Campaing Test"
	createdBy = "teste@teste.com"
	content   = "Body Hi"
	contacts  = []string{"teste@gmail.com", "teste2@gmail.com"}
	fake      = faker.New()
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(campaign.Name, name)
	assert.Equal(createdBy, campaign.CreatedBy)

}

func TestNewCampaignID(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(t, campaign.ID)
}
func TestNewCampaignStatus(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(Pending, campaign.Status)
}

func TestNewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts, createdBy)

	assert.Equal("name is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {

	assert := assert.New(t)
	fake := faker.New()

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Equal("name is required with max 24", err.Error())

}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)
	fake := faker.New()

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts, createdBy)

	assert.Equal("content is required with max 1024", err.Error())

}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil, createdBy)

	assert.Equal("contacts is required with min 1", err.Error())

}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"}, createdBy)

	assert.Equal("email is invalid", err.Error())

}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, contacts, "")

	assert.Equal("createdby is invalid", err.Error())

}

func Test_Campaign_ChangeStatus_Done(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Done()

	assert.Equal(campaign.Status, Done)

}
func Test_Campaign_ChangeStatus_Delete(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Delete()

	assert.Equal(campaign.Status, Delete)

}

func Test_Campaign_ChangeStatus_Fail(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Fail()

	assert.Equal(campaign.Status, Fail)

}

func Test_Campaign_ChangeStatus_Started(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Started()

	assert.Equal(campaign.Status, Started)

}
