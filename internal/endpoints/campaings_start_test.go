package endpoints

import (
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_BY_ID_Should_return_campaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	campaignId := "test_id"

	service.On("Start", mock.Anything).Return(nil)

	handler := Handler{CampaignService: service}

	req, rr := newHttpTest("PATCH", "/", nil)
	addParameter(req, "id", campaignId)

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(200, status)
	assert.Nil(err)

}

func Test_CampaignStart_BY_ID_Should_return_campaign_error(t *testing.T) {
	setUp()
	assert := assert.New(t)

	errExpected := errors.New("something wrong")
	service.On("Start", mock.Anything).Return(errExpected)
	req, rr := newHttpTest("PATCH", "/", nil)

	_, _, err := handler.CampaignStart(rr, req)

	assert.Equal(errExpected, err)

}
