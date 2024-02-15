package endpoints

import (
	"emailn/internal/contract"
	"net/http"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	var requestBody contract.NewCampaign

	render.DecodeJSON(request.Body, &requestBody)
	email := request.Context().Value("email").(string)
	requestBody.CreatedBy = email

	id, err := h.CampaignService.Create(requestBody)


	return map[string]string{"id": id}, 201, err

}
