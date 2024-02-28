package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJlbWFpbCI6InRlc3RlQGdtYWlsLmNvbSJ9.BkJ2TcT7_90yfHzGZkRoutb88Ibbj4QH1PUWn0IRvcc"

var validEmail string = "teste@gmail.com"

func Test_Auth_When_AuthorizationIsMissing_ReturnError(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})

	handlerFunc := Auth(nextHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "tokenString missing")

}

func Test_Auth_When_AuthorizationIsInvalid_ReturnError(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})
	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return "", errors.New("Invalid token")
	}
	handlerFunc := Auth(nextHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer testtoken")
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "invalid token")

}

func Test_Auth_When_AuthorizationIsValid_CallNextHandler(t *testing.T) {
	assert := assert.New(t)
	emailExpected := "teste@gmail.com"
	var email string

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email = r.Context().Value("email").(string)
	})

	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return emailExpected, nil
	}
	handlerFunc := Auth(nextHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", validToken)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Equal(validEmail, email)

}
