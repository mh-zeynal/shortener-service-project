package utils

import "back-end/model"

func GenerateResponseMessage(message string, destination string, isError bool) model.ResponseMessage {
	newMessage := model.ResponseMessage{IsError: isError, Message: message, DestinationLink: destination}
	return newMessage
}
