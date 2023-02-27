package utils

import "back-end/model"

func GenerateResponseMessage(message string, destination string) model.ResponseMessage {
	newMessage := model.ResponseMessage{Message: message, DestinationLink: destination}
	return newMessage
}
