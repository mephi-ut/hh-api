package jwt

import (
	ginJwt "github.com/appleboy/gin-jwt"
	auth "github.com/mephi-ut/hh-api/auth"
	cfg "github.com/mephi-ut/hh-api/config"
	"github.com/gin-gonic/gin"
	"time"
)

func GetJwtMiddleware(unauthorizedFunc func(c *gin.Context, code int, message string)) *ginJwt.GinJWTMiddleware {
	return &ginJwt.GinJWTMiddleware{
		Realm:      "HH-IT-MEPHI",
		Key:        []byte(cfg.Get().Secret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(login string, password string, c *gin.Context) (string, bool) {
			return auth.SignIn(login, password)
		},
		Authorizator: func(login string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if unauthorizedFunc == nil {
				return
			}
			unauthorizedFunc(c, code, message)
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
