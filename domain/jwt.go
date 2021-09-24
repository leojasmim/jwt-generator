package domain

import (
	"encoding/json"

	"github.com/golang-jwt/jwt"
)

type JwtGenerator interface {
	Sign(payload string, privateKey string) (string, error)
}

func GetJwtGenerator(alg string) JwtGenerator {
	switch alg {
	case "RS256":
		return NewRS256()
	default:
		return NewRS256()
	}
}

func createClaims(payload string) jwt.MapClaims {
	claimsMap := make(jwt.MapClaims)
	err := json.Unmarshal([]byte(payload), &claimsMap)
	if err != nil {
		panic(err)
	}
	dump(claimsMap)
	return claimsMap
}

func dump(m map[string]interface{}) {
	for _, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			dump(mv)
		}
	}
}
