package main

import (

	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"

	"net/http"

	"github.com/go-chi/chi/v5"

)

func main() {
	router := chi.NewRouter()

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}
	
	router.Post("/campaigns", handler.CampaignsPost)
	http.ListenAndServe(":3000", router)
}
