package endpoints

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	var requestBody contract.NewCampaign

	render.DecodeJSON(request.Body, &requestBody)

	id, err := h.CampaignService.Create(requestBody)

	if err != nil {
		if errors.Is(err, internalerror.ErrInternal) {
			return nil, 500, internalerror.ErrInternal

		} else {
			return nil, 400, err
		}

	}
	return map[string]string{"id": id}, 201, nil

}
