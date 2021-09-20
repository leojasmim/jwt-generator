package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwtgen "github.com/leojasmim/jwt-generator/domain"
)

func GenerateJWT(c *gin.Context) {
	var payload = c.PostForm("payload")
	var sk = c.PostForm("sk")

	generator := getGenerator(c)

	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("an internal error occurred: %v", rec)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
	}()

	if payload == "" || sk == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "payload and sk is required"})
		return
	}

	token, _ := generator.Sign(payload, sk)

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func getGenerator(c *gin.Context) jwtgen.JwtGenerator {
	if g, exists := c.Get("generator"); exists {
		return g.(jwtgen.JwtGenerator)
	}
	return jwtgen.GetJwtGenerator("default")
}
