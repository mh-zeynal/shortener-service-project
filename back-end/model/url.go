package model

type URL struct {
	Id int       `json:"id"`
	Short string `json:"short"`
	Long string  `json:"long"`
	Date string  `json:"date"`
	description string `json:"description"`
	name string `json:"name"`
}

func (u URL) TableName() string {
	return "urls"
}


