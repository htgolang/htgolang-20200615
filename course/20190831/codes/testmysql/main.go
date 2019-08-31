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
	// 操作 DQL DML
	// 查看 提交sql => 获取结果
	rows, err := db.Query("select id,name,password,sex,birthday,tel,addr,`desc`,create_time from todolist_user")

	var (
		id              int
		name            string
		password        string
		sex             bool
		birthday        time.Time
		tel, addr, desc string
		createTime      time.Time
	)

	fmt.Println(rows, err)

	for rows.Next() {
		rows.Scan(&id, &name, &password, &sex, &birthday, &tel, &addr, &desc, &createTime)
		fmt.Println(id, name, password, sex, birthday, tel, addr, desc, createTime)
	}

	db.Close()
}
