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

	Name       string `gorm:"type:varchar(32); not null; default:''"`
	Password   string
	Birthday   time.Time `gorm:"type:date"`
	Sex        bool
	Tel        string `gorm:"column:telephone"`
	Addr       string
	Desciption string `gorm:"type:text"`
}

func (u *User) TableName() string {
	return "user3"
}

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// db.LogMode(true)
	db.AutoMigrate(&User{})

	var users []User
	db.Raw("select name from user3 where name=?", "kk_2").Scan(&users)
	fmt.Println(users)

	var name string
	db.Raw("select name from user3 where id=?", 2).Row().Scan(&name)
	fmt.Println(name)

	fmt.Println(db.Debug().Exec("insert into user3(name) values(?)", "kk4").RowsAffected)
	fmt.Println(db.Exec("delete from user3 where name = ?", "kk4").RowsAffected)

	db.Close()
}
