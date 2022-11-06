package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email 	 string `json:"email"`
	Name 	 string `json:"name"`
}

func (u User) TableName() string {
	return "service_users"
}