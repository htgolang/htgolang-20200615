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
	return "user2"
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
	// 	}
	// 	db.Create(&user)
	// }
	var user User
	db.First(&user)

	fmt.Println(user)
	user.Sex = false
	db.Save(&user)
	db.Model(&User{}).Where("id=2").UpdateColumn("sex", true)
	db.Model(&User{}).Where("id=3").UpdateColumns(map[string]interface{}{"sex": true})

	db.Model(&User{}).Where("id=4").Update("sex", true)
	db.Model(&User{}).Where("id=5").Updates(map[string]interface{}{"sex": true})

	db.Unscoped().Delete(&user)

	var users []User
	db.Find(&users)
	fmt.Println(users)
	db.Close()
}
