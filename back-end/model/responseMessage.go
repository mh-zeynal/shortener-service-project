package model

type ResponseMessage struct {
	Message         string `json:"message"`
	DestinationLink string `json:"destinationLink"`
}
