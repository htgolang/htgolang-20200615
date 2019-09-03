package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "user"
	dbPassword string = "ispasswd"
	dbHost     string = "172.25.50.250"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

// var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
// 	dbUser, dbPassword, dbHost, dbPort, dbName)

func connectDB() (*sql.DB, error) {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	err = db.Ping()
	return db, err
}
