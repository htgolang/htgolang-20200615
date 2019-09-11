package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "user"
	dbPassword string = "ispasswd"
	dbHost     string = "172.25.50.250"
	dbPort     int    = 3306
	dbName     string = "todolist2"
	// create database todolist2 default charset uft8mb4;
	// insert into user_describes(created_at,updated_at,name,password,birthday,createtime) values(now(),now(),'mark',md5('123'),now(),now());
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil || db.DB().Ping() != nil {
		panic("Error Connect DB")
	}
	db.AutoMigrate(&UserDescribe{})
}
