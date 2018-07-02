package serverMethods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error": err.Error(),
	})
	c.Abort()
}

func jsonOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
