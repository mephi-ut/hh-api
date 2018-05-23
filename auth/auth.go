package auth

import (
	"database/sql"
	cfg "github.com/mephi-ut/hh-api/config"
	m "github.com/mephi-ut/hh-api/models"
)

func init() {
	cfg.AddReloadHook(Rehash)
}

func Rehash() {
}

func SignIn(login, password string) (loginFixed string, success bool) {
	user, err := m.User.First(m.UserF{Nickname: login})
	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		panic(err)
		return "", false
	}
	return login, user.CheckPassword([]byte(password))
}
