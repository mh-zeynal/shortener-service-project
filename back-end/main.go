package main

import (
	"back-end/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api := e.Group("/api")
	api.POST("/shorten", handlers.HandleShortener)
	api.GET("/:s", handlers.HandleRedirects)
	api.POST("/signUp", handlers.SignUpUser)
	api.POST("/login", handlers.LoginUser)
	api.GET("/user_urls", handlers.GetUserUrls)
	e.Logger.Info(e.Start(":9090"))
}
