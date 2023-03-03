package middleWares

import (
	"back-end/constants"
	"back-end/model"
	"back-end/utils"
	"net/http"

	"github.com/labstack/echo"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("token")
		var msg model.ResponseMessage
		if cookie == nil {
			msg = utils.GenerateResponseMessage(constants.UNAUTHORIZED_USER, "", false)
			return c.JSON(http.StatusUnauthorized, msg)
		}
		_, err := utils.ExtractTokenClaimsFromCookie(*cookie)
		if err != nil {
			msg = utils.GenerateResponseMessage(constants.UNAUTHORIZED_USER, "", false)
			return c.JSON(http.StatusUnauthorized, msg)
		}
		return next(c)
	}
}
