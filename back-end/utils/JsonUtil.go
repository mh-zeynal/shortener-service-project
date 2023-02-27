package utils

import (
	"back-end/model"
	"encoding/json"
	"errors"
	"github.com/labstack/echo"
)

func ConvertBodyToUser(c echo.Context) (*model.User, error) {
	tempUser := model.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&tempUser)
	if err != nil {
		return nil, errors.New("empty request body")
	}
	return &tempUser, nil
}

func ConvertToLoginForm(c echo.Context) (*model.LoginForm, error) {
	tempForm := model.LoginForm{}
	err := json.NewDecoder(c.Request().Body).Decode(&tempForm)
	if err != nil {
		return nil, errors.New("empty request body")
	}
	return &tempForm, nil
}
