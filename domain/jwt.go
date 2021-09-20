package domain

import "github.com/leojasmim/jwt-generator/domain/rs256"

type JwtGenerator interface {
	Sign(payload string, privateKey string) (string, error)
}

func GetJwtGenerator(alg string) JwtGenerator {
	switch alg {
	case "RS256":
		return rs256.New()
	default:
		return rs256.New()
	}
}
