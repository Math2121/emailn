package endpoints

import (
	internalerror "emailn/internal/internalError"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_returns_internal_error(t *testing.T){
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 500, internalerror.ErrInternal
	}


	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET","/", nil)
	res := httptest.NewRecorder()


	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerror.ErrInternal.Error())


}

func Test_HandlerError_when_endpoint_returns_domain_error(t *testing.T){
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 500, errors.New("domain error")
	}


	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET","/", nil)
	res := httptest.NewRecorder()
	
	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}