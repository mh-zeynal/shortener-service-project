package model

type UserLinks struct {
	Links   []URL `json:"links"`
	IsError bool  `json:"isError"`
}
