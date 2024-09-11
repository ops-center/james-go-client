package inbox

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

func (t *tokenService) CreateJwtTokenForObject(object Object, setAliasAsSubject bool) (string, error) {
	defaultExpTime := time.Now().Add(DefaultTokenExpDuration).Unix()

	var sub string
	var err error

	if setAliasAsSubject {
		sub, err = object.GetAddressAlias()
		if err != nil {
			return "", fmt.Errorf("failed to get address alias: %v", err)
		}

		sub = fmt.Sprintf("%s@%s", sub, GlobalMailDomain)
	} else {
		sub, err = generateObjectAddr(object)
		if err != nil {
			return "", fmt.Errorf("failed to generate object addr: %v", err)
		}
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

func GetAdminToken(privateKey string) (string, error) {
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", err
	}
	service := newTokenService(rsaPrivateKey)
	adminClaims := jwt.MapClaims{
		"sub":  "admin",
		"exp":  time.Now().Add(AdminTokenExpDuration).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  "ACE Platform",
		"type": "admin",
	}
	adminToken, err := service.CreateJwtTokenFromClaims(adminClaims)
	if err != nil {
		return "", err
	}

	return adminToken, nil
}
