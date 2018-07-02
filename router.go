package main

import (
	jwt "github.com/mephi-ut/hh-api/jwt"
	"github.com/mephi-ut/hh-api/cors"
	m "github.com/mephi-ut/hh-api/serverMethods"
	mw "github.com/mephi-ut/hh-api/serverMethods/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	setupJsonRouter(r)
	setupFrontendRouter(r)
}

func setupJsonRouter(r *gin.Engine) {
	// JWT
	jwtMiddleware := jwt.GetJwtMiddleware(
		func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	)
	corsed := r.Group("/")
	corsed.Use(cors.MiddlewareFunc())
	authed := r.Group("/")
	authed.Use(cors.MiddlewareFunc())
	authed.Use(jwtMiddleware.MiddlewareFunc()) // require to be authed
	authed.Use(mw.Authed)                      // some routines for an already authed
	corsed.POST("/auth.json", jwtMiddleware.LoginHandler)
	corsed.OPTIONS("/auth.json", m.Null)
	corsed.POST("/refresh_token.json", jwtMiddleware.RefreshHandler)
	corsed.OPTIONS("/refresh_token.json", m.Null)

	// Mix of my methods and JWT
	signUp := r.Group("/")
	signUp.Use(mw.SignUp)
	signUp.POST("/sign_up.json", jwtMiddleware.LoginHandler)

	// My methods
	r.GET("/ping.json", m.Ping)
	corsed.POST("/vacancies/:vacancy_name/test.json", m.PostTestForm)
	corsed.OPTIONS("/vacancies/:vacancy_name/test.json", m.Null)
	authed.GET("/profile.json", m.GetProfile)
	authed.PUT("/profile.json", m.UpdateProfile)
	corsed.OPTIONS("/profile.json", m.Null)
}

func setupFrontendRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	r.StaticFile("/login", "frontend/build/index.html")
	r.StaticFile("/signup", "frontend/build/index.html")
	for _, file := range []string{"index.html", "service-worker.js"} {
		r.StaticFile(file, "frontend/build/"+file)
	}
}
