package dialogflow

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/NYTimes/gizmo/auth/gcp"
)

type key int

const claimsKey key = 1

// GetUserClaims will return the Google identity claim set if it exists in the
// context. This can be used in coordination with the Authenticator.Middleware.
func GetUserClaims(ctx context.Context) (gcp.IdentityClaimSet, error) {
	var claims gcp.IdentityClaimSet
	clms := ctx.Value(claimsKey)
	if clms == nil {
		return claims, errors.New("claims not found")
	}
	return clms.(gcp.IdentityClaimSet), nil
}

func decodeClaims(token string) (gcp.IdentityClaimSet, error) {
	var claims gcp.IdentityClaimSet
	s := strings.Split(token, ".")
	if len(s) < 2 {
		return claims, errors.New("jws: invalid token received")
	}
	decoded, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return claims, err
	}
	err = json.Unmarshal(decoded, &claims)
	if err != nil {
		return claims, err
	}
	return claims, nil
}
