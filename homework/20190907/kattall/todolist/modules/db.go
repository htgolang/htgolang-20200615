package modules

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "123123"
	dbHost     string = "localhost"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

var db *gorm.DB

func init(){
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil || db.DB().Ping() != nil {
		panic("Error Connect DB")
	}
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Task{})
}