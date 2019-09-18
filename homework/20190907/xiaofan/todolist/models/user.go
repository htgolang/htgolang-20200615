package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(32); not null; default:''"`
	Password   string    `gorm:"type:varchar(32); not null; default:''"`
	Birthday   time.Time `gorm:"type:date; not null"`
	Tel        string    `grom:"type:varchar(16); not null; default:''"`
	Addr       string    `grom:"type:varchar(512); not null; default:''"`
	Desc       string    `gorm:"column:description; type:text; not null; default:''"`
	CreateTime time.Time `gorm:"column:create_time; type:datetime"`
	Sex        int
}

func GetUsers(query string) []User {
	var users []User
	if query == "" {
		db.Find(&users)
	} else {
		query = `%` + query + `%`
		fmt.Println(query)
		db.Where("name like ?", query).Or("addr like ?", query).Or("description like ?", query).Find(&users)
	}
	return users
}

func CreateUser(name, pwd, sex, bir, tel, addr, desc string) {
	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}

	sexI, err := strconv.Atoi(sex)
	if err != nil {
		fmt.Println(err)
	}

	db.Create(&User{
		Name:       name,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(pwd))),
		Birthday:   birParse,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		CreateTime: time.Now(),
		Sex:        sexI,
	})
}

func GetUserById(id int) (User, error) {
	var user User
	err := db.First(&user, "id=?", id).Error
	return user, err
}

func ModifyUser(id int, name, bir, tel, addr, desc string) {
	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}

	var user User
	if db.First(&user, "id=?", id).Error == nil {
		user.Name = name
		user.Birthday = birParse
		user.Tel = tel
		user.Addr = addr
		user.Desc = desc
	}
	db.Save(&user)
}

func DeleteUser(id int) {
	var user User
	if db.First(&user, "id=?", id).Error == nil {
		db.Delete(&user)
	}
}

func CheckUser(username, password string) bool {
	var user User
	if db.First(&user, "name=? and password=md5(?)", username, password).Error == nil {
		return true
	}
	return false
}

func ModifyPassword(name, password string) {
	var user User
	if db.First(&user, "name=?", name).Error == nil {
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	}
	db.Save(&user)
}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err
}

func ValidateCreateUser(name, password, bir, tel string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "用户名长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "用户名重复"
	}

	if len(password) > 30 || len(password) < 6 {
		fmt.Println(len(password))
		errors["password"] = "密码长度必须在6~30之间"
	}

	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}
	if birParse.Year() > 2019 || birParse.Year() < 1960 {
		errors["bir"] = "出生年必须在1960~2019之间"
	}

	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	} else if _, err := strconv.Atoi(tel); err != nil {
		errors["tel"] = "手机号码必须为数字"
	}

	return errors
}

func ValidateModifyUser(name, bir, tel string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "用户名长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "用户名重复"
	}

	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}
	if birParse.Year() > 2019 || birParse.Year() < 1960 {
		errors["bir"] = "出生年必须在1960~2019之间"
	}

	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	} else if _, err := strconv.Atoi(tel); err != nil {
		errors["tel"] = "手机号码必须为数字"
	}

	return errors
}
