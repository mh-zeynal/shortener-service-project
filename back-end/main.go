package main

import (
	"back-end/handlers"
	"back-end/middlewares"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api := e.Group("/api")
	api.POST("/shorten", handlers.HandleShortener)
	api.GET("/:s", handlers.HandleRedirects)
	api.GET("/admin", handlers.HandleAdminAccess, middlewares.AuthenticateUser)
	api.POST("/signUp", handlers.SignUpUser)
	e.Logger.Info(e.Start(":9090"))
}