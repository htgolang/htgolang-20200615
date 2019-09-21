package models

import (
	"time"

	"todolist/utils"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name       string    `gorm:"type:varchar(32); not null; default:'' "`
	Password   string    `gorm:"type:varchar(1024); not null; default:'' "`
	Birthday   time.Time `gorm:"type:date; not null"`
	Sex        bool      `gorm:"not null; default:false"`
	Tel        string    `gorm:"type:varchar(16); not null; default:''"`
	Addr       string    `gorm:"type:varchar(512); not null; default:''"`
	Desc       string    `gorm:"column:description; type:text; not null; default:''"`
	CreateTime time.Time `gorm:"column:create_time; type:datetime"`
}

func (u User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func GetUsers(q string) []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err
}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6~30之间"
	}
	return errors
}

func CreateUser(name, password, birthday, tel, addr, desc string) {

	day, _ := time.Parse("2006-01-02", birthday)

	user := User{
		Name:     name,
		Password: utils.Md5(password),
		Birthday: day,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	db.Create(&user)
}
