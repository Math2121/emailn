package main

import (
	// "emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	_ = database.NewDb()
	// repository := database.CampaignRepository{Db: db}

	// for {
	// 	campaigns, _ := repository.GetCampaignsToBeSent()

	// 	for _, campaign := range campaigns {
	// 		print(campaign.ID)
	// 	}

	// 	println("Sleeping for 1 hour")

	// }

}
