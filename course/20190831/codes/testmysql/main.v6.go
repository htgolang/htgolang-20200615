package main

import (
	"database/sql"
	"fmt"
	"os"

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

	sql := "update todolist_user set name=?, password=?, `desc`=? where id=?;"
	// 操作  DML
	//
	fmt.Println(sql)

	result, err := db.Exec(sql, "kk2", "kk2", "xxxxx", 31)

	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	db.Close()
}
