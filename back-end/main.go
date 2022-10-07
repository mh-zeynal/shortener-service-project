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
	api.GET("/admin", handlers.HandleAdminAccess)
	e.Logger.Info(e.Start(":9090"))
}