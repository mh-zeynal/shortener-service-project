package db

import (
	"back-end/utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MakeDbConnection() {
	var dbUser = utils.GetVariable("DB_User")
	var dbPass = utils.GetVariable("DB_Pass")
	//var dbCont = utils.GetVariable("DB_CONTAINER")
	//var dbContPort = utils.GetVariable("DB_CONTAINER_PORT")
	var dbSchema = utils.GetVariable("DB_SCHEMA")
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	dbUser, dbPass, dbCont, dbContPort, dbSchema)
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbSchema + "?parseTime=true"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("database connection has failed")
	}
}
