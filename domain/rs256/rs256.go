package rs256

import (
	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
)

type RS256Generator struct {
	Algorithm jwt.SigningMethod
}

func New() RS256Generator {
	return RS256Generator{
		Algorithm: jwt.SigningMethodRS256,
	}
}

//SignRS256 Sign payload with asymmetric algorithm RSA Signature with SHA-256
func (j RS256Generator) Sign(payload string, privateKey string) (string, error) {

	claims := generateClaims(payload)

	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	token := jwt.NewWithClaims(j.Algorithm, claims)

	return token.SignedString(signKey)
}

func generateClaims(payload string) jwt.MapClaims {
	claimsMap := make(jwt.MapClaims)
	err := json.Unmarshal([]byte(payload), &claimsMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", claimsMap)
	return claimsMap
}

func dumpMap(space string, m map[string]interface{}) {
	for _, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			dumpMap(space+"\t", mv)
		}
	}
}
