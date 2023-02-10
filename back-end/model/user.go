package model

import (
	"time"
)

type User struct {
	User_Id    uint      `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
