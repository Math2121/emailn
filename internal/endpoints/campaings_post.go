package endpoints

import (
	"emailn/internal/contract"
	internalerror "emailn/internal/internalError"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost(writer http.ResponseWriter, request *http.Request) {
		var requestBody contract.NewCampaign

		render.DecodeJSON(request.Body, &requestBody)

		id, err := h.CampaignService.Create(requestBody)

		if err != nil {
			if errors.Is(err, internalerror.ErrInternal) {
				render.Status(request, 400)
				render.JSON(writer, request, map[string]string{"error": err.Error()})
			} else {
				render.Status(request, 402)
				render.JSON(writer, request, map[string]string{"error": err.Error()})
			}

			return
		}

		render.Status(request, 201)
		render.JSON(writer, request, map[string]string{"id": id})
}