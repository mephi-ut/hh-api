package serverMethods

import (
	"github.com/gin-gonic/gin"
)

func Null(c *gin.Context) {
	c.JSON(200, gin.H{})
}
