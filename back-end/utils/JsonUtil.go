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

func ConvertResponseMessageToJson(message string, destination string) (*string, error) {
	newMessage := model.ResponseMessage{Message: message, DestinationLink: destination}
	b, err := json.Marshal(newMessage)
	if err != nil {
		return nil, err
	}
	temp := string(b)
	return &temp, nil
}
