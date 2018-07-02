package serverMethodsMiddleware

import (
	"database/sql"
	"fmt"
	"net/http"

	m "github.com/mephi-ut/hh-api/models"
	"github.com/gin-gonic/gin"
)

func TryAuthed(c *gin.Context) error {
	nickname := c.GetString("userID");
	if nickname == "" {
		return fmt.Errorf(`nickname == ""`)
	}
	me, err := m.User.First(m.UserF{Nickname: nickname})
	if err == sql.ErrNoRows {
		return fmt.Errorf("User not found")
	}
	if err != nil {
		panic(err)
	}
	c.Set("me", me)
	return nil
}

func Authed(c *gin.Context) {
	err := TryAuthed(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}
