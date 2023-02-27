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
	var msg model.ResponseMessage
	var userRecord *model.User
	if err != nil {
		msg = utils.GenerateResponseMessage(err.Error(), "")
		return c.JSON(http.StatusBadRequest, msg)
	}
	if userRecord, _ = db.GetUserByUsername(user.Username); userRecord != nil {
		msg = utils.GenerateResponseMessage(constants.DUPLICATE_USERNAME, "")
		return c.JSON(http.StatusOK, msg)
	} else if userRecord, _ = db.GetUserByEmail(user.Email); userRecord != nil {
		msg = utils.GenerateResponseMessage(constants.DUPLICATE_EMAIL, "")
		return c.JSON(http.StatusOK, msg)
	} else if err == sql.ErrNoRows || userRecord == nil {
		db.InsertNewUser(*user)
	}
	token, _ := utils.GenerateToken(*user)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg = utils.GenerateResponseMessage(constants.SUCCESSFULL_SIGNUP, "")
	return c.JSON(http.StatusOK, msg)
}

func LoginUser(c echo.Context) error {
	user, err := utils.ConvertToLoginForm(c)
	var msg model.ResponseMessage
	if err != nil {
		msg = utils.GenerateResponseMessage(err.Error(), "")
		return c.JSON(http.StatusBadRequest, msg)
	}
	userRecord, err := db.GetUserByUsername(user.Username)
	if err == sql.ErrNoRows || userRecord == nil {
		msg = utils.GenerateResponseMessage(constants.NO_MATCH_USERNAME, "")
		return c.JSON(http.StatusOK, msg)
	}
	if user.Password != userRecord.Password {
		msg = utils.GenerateResponseMessage(constants.NO_MATCH_PASSWORD, "")
		return c.JSON(http.StatusOK, msg)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg = utils.GenerateResponseMessage(constants.SUCCESSFULL_LOGIN, "")
	return c.JSON(http.StatusOK, msg)
}
