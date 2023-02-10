package middleWares

import (
	"back-end/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("token")
		if cookie == nil {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		claims, err := utils.ExtractTokenClaimsFromCookie(*cookie)
		if err == nil {
			fmt.Println(claims["usr"], claims["exp"])
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "unauthorized")
	}
}
