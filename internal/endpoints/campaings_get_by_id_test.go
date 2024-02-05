package endpoints

import (
	"emailn/internal/contract"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalmock "emailn/internal/test/internal-mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Campaigns_Get_BY_ID_Should_return_campaign(t *testing.T) {
	assert := assert.New(t)

	campaing := contract.CampaingResponse{
		ID:      "11",
		Name:    "Teste",
		Content: "Hi",
		Status:  "Allowed",
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaing, nil)

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaing.ID, response.(*contract.CampaingResponse).ID)
	assert.Equal(campaing.Name, response.(*contract.CampaingResponse).Name)
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
