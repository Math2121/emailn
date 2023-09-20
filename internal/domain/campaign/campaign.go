package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	if name == ""{
		return nil, errors.New("Name is required")
	}
	if content == ""{
		return nil, errors.New("content is required")
	}
	
	if len(contacts) == 0 {
		return nil, errors.New("Contact is required")
	}

	for indx, value := range emails {
		contacts[indx].Email = value
	}
	
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts: contacts,
	}, nil
}
