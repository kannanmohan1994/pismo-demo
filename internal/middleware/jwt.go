package middleware

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"xm/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

const (
	KeyIDKey = "kid"

	// tokens
	TokenIssuer  = "custom-issuer"
	TokenSubject = "custom_subject"
	TokenID      = "custom_token_id"

	// jwk
	JWKFilePathTemplate = "%v/%v"
	JWKFileName         = "jwk.json"
)

type JWTClaims struct {
	UserId string `json:"user_id"`
	Mobile string `json:"mobile"`
	jwt.RegisteredClaims
}

var jwkSet jwk.Set

func initJWKSet(cfg config.Config) {
	jwkFolderPath := cfg.JWKFilePath
	path := fmt.Sprintf(JWKFilePathTemplate, jwkFolderPath, JWKFileName)

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		jwkSet = nil
		return
	}

	var pbError error
	jwkSet, pbError = jwk.Parse(data)
	if pbError != nil {
		jwkSet = nil
		return
	}
}

func (m *middleware) CreateJWTTokenWithClaims(claims jwt.Claims, key jwk.Key, kid string) (string, error) {
	algorithm := key.Algorithm()
	siginingMethod := jwt.GetSigningMethod(algorithm)

	jwt_ := jwt.NewWithClaims(siginingMethod, claims)

	jwt_.Header["kid"] = kid

	var privateKey rsa.PrivateKey
	rErr := key.Raw(&privateKey)
	if rErr != nil {
		m.logger.Errorf("Error while creating raw key: ", rErr)
		return "", rErr
	}

	jwtSignedString, err := jwt_.SignedString(&privateKey)
	if err != nil {
		return "", err
	}

	return jwtSignedString, nil
}

// GetPublicKey parse the JWT with the cognito JWK
func (m *middleware) GetPublicKey(token *jwt.Token) (interface{}, error) {
	if jwkSet == nil {
		return nil, errors.New("jwkset is not loaded")
	}

	keyID, ok := token.Header[KeyIDKey].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	key, ok := jwkSet.LookupKeyID(keyID)
	if ok {
		var privateKey rsa.PrivateKey
		rErr := key.Raw(&privateKey)
		if rErr != nil {
			m.logger.Errorf("unable to find key", rErr)
			return nil, rErr
		}
		return privateKey.Public(), nil
	}
	return nil, fmt.Errorf("unable to find key %q", keyID)
}

func (m *middleware) GetKeyByKID(KID string) (jwk.Key, error) {
	key, ok := jwkSet.LookupKeyID(KID)
	if ok {
		return key, nil
	}
	return nil, fmt.Errorf("unable to find key %q", KID)
}
