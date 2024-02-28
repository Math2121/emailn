package credentials

import (
	"context"
	"errors"
	"os"
	"strings"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/coreos/go-oidc/v3/oidc"
)

func ValidateToken(token string, ctx context.Context) (string, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	provider, err := oidc.NewProvider(ctx, os.Getenv("KEYCLOAK"))

	if err != nil {
		return "",errors.New(" error to conect to provider")

	}

	// verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
	_, err = verifier.Verify(ctx, token)

	if err != nil {
		return "", errors.New("Invalid token")
	}

	tokenJwt, _ := jwtgo.Parse(token, nil)
	claims := tokenJwt.Claims.(jwtgo.MapClaims)

	return claims["email"].(string), nil
}
