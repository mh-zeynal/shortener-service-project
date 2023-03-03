package model

import (
	"time"
)

type URL struct {
	Id           int       `json:"id"`
	Short_url    string    `json:"shortUrl"`
	Original_url string    `json:"originalUrl"`
	Title        string    `json:"title"`
	Created_at   time.Time `json:"createdAt"`
	User_id      uint      `json:"userId"`
}
