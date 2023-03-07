package main

import (
	"back-end/handlers"
	"back-end/middleWares"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api := e.Group("/api")
	api.POST("/login", handlers.LoginUser)
	api.POST("/signup", handlers.SignUpUser)
	api.GET("/:s", handlers.HandleRedirects, middleWares.AuthenticateUser)
	api.GET("/user_urls", handlers.GetUserUrls, middleWares.AuthenticateUser)
	api.POST("/shorten", handlers.HandleShortener, middleWares.AuthenticateUser)
	api.POST("/editUrl", handlers.EditUrlRecord, middleWares.AuthenticateUser)
	api.POST("/deleteUrl", handlers.DeleteUrl, middleWares.AuthenticateUser)
	e.Logger.Info(e.Start(":9090"))
}
