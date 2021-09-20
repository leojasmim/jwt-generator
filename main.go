package main

import (
	"os"

	"github.com/gin-gonic/gin"
	app "github.com/leojasmim/jwt-generator/application"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/jwt", app.GenerateJWT)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
