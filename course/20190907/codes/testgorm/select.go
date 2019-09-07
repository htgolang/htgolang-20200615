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
	db.LogMode(true)
	db.AutoMigrate(&User{})

	// for i := 0; i < 10; i++ {
	// 	user := User{
	// 		Name:     fmt.Sprintf("kk_%d", i),
	// 		Password: fmt.Sprintf("password_%d", i),
	// 		Birthday: time.Date(1988, 11, i, 0, 0, 0, 0, time.UTC),
	// 		Sex:      false,
	// 	}
	// 	db.Create(&user)
	// }

	var user User
	db.First(&user, "name = ?", "kk_2")
	fmt.Println(user)

	var user2 User
	db.Last(&user2, "name=?", "kk_5")
	fmt.Println(user2)

	var users []User
	// db.Where("name=?", "kk_6").Find(&users)

	// db.Where("name = ? and password = ?", "kk_5", "password_6").Find(&users)
	// db.Where("name = ?", "kk_5").Or("password = ?", "password_6").Find(&users)
	// db.Select([]string{"name", "password"}).Find(&users)

	db.Limit(3).Offset(5).Order("id desc, name asc").Find(&users)
	fmt.Println(users)

	var count int

	db.Model(&User{}).Where("name=?", "kk_2").Count(&count)
	db.Table("user").Where("name=?", "kk_2").Count(&count)
	fmt.Println(count)

	rows, _ := db.Model(&User{}).Select("name, password").Rows()

	for rows.Next() {
		var name, password string
		rows.Scan(&name, &password)
		fmt.Println(name, password)
	}

	rows, _ = db.Model(&User{}).Select("name, count(*) as cnt").Group("name").Having("count(*) > ?", 1).Rows()
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		fmt.Println(name, count)
	}

	db.Close()
}
