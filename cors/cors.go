package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MiddlewareFunc() func(*gin.Context) {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			switch origin {
			case "http://new.hh.it.mephi.ru:3000", "https://hh.it.mephi.ru", "https://hh.it.mephi.ru.":
				return true
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
