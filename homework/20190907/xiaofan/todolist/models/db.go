package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     = "root"
	dbPassword = "abc@123"
	dbHost     = "139.224.36.94"
	dbPort     = 3306
	dbName     = "gorm"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil || db.DB().Ping() != nil {
		panic(err)
	}
	//db.LogMode(true)
	db.AutoMigrate(&User{}, &Task{})
}
