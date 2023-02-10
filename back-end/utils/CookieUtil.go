package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func GenerateCookie(c echo.Context, cookieName string, cookieValue string) *http.Cookie {
	fmt.Println(time.Now())
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = cookieValue
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour * 30)
	return cookie
}
