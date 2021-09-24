package domain

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type RS256Generator struct {
	algorithm jwt.SigningMethod
}

func NewRS256() RS256Generator {
	return RS256Generator{
		algorithm: jwt.SigningMethodRS256,
	}
}

//SignRS256 Sign payload with asymmetric algorithm RSA Signature with SHA-256
func (j RS256Generator) Sign(payload string, privateKey string) (string, error) {

	claims := createClaims(payload)

	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	token := jwt.NewWithClaims(j.algorithm, claims)

	return token.SignedString(signKey)
}
