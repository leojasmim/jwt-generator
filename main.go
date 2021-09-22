package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leojasmim/jwt-generator/application"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())

	r.POST("/jwt", application.GenerateJWT)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
