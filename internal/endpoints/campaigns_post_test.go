package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/contract"
	inertnalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body contract.NewCampaign, createdBy string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdBy)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_Campaigns_Post_Should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	createdByExpected := "teste@email.com"

	body := contract.NewCampaign{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"test@example.com"},
	}

	service := new(inertnalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content && request.CreatedBy == createdByExpected {
			return true
		}

		return false
	})).Return("201", nil)

	handler := Handler{CampaignService: service}

	req, res := setup(body, createdByExpected)

	_, status, err := handler.CampaignsPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_Campaigns_Should_inform_error(t *testing.T) {
	assert := assert.New(t)
	createdByExpected := "teste@email.com"
	body := contract.NewCampaign{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"test@example.com"},
	}

	service := new(inertnalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error test"))

	handler := Handler{CampaignService: service}
	req, res := setup(body, createdByExpected)

	_, _, err := handler.CampaignsPost(res, req)

	assert.NotNil(err)

}
