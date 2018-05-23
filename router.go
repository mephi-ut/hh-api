package main

import (
	jwt "github.com/mephi-ut/hh-api/jwt"
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
	jwtMiddleware := jwt.GetJwtMiddleware()
	authed := r.Group("/")
	authed.Use(jwtMiddleware.MiddlewareFunc()) // require to be authed
	authed.Use(mw.Authed)                      // some routines for an already authed
	r.POST("/auth.json", jwtMiddleware.LoginHandler)
	r.POST("/refresh_token.json", jwtMiddleware.RefreshHandler)

	// Mix of my methods and JWT
	signUp := r.Group("/")
	signUp.Use(mw.SignUp)
	signUp.POST("/sign_up.json", jwtMiddleware.LoginHandler)

	// My methods
	r.GET("/ping.json", m.Ping)
	authed.GET("/whoami.json", m.Whoami)
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
