package handlers

import (
	"back-end/constants"
	"back-end/db"
	"back-end/model"
	"back-end/utils"
	"database/sql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func SignUpUser(c echo.Context) error {
	user, err := utils.ConvertBodyToUser(c)
	var msg model.ResponseMessage
	var userRecord *model.User
	if err != nil {
		msg = utils.GenerateResponseMessage(constants.BAD_REQUEST, "", true)
		return c.JSON(http.StatusOK, msg)
	}
	if userRecord, _ = db.GetUserByUsername(user.Username); userRecord != nil {
		msg = utils.GenerateResponseMessage(constants.DUPLICATE_USERNAME, "", true)
		return c.JSON(http.StatusOK, msg)
	} else if userRecord, _ = db.GetUserByEmail(user.Email); userRecord != nil {
		msg = utils.GenerateResponseMessage(constants.DUPLICATE_EMAIL, "", true)
		return c.JSON(http.StatusOK, msg)
	} else if err == sql.ErrNoRows || userRecord == nil {
		db.InsertNewUser(*user)
	}
	token, _ := utils.GenerateToken(*user)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg = utils.GenerateResponseMessage(constants.SUCCESSFULL_SIGNUP, "", false)
	return c.JSON(http.StatusOK, msg)
}

func LoginUser(c echo.Context) error {
	user, err := utils.ConvertToLoginForm(c)
	var msg model.ResponseMessage
	if err != nil {
		msg = utils.GenerateResponseMessage(constants.BAD_REQUEST, "", true)
		return c.JSON(http.StatusOK, msg)
	}
	userRecord, err := db.GetUserByUsername(user.Username)
	if err == sql.ErrNoRows || userRecord == nil {
		msg = utils.GenerateResponseMessage(constants.NO_MATCH_USERNAME, "", true)
		return c.JSON(http.StatusOK, msg)
	}
	if user.Password != userRecord.Password {
		msg = utils.GenerateResponseMessage(constants.NO_MATCH_PASSWORD, "", true)
		return c.JSON(http.StatusOK, msg)
	}
	token, _ := utils.GenerateToken(*userRecord)
	c.SetCookie(utils.GenerateCookie(c, "token", token))
	msg = utils.GenerateResponseMessage(constants.SUCCESSFULL_LOGIN, "", false)
	return c.JSON(http.StatusOK, msg)
}

func SendUserBriefData(c echo.Context) error {
	token, _ := c.Cookie("token")
	var msg model.ResponseMessage
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mh_secret"), nil
	})
	if err != nil {
		msg = utils.GenerateResponseMessage(constants.UNAUTHORIZED_USER, "", true)
		return c.JSON(http.StatusUnauthorized, msg)
	}
	briefData, _ := db.GetUserBriefDataByUsername(claims["usr"].(string))
	return c.JSON(http.StatusOK, briefData)
}

func LogoutUser(c echo.Context) error {
	token, _ := c.Cookie("token")
	token.MaxAge = -99999
	token.Value = ""
	token.Path = "/"
	token.Expires = time.Unix(0, 0)
	token.HttpOnly = true
	c.SetCookie(token)
	msg := utils.GenerateResponseMessage(constants.SUCCESSFULL_LOGOUT, "", false)
	return c.JSON(http.StatusOK, msg)
}
