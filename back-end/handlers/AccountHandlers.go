package handlers

import (
	"back-end/constants"
	"back-end/db"
	"back-end/model"
	"back-end/utils"
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

func SignUpUser(c echo.Context) error {
	user, err := utils.ConvertBodyToUser(c)
	var msg *string
	var userRecord *model.User
	if err != nil {
		msg, _ = utils.ConvertResponseMessageToJson(err.Error(), "")
		return c.JSON(http.StatusBadRequest, msg)
	}
	if userRecord, _ = db.GetUserByUsername(user.Username); userRecord != nil {
		msg, _ = utils.ConvertResponseMessageToJson(constants.DUPLICATE_USERNAME, "")
		return c.JSON(http.StatusConflict, msg)
	} else if userRecord, _ = db.GetUserByEmail(user.Email); userRecord != nil {
		msg, _ = utils.ConvertResponseMessageToJson(constants.DUPLICATE_EMAIL, "")
		return c.JSON(http.StatusConflict, msg)
	} else if err == sql.ErrNoRows || userRecord == nil {
		db.InsertNewUser(*user)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg, _ = utils.ConvertResponseMessageToJson(constants.SUCCESSFULL_SIGNUP, "")
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
		msg, _ = utils.ConvertResponseMessageToJson(constants.NO_MATCH_USERNAME, "")
		return c.JSON(http.StatusUnauthorized, msg)
	}
	if user.Password != userRecord.Password {
		msg, _ = utils.ConvertResponseMessageToJson(constants.NO_MATCH_PASSWORD, "")
		return c.JSON(http.StatusUnauthorized, msg)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg, _ = utils.ConvertResponseMessageToJson(constants.SUCCESSFULL_LOGIN, "")
	return c.JSON(http.StatusOK, msg)
}
