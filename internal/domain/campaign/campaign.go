package campaign

import (
	internalerror "emailn/internal/internalError"
	"time"

	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
	Fail    string = "Fail"
	Cancel  string = "Cancel"
	Delete  string = "Delete"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;not null"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required"`
	Status    string    `gorm:"size:20"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedBy string    `validate:"email" gorm:"size:100;not null"`
	UpdatedOn time.Time 
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for indx, value := range emails {
		contacts[indx].Email = value
		contacts[indx].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Status:    Pending,
		Contacts:  contacts,
		CreatedBy: createdBy,
	}

	error := internalerror.ValidateStruct(campaign)

	if error == nil {
		return campaign, nil
	}

	return nil, error
}

func (c *Campaign) Cancel() {
	c.Status = Cancel
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = Delete
	c.UpdatedOn = time.Now()
}

// TODO: make unit test
func (c *Campaign) Fail() {
	c.Status = Fail
	c.UpdatedOn = time.Now()
}

// TODO: make unit test
func (c *Campaign) Started() {
	c.Status = Started
	c.UpdatedOn = time.Now()
}

// TODO: make unit test
func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedOn = time.Now()
}
