package endpoints

import (
	"context"
	"emailn/internal/infrastructure/credentials"

	"github.com/go-chi/render"
	"net/http"
)

type ValidateTokenFunc func(token string, ctx context.Context) (string, error)

var ValidateToken ValidateTokenFunc = credentials.ValidateToken

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "tokenString missing"})
			return
		}

		email, err := ValidateToken(tokenString, r.Context())
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
