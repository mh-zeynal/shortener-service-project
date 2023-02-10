package handlers

import (
	"back-end/db"
	"back-end/utils"
	"database/sql"
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

func SignUpUser(c echo.Context) error {
	user, err := utils.ConvertBodyToUser(c)
	if err != nil {
		msg, _ := utils.ConvertResponseMessageToJson(err.Error(), "")
		return c.JSON(http.StatusBadRequest, msg)
	}
	userRecord, err := db.GetUserByUsername(user.Username)
	if err == sql.ErrNoRows || userRecord == nil {
		db.InsertNewUser(*user)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg, _ := utils.ConvertResponseMessageToJson("user signed in", "")
	return c.JSON(http.StatusOK, msg)
}

func LoginUser(c echo.Context) error {
	user, err := utils.ConvertBodyToUser(c)
	var msg *string
	if err != nil {
		msg, _ = utils.ConvertResponseMessageToJson(err.Error(), "")
		return c.JSON(http.StatusBadRequest, msg)
	}
	userRecord, err := db.GetUserByUsername(user.Username)
	if err == sql.ErrNoRows || userRecord == nil {
		err = errors.New("no user with this username")
		msg, _ = utils.ConvertResponseMessageToJson(err.Error(), "")
		return c.JSON(http.StatusUnauthorized, msg)
	}
	if user.Password != userRecord.Password {
		err = errors.New("no user with this username")
		msg, _ = utils.ConvertResponseMessageToJson(err.Error(), "")
		return c.JSON(http.StatusUnauthorized, msg)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg, _ = utils.ConvertResponseMessageToJson("user logged in successfully", "")
	return c.JSON(http.StatusOK, msg)
}
