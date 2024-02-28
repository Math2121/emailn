package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/domain/campaign"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupInternal(body campaign.NewCampaignRequest, createdBy string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdBy)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_Campaigns_Post_Should_save_new_campaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	createdByExpected := "teste@email.com"

	body := campaign.NewCampaignRequest{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"test@example.com"},
	}

	service.On("Create", mock.MatchedBy(func(request campaign.NewCampaignRequest) bool {
		if request.Name == body.Name && request.Content == body.Content && request.CreatedBy == createdByExpected {
			return true
		}

		return false
	})).Return("201", nil)

	req, rr := newHttpTest("POST", "/", body)
	req = addContext(req, "email", createdByExpected)

	_, status, err := handler.CampaignsPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_Campaigns_Should_inform_error(t *testing.T) {
	setUp()
	assert := assert.New(t)
	createdByExpected := "teste@email.com"
	body := campaign.NewCampaignRequest{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"test@example.com"},
	}

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error test"))

	req, rr := newHttpTest("POST", "/", body)
	req = addContext(req, "email", createdByExpected)

	_, _, err := handler.CampaignsPost(rr, req)

	assert.NotNil(err)

}
