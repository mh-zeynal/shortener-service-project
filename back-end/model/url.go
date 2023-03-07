package model

import (
	"gopkg.in/guregu/null.v3"
	_ "gopkg.in/guregu/null.v3"
	"time"
)

type URL struct {
	Id           int         `json:"id"`
	Short_url    string      `json:"shortUrl"`
	Original_url string      `json:"originalUrl"`
	Title        string      `json:"title"`
	Created_at   time.Time   `json:"createdAt"`
	Description  null.String `json:"description"`
	User_id      uint        `json:"userId"`
}
