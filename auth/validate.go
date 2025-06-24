package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

const edDsaJWTHeader = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9"
const edDsaSigLen = 86 // base64url-encoded Ed25519 signature length

func ParseToken(token string, claims jwt.Claims, compact bool) (*jwt.Token, error) {
	if compact {
		var err error
		token, err = EdDsaJWTTokenCompactToFull(token)
		if err != nil {
			return nil, err
		}
	}
	return jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
	)
}

func EdDsaJWTTokenCompactToFull(compact string) (string, error) {
	if len(compact) <= edDsaSigLen {
		return "", errors.New("too short to be a valid compact JWT")
	}
	payload := compact[:len(compact)-edDsaSigLen]
	signature := compact[len(compact)-edDsaSigLen:]
	fullToken := edDsaJWTHeader + "." + payload + "." + signature
	return fullToken, nil
}
