package james

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
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

func (t *tokenService) CreateJwtTokenFromClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (t *tokenService) GetUserTokenCookieName() string {
	return TokenCookieName
}

func (t *tokenService) GetUserTokenCookieDuration() time.Duration {
	return TokenCookieDuration
}

type AppCredentials struct {
	PrivateKey string
	AdminToken string
}

func GenerateAppCredentials() (*AppCredentials, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.New("failed to generate private key")
	}

	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	tokenService := newTokenService(privateKey)
	adminClaims := jwt.MapClaims{
		"sub":  "admin",
		"exp":  time.Now().Add(AdminTokenExpDuration).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  "ACE Platform",
		"type": "admin",
	}
	adminToken, err := tokenService.CreateJwtTokenFromClaims(adminClaims)
	if err != nil {
		return nil, err
	}

	return &AppCredentials{
		PrivateKey: string(pem.EncodeToMemory(&privateKeyPEM)),
		AdminToken: adminToken,
	}, nil
}
