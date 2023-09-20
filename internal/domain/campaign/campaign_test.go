package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing Test"
	content  = "html"
	contacts = []string{"teste@gmail.com", "teste2@gmail.com"}
)

func TestNewCampaign(t *testing.T) {

	campaign,_ := NewCampaign(name, content, contacts)
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

	campaign,_ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func TestNewCampaign_CreatedOnIsNotNil(t *testing.T) {

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(t, campaign.CreatedOn, now)
}


func Test_NeCampign_MustValidateName(t *testing.T){
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name is required", err.Error())

}


func Test_NeCampign_MustValidateContent(t *testing.T){
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())

}


func Test_NeCampign_MustValidateContact(t *testing.T){
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contact is required", err.Error())

}