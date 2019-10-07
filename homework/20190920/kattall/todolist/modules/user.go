package modules

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
	"todolist/utils"
)

type User struct {
	gorm.Model
	Name        string     `gorm: "type: varchar(32); unique; not null; default: ''"`
	Password    string     `gorm: "type: varchar(64); not null; default: ''"`
	Birthday    *time.Time `gorm: "type: date"`
	Sex         bool       `gorm: "not null; default: false"`
	Tel         string     `gorm: "type: varchar(16); not null; default: ''"`
	Addr        string     `gorm: "type: varchar(1024); not null; default: ''"`
	Desc        string     `gorm: "type: text; column: description; not null; default: ''"`
	Create_time *time.Time `gorm: "column: create_time; type: datetime"`
}

func (u User) TableName() string {
	return "todolist_user"
}

func (u User) VaildatePassword(password string) bool {
	fmt.Println(u.Password, utils.Md5(password))
	return u.Password == utils.Md5(password)
}

func GetUsers(q string) []User {
	var users []User
	db.Find(&users)
	new_users := make([]User, 0)
	for _, user := range users {
		if q == "" || strings.Contains(user.Name, q) || strings.Contains(user.Tel, q) || strings.Contains(user.Desc, q) || strings.Contains(user.Addr, q) {
			new_users = append(new_users, user)
		}
	}
	return new_users
}

func GetUserByName(name string) (User, bool) {
	var user User
	if db.First(&user, "name = ?", name).Error == nil {
		return user, true
	} else {
		return user, false
	}
}

func ValidateCreateUser(name, password, password2, birthday, tel, desc, addr string) map[string]string {
	fmt.Println("validateCreateUser:", name, password, password2, birthday, tel, desc, addr)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	if password != password2 {
		errors["password"] = "两次密码输入不一样"
	} else if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6-30之间"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}

	return errors
}

func CreateUser(name, password, birthday string, sex bool, tel, desc, addr string) {
	bir, _ := time.Parse("2006-01-02", birthday)
	now := time.Now()
	user := User{
		Name:        name,
		Password:    utils.Md5(password),
		Birthday:    &bir,
		Sex:         sex,
		Tel:         tel,
		Addr:        addr,
		Desc:        desc,
		Create_time: &now,
	}
	db.Create(&user)
}

func ValidateModifyUser(name string, birthday string) map[string]string {
	fmt.Println("ValidateModifyUser:", name, birthday)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}
	return errors
}

func ModifyUser(id int, name, birthday, tel, desc, addr string, sexx bool) {
	var user User
	if db.First(&user, "id = ?", id).Error == nil {
		user.Name = name
		bir, _ := time.Parse("2006-01-02", birthday)
		user.Birthday = &bir
		user.Tel = tel
		user.Desc = desc
		user.Addr = addr
		user.Sex = sexx
	}
	db.Save(&user)
}

func DeleteUser(id int) {
	var user User
	if db.First(&user, "id = ?", id).Error == nil {
		db.Delete(&user)
	}
}

func GetUserById(id int) (User, error) {
	var user User
	err := db.Find(&user, "id = ?", id).Error
	return user, err
}
