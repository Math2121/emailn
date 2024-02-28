package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"emailn/internal/domain/campaign"
	internalmock "emailn/internal/test/internal-mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Campaigns_Get_BY_ID_Should_return_campaign(t *testing.T) {
	assert := assert.New(t)

	campaingRes := campaign.CampaingResponse{
		ID:      "11",
		Name:    "Teste",
		Content: "Hi",
		Status:  "Allowed",
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaingRes, nil)

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaingRes.ID, response.(*campaign.CampaingResponse).ID)
	assert.Equal(campaingRes.Name, response.(*campaign.CampaingResponse).Name)
}

func Test_Campaigns_Get_BY_ID_Should_return_error(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)

	errorExpected := errors.New("something wrong")
	service.On("GetById", mock.Anything).Return(nil, errorExpected)

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errResponse := handler.CampaignGetById(rr, req)

	assert.Equal(errorExpected.Error(), errResponse.Error())
}
