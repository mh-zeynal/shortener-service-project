package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
	"url/internal/model"
)

var db *gorm.DB

func IsConnectionStablished() bool {
	return db != nil
}

func MakeDatabase(user string, pass string, dbname string)  {
	dsn := user + ":" + pass + "@/" + dbname
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
