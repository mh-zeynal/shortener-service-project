package model

type ResponseMessage struct {
	IsError         bool   `json:"isError"`
	Message         string `json:"message"`
	DestinationLink string `json:"link"`
}
