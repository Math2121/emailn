package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"

	"path/filepath"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(filepath.Join("./", ".env"))

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	db := database.NewDb()

	router.Use(endpoints.Auth)

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	router.Route("/api/v1/campaings", func(r chi.Router) {
		r.Post("/", endpoints.HandlerError(handler.CampaignsPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))

	})

	http.ListenAndServe(":3000", router)
}
