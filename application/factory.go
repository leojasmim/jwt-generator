package application

import (
	jwtgen "github.com/leojasmim/jwt-generator/domain"

	"github.com/gin-gonic/gin"
)

func GetJwtGenerator(c *gin.Context) {
	generator := jwtgen.GetJwtGenerator(c.PostForm("alg"))
	c.Set("generator", generator)

}
