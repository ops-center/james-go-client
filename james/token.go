package james

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	privateKey *rsa.PrivateKey
}

func newTokenService(privateKey *rsa.PrivateKey) tokenService {
	return tokenService{
		privateKey: privateKey,
	}
}

func (t *tokenService) CreateJwtTokenForObject(object Object) (string, error) {
	defaultExpTime := time.Now().Add(DefaultTokenExpDuration).Unix()

	sub, err := generateObjectAddr(object)
	if err != nil {
		return "", fmt.Errorf("generate object addr: %v", err)
	}
	payload := jwt.MapClaims{
		"sub": sub,
		"exp": defaultExpTime,
	}
	additionalClaims := object.AdditionalTokenClaims()
	if additionalClaims != nil {
		for key, value := range *additionalClaims {
			payload[key] = value
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	signedToken, err := token.SignedString(t.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token, err: %v", err)
	}

	return signedToken, err
}

func (t *tokenService) GetUserTokenCookieName() string {
	return TokenCookieName
}

func (t *tokenService) GetUserTokenCookieDuration() time.Duration {
	return TokenCookieDuration
}
