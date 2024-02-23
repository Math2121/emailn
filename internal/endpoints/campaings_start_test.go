package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalmock "emailn/internal/test/internal-mock"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_BY_ID_Should_return_campaign(t *testing.T) {
	assert := assert.New(t)

	campaignId := "test_id"

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(nil)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("PATCH", "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(200, status)
	assert.Nil(err)

}

func Test_CampaignStart_BY_ID_Should_return_campaign_error(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("Start", mock.Anything).Return(errExpected)

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignStart(rr, req)

	assert.Equal(errExpected, err)

}
