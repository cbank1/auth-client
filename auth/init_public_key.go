package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"log"
)

var publicKey ed25519.PublicKey

func InitPublicKey(publicKeyBase64 string) {
	pKeyBytes, err := base64.RawURLEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		log.Fatal(err)
	}
	publicKey = pKeyBytes
}

func GetPublicKey() ed25519.PublicKey {
	if publicKey == nil {
		log.Fatal("public key not initialized")
	}
	return publicKey
}
