package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing Test"
	content  = "Body Hi"
	contacts = []string{"teste@gmail.com", "teste2@gmail.com"}
)

func TestNewCampaign(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)
	println(campaign.ID)

	if campaign.Name != name {
		t.Error("expected correct name ", campaign.Name)
	}
	if campaign.Content != content {
		t.Error("expected correct content ", campaign.Content)
	}

	if len(campaign.Contacts) != len(contacts) {
		t.Error("expected correct contacts ", campaign.Contacts)
	}

}

func TestNewCampaignID(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func TestNewCampaign_CreatedOnIsNotNil(t *testing.T) {

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(t, campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {

	assert := assert.New(t)
	fake := faker.New()

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())

}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)
	fake := faker.New()

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", err.Error())

}


func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())

}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())

}