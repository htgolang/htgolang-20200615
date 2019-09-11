package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"user/utils"

	//	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserDescribe struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(32);not null;default:''"`
	Addr       string     `gorm:"type:varchar(512);not null;default:''"`
	Desc       string     `gorm:"column:description;type:text;default:''"`
	Sex        string     `gorm:"not null;default:false"`
	Tel        string     `gorm:"type:varchar(16);not null;default:''"`
	Birthday   string     `gorm:"type:date;not null"`
	Password   string     `gorm:"type:varchar(1024);not null;default:''"`
	Createtime *time.Time `gorm:"type:datetime;not null"`
	ChangeTime *time.Time `gorm:"type:datetime;not null"`
}

func (u UserDescribe) ValidatePassword(password string) bool {
	fmt.Println("utils.Md5New(password):", utils.Md5New(password))
	fmt.Println("u.Passwd:", u.Password)
	return utils.Md5New(password) == u.Password
}

func GetUsers(q string) []UserDescribe {
	var user8 []UserDescribe
	db.Select([]string{"id,name,birthday,sex,tel,addr,description,createtime,change_time"}).Find(&user8)
	//db.Find(&user8)

	users := make([]UserDescribe, 0)
	for _, user := range user8 {
		if q == "" ||
			strings.Contains(user.Name, q) ||
			strings.Contains(user.Tel, q) ||
			strings.Contains(user.Addr, q) ||
			strings.Contains(user.Desc, q) {
			//fmt.Println("GetUsers:", user.Name, user.Tel, user.Addr, user.Desc)
			users = append(users, user)
		}
		//fmt.Println("GetUsers users:", users)
	}
	return users
}
func GetUserByName(name string) (UserDescribe, error) {
	var user UserDescribe
	err := db.First(&user, "name = ?", name).Error
	//fmt.Println(user, err)
	return user, err
}

func ValidateCreateUser(name, password string, birthday, sex, tel, desc, addr string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6~30位之间"
	}
	_, err := strconv.ParseInt(tel, 11, 0)
	if err != nil {
		errors["tel"] = "请确定你输入的是11位手机号码或者7位数的电话号码"
	} else if len(tel) != 7 && len(tel) != 11 {
		fmt.Println(len(tel))
		errors["tel"] = "电话号码必须是7位或者11位"
	}
	return errors
}

func CreateUser(name, password string, birthday, sex, tel, addr, desc string) {
	nows := time.Now()
	user := UserDescribe{
		Name:       name,
		Password:   utils.Md5New(password),
		Birthday:   birthday,
		Sex:        sex,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		Createtime: &nows,
		ChangeTime: &nows,
	}
	db.Create(&user)
}

// (id, name, birthday, tel, desc, addr, password)
// name, password, birthday, tel, desc, addr
func ModifyUser(id int, name, password, birthday, sex, tel, addr, desc string) {
	var user UserDescribe
	nows := time.Now()
	//fmt.Println("ModifyUser- ID :", id)
	if db.First(&user, "id=?", id).Error == nil {
		user.Name = name
		user.Password = utils.Md5New(password)
		user.Birthday = birthday
		user.Sex = sex
		user.Tel = tel
		user.Addr = addr
		user.Desc = desc
		user.ChangeTime = &nows
	}
	db.Save(&user)
}
func GetUserByID(id int) (UserDescribe, error) {
	var user UserDescribe
	err := db.First(&user, "id = ?", id).Error
	return user, err
}

func DeleteUser(id int) {
	var user UserDescribe
	if db.First(&user, "id=?", id).Error == nil {
		db.Delete(&user)
	}
}
