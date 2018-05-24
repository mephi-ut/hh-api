package serverMethods

import (
	"github.com/mephi-ut/hh-api/serverMethods/helpers"
	"github.com/gin-gonic/gin"
)

func Whoami(c *gin.Context) {
	me := helpers.GetMe(c)

	c.JSON(200, gin.H{
		"UserId":   me.GetId(),
		"Nickname": me.GetNickname(),
		"Email":    me.GetEmail(),
	})
}
