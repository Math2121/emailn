package endpoints

import (
	"context"

	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, 401)

			render.JSON(w, r, map[string]string{"error": "tokenString missing"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "Error creating provider"})
			return
		}

		// verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), tokenString)

		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "Invalid tokenString"})
			return
		}

		// decodificar token
		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)

		email := claims["email"]

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
