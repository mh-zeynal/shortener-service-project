package model

type URL struct {
	Id int       `json:"id"`
	Short string `json:"short"`
	Long string  `json:"long"`
	Date string  `json:"date"`
}

func (u URL) TableName() string {
	return "urls"
}


