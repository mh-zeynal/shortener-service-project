package model

import (
	"time"
)

type URL struct {
	Id           int       `json:"id"`
	Short_url    string    `json:"short_url"`
	Original_url string    `json:"original_url"`
	Title        string    `json:"title"`
	Created_at   time.Time `json:"created_at"`
	User_id      uint      `json:"user_id"`
}
