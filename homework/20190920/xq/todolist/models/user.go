package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xlotz/todolist/utils"

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
	Desc       string    `gorm:"column:description; type:text; not null"`
	CreateTime time.Time `gorm:"type:datetime; "`

}

func (u User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func GetUsers(q string) []User {
	users := []User{}

	if q == "" {
		db.Find(&users)
	}else {
		q = `%` + q + `%`
		db.Where("name like ?", q).Or("addr like ?", q).Find(&users)
	}


	return users

}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err
}
func GetUserById(id int) (User, error) {
	var user User
	err := db.First(&user, "id=?", id).Error
	return user, err
}

func ValidateModifyUser(name, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	} else if _, err := strconv.Atoi(tel); err != nil {
		errors["tel"] = "手机号码必须为数字"
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
		CreateTime: time.Now(),

	}
	fmt.Println(user)
	db.Save(&user)

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

	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	} else if _, err := strconv.Atoi(tel); err != nil {
		errors["tel"] = "手机号码必须为数字"
	}
	return errors
}

func ModifyUser(id int, name, tel, addr, desc string){

	var user User
	//day, _ := time.Parse("2006-01-02", birthday)

	if db.First(&user, "id = ?", id).Error == nil {
		user.Name = name
		//user.Birthday = day
		user.Tel = tel
		user.Desc = desc
		user.Addr = addr

	}

	db.Save(&user)

}

func DeleteUserFromDB(id int){

	var users User
	if db.Delete(&users, "id=?", id).Error == nil {

		fmt.Println(&users)
	}

}
