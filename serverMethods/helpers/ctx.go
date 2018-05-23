package helpers

import (
	m "github.com/mephi-ut/hh-api/models"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) m.UserI {
	me, ok := c.Get("me")
	if !ok {
		panic(`"me" is not set :(`)
	}
	return me.(m.UserI)
}
