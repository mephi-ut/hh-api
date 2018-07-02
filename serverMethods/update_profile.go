package serverMethods

import (
	"database/sql"
	"net/http"

	m "github.com/mephi-ut/hh-api/models"
	"github.com/mephi-ut/hh-api/serverMethods/helpers"
	"github.com/gin-gonic/gin"
)

type updateProfileParams struct {
	Nickname string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func UpdateProfile(c *gin.Context) {
	var json updateProfileParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if json.Nickname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nickname is not set"})
	}

	me := helpers.GetMe(c)
	user, err := m.User.First(m.UserF{Id: me.GetId()})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find the user"})
		c.Abort()
		return
	}
	user.Nickname = json.Nickname

	if json.Password != "" {
		passwordHash := m.HashPassword(json.Password)
		user.PasswordHash = passwordHash
	}
	if json.Email != "" {
		user.Email = &json.Email
	} else {
		user.Email = nil
	}
	if json.Phone != "" {
		user.Phone = &json.Phone
	} else {
		user.Phone = nil
	}
	err = user.Update()
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot update user profile (the nickname is already used?)"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
