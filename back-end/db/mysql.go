package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
	"back-end/model"
)

var DB *sql.DB

func MakeDbConnection() {
	var dbUser = utils.GetVariable("DB_User")
	//var dbPass = utils.GetVariable("DB_Pass")
	//var dbCont = utils.GetVariable("DB_CONTAINER")
	//var dbContPort = utils.GetVariable("DB_CONTAINER_PORT")
	//var dbSchema = utils.GetVariable("DB_SCHEMA")
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	dbUser, dbPass, dbCont, dbContPort, dbSchema)
	dsn := dbUser + ":" + "m@96@s97" + "@tcp(127.0.0.1:3306)/" + "gamedatabase?parseTime=true"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("database connection has failed")
	}
}

func AddNewRow(short string, long string)  {
	temp := model.URL{Short: short, Long: long, Date: time.Now().String()}
	_ = db.Create(&temp)
}

func DeleteRow(id int)  {
	_ = db.Delete(&model.URL{}, id)
}

func GetAllRows() *[]model.URL {
	rows := make([]model.URL, 0)
	_ = db.Find(&rows)
	return &rows
}

func GetRowViaLongUrl(long string) model.URL {
	temp := model.URL{}
	rows, _ := db.Model(&model.URL{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var user model.URL
		_ = db.ScanRows(rows, &user)
		if user.Long == long {
			temp = user
			break
		}
	}
	return temp
}

func GetRowViaShortUrl(short string) model.URL {
	temp := model.URL{}
	rows, _ := db.Model(&model.URL{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var user model.URL
		_ = db.ScanRows(rows, &user)
		if user.Short == short {
			temp = user
			break
		}
	}
	return temp
}

func IsLongAvailable(long string) bool {
	rows, _ := db.Model(&model.URL{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var user model.URL
		_ = db.ScanRows(rows, &user)
		if user.Long == long {
			return true
		}
	}
	return false
}

func IsShortAvailable(short string) bool {
	rows, _ := db.Model(&model.URL{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var user model.URL
		_ = db.ScanRows(rows, &user)
		if user.Short == short {
			return true
		}
	}
	return false
}