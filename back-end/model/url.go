package model

type URL struct {
	Id    int     `json:"id"`
	Short string  `json:"short"`
	Long  string  `json:"long"`
	Date  string  `json:"date"`
	Name  string  `json:"name"`
	User_id int `json:"user_id"`
	User User `json:"user"`
}

func (u URL) TableName() string {
	return "urls"
}


