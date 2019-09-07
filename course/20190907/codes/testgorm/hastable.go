package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "881019"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "testgorm"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

type User struct {
	gorm.Model
	Name     string
	Password string
	Birthday time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
}

type User2 struct {
	Name     string
	Password string
	Birthday time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
}

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(db.HasTable(&User{}))
	fmt.Println(db.HasTable("user2"))

	db.Close()
}
