package mail

import (
	"crypto/tls"
	"emailn/internal/domain/campaign"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {

	message := gomail.NewMessage()
	dialer := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	var emails []string

	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	message.SetHeader("From", "teste@gmail.com")
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", campaign.Name)
	message.SetBody("text/html", campaign.Content)

	return dialer.DialAndSend(message)
}
