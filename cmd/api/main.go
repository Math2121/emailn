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

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	
	router.Post("/campaigns",endpoints.HandlerError(handler.CampaignsPost))
	router.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	
	http.ListenAndServe(":3000", router)
}
