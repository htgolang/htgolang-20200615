package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "881019"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "todolist2"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	sql := "insert into todolist_user(name, password, sex, birthday, tel, addr, `desc`, create_time) values(?, md5(?), ?, ?, ?, ?, ?, ?);"
	// 操作  DML
	//
	fmt.Println(sql)

	result, err := db.Exec(sql, "kk2", "kk", 1, "1988-11-12", "123", "123", "123", time.Now())

	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	db.Close()
}
