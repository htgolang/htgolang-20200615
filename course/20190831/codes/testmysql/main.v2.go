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
	//操作 DQL DML
	//查询DQL
	rows, err := db.Query("select create_time, complete_time from todolist_task")

	var (
		id           int
		name         string
		progress     int
		status       int
		user         string
		desc         string
		completeTime time.Time
		createTime   time.Time
	)
	for rows.Next() {
		rows.Scan(&completeTime, &createTime)
		fmt.Println(id, name, progress, status, user, desc, completeTime, createTime)
	}

	db.Close()
}
