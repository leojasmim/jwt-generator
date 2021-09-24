package domain

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_CreateClaimsByString(t *testing.T) {
	payload := `{ "message" : "teste_jwt_generator", "inner" : { "key" : "nome", "value" : "abcde" } }`

	result := createClaims(payload)

	assert.Equal(t, result["message"], "teste_jwt_generator")
}

func Test_CreateClaimsByMap(t *testing.T) {
	payload := make(map[string]interface{})
	payload["message"] = "teste_jwt_generator"

	result := createClaims(payload)

	assert.Equal(t, result["message"], "teste_jwt_generator")
}
