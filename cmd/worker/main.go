package main

import (
	// "emailn/internal/domain/campaign"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}

	campaigns, err := repository.GetCampaignsToBeSent()

	campaignService := campaign.ServiceImp{
		Repository: &repository,
		SendMail:   mail.SendMail,
	}
	for {
		if err != nil {
			println(err.Error())
		}
		for _, campaign := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaign)
			println("Campaigns sent: ", campaign.ID)
		}
		time.Sleep(60 * 60 * time.Second)
	}

}
