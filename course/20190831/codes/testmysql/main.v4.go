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

	userName, userPassword := "xiangge11", "kk12"
	sql := "select name,password from todolist_user where name=? and password=md5(?)"
	// 操作 DQL DML
	// 查看 提交sql => 获取结果
	fmt.Println(sql)
	row := db.QueryRow(sql, userName, userPassword)

	var (
		name     string
		password string
	)

	err = row.Scan(&name, &password)
	fmt.Println(err)
	fmt.Println(name, password)
	db.Close()
}
