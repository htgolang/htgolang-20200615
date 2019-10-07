package models

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name       string     `gorm:"type:varchar(32); not null; default:'' "`
	Birthday   time.Time  `gorm:"type:date; not null"`
	Sex        bool       `gorm:"not null; default:false"`
	Addr       string     `gorm:"type:varchar(512); not null; default:'' "`
	Tel        string     `gorm:"type:varchar(16); not null; default:'' "`
	Desc       string     `gorm:"column:description; varchar(1024); not null; default:'' "`
	Password   string     `gorm:"type:varchar(1024); not null; default:'' "`
	CreateTime *time.Time `gorm:"column:create_time; type:datetime "`
}

type User_str struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(32); not null; default:'' "`
	Birthday   string     `gorm:"type:date; not null"`
	Sex        bool       `gorm:"not null; default:false"`
	Addr       string     `gorm:"type:varchar(512); not null; default:'' "`
	Tel        string     `gorm:"type:varchar(16); not null; default:'' "`
	Desc       string     `gorm:"column:description; type:text; not null; default:'' "`
	Password   string     `gorm:"type:varchar(1024); not null; default:'' "`
	CreateTime *time.Time `gorm:"column:create_time; type:datetime "`
}

type Errors struct {
	Name           string
	Password       string
	PasswordVerify string
}

var Register bool
var Login_User string

func (u User) ValidatePassword(passwd string) bool {
	log.Printf("%s Account Verify Success", u.Name)
	return passwd == u.Password
}

func Init() {
	logfile := "user.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
	}

}

func GetUsers(q string) []User_str {
	var users []User
	if q == "" {
		db.Find(&users)
	} else {
		db.Where("name like ?", "%"+q+"%").Or("tel like ?", "%"+q+"%").Or("addr like ?", "%"+q+"%").Or("description like ?", "%"+q+"%").Find(&users)
	}

	users_str := make([]User_str, 0)
	var user_str User_str
	for _, user := range users {
		user_str.ID = user.ID
		user_str.Name = user.Name
		user_str.Sex = user.Sex
		user_str.Birthday = user.Birthday.Format("2006-01-02")
		user_str.Addr = user.Addr
		user_str.Tel = user.Tel
		user_str.Desc = user.Desc
		users_str = append(users_str, user_str)
	}

	return users_str
}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err

}

func ValidateCreateUser(name, password, passwordVerify string, nameexist bool) Errors {
	// 修改密码时，nameexist传入True参数，创建用户时，nameexist传入false

	//errors := map[string]string{}
	var errors Errors

	if nameexist {
		_, err := GetUserByName(name)
		if err != nil {
			errors.Name = "Name不存在"
		}
	} else {
		if len(name) > 12 || len(name) < 4 {
			errors.Name = "Name长度需在4-12位之间"
		} else if _, err := GetUserByName(name); err == nil {
			errors.Name = "Name已存在"
		}
	}

	if len(password) > 30 || len(password) < 6 {
		errors.Password = "password长度需在6-30位之间"
	}

	if password != passwordVerify {
		errors.PasswordVerify = "密码输入不一致"
	}

	return errors
}

func CreateUser(name, password, birthday, tel, addr, desc string, sex bool) {
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", birthday, local)
	var user User = User{
		Name:       name,
		Password:   password,
		Birthday:   t,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		Sex:        sex,
		CreateTime: &now,
	}

	if db.NewRecord(user) {
		db.Create(&user)
	}

}

func GetUserById(id int) (User_str, error) {
	var user User
	err := db.First(&user, "id=?", id).Error

	var user_str User_str
	user_str.ID = user.ID
	user_str.Name = user.Name
	user_str.Sex = user.Sex
	user_str.Birthday = user.Birthday.Format("2006-01-02")
	user_str.Addr = user.Addr
	user_str.Tel = user.Tel
	user_str.Desc = user.Desc

	return user_str, err

}

func ModifyUser(id int, name, birthday, addr, tel, desc string, sex bool) {
	var user User
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", birthday, local)
	if db.First(&user, "id=?", id).Error == nil {
		user.Name = name
		user.Birthday = t
		user.Addr = addr
		user.Tel = tel
		user.Desc = desc
		user.Sex = sex

		db.Save(&user)

	}
	log.Printf("%s user Create Success", name)

}

func DeleteUser(id int) {
	var user User
	if db.First(&user, "id=?", id).Error == nil {
		db.Delete(&user)
	}

	log.Printf("ID=%d user Delete Success", id)

}

func ModifyPasswd(name, passwd string) {
	var user User
	var passwd_md5 string
	passwd_md5 = fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

	if db.First(&user, "name=?", name).Error == nil {
		user.Password = passwd_md5
		db.Save(&user)
	}
	log.Printf("Name为 %s 的Passwd Modidy Success", name)
}
