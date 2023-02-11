package middleWares

import (
	"back-end/constants"
	"back-end/utils"
	"net/http"

	"github.com/labstack/echo"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("token")
		var msg *string
		if cookie == nil {
			msg, _ = utils.ConvertResponseMessageToJson(constants.UNAUTHORIZED_USER, "")
			return c.JSON(http.StatusUnauthorized, *msg)
		}
		_, err := utils.ExtractTokenClaimsFromCookie(*cookie)
		if err != nil {
			msg, _ = utils.ConvertResponseMessageToJson(constants.UNAUTHORIZED_USER, "")
			return c.JSON(http.StatusUnauthorized, *msg)
		}
		return next(c)
	}
}
