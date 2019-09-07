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
	Id         int    `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(32); not null; default:''"`
	Password   string
	Birthday   time.Time `gorm:"type:date"`
	Sex        bool
	Tel        string `gorm:"column:telephone"`
	Addr       string
	Desciption string `gorm:"type:text"`
}

func (u *User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	db.AutoMigrate(&User{})

	user := User{
		Name:     "kk2",
		Password: "123456",
		Birthday: time.Date(1988, 11, 11, 0, 0, 0, 0, time.UTC),
		Sex:      false,
	}

	fmt.Println(db.NewRecord(user))
	fmt.Println(user)
	db.Create(&user)

	fmt.Println(user)
	fmt.Println(db.NewRecord(user))
	if db.NewRecord(user) {

		db.Create(&user)
	}
	db.Close()
}
