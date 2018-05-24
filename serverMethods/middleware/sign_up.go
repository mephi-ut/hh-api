package serverMethodsMiddleware

import (
	"bytes"
	m "github.com/mephi-ut/hh-api/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type signUpParams struct {
	Nickname string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ClosingBuffer struct {
	*bytes.Buffer
}
func (cb *ClosingBuffer) Close() (error) {
	return nil
}

func SignUp(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	c.Request.Body = &ClosingBuffer{bytes.NewBuffer(body)}

	var json signUpParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Request.Body = &ClosingBuffer{bytes.NewBuffer(body)}

	if json.Nickname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nickname is not set"})
	}

	passwordHash := m.HashPassword(json.Password)
	user := m.NewUser()
	user.Nickname = json.Nickname
	user.PasswordHash = passwordHash
	if json.Email != "" {
		user.Email = &json.Email
	}
	err = user.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create user (the nickname is already used?)"})
		c.Abort()
		return
	}

	c.Next()
}
