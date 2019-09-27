package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	Desc       string    `gorm:"column:description; type:text; not null"`
	CreateTime time.Time `gorm:"column:create_time; type:datetime"`
}

func (u User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func loadUsers() (map[int]User, error) {
	if bytes, err := ioutil.ReadFile("datas/users.json"); err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	} else {
		var users map[int]User
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}

func storeUsers(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/users.json", bytes, 0X066)
}

func GetUsers(q string) []User {
	var user []User
	db.Find(&user)
	return user
}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err
}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 3 {
		errors["name"] = "名称长度必须在3~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 4 {
		errors["password"] = "密码长度必须在4~30之间"
	}
	return errors
}

func ValidateModifyUser(name, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 3 {
		errors["name"] = "名称长度必须在3~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	return errors
}

func GetUserById(id int) User {
	var user User
	err := db.First(&user, "id=?", id).Error
	if err != nil {
		panic(err)
	}
	return user
}

func GetUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for uid := range users {
		if id < uid {
			id = uid
		}
	}
	return id + 1, nil
}

func CreateUser(name, password, birthday, tel, addr, desc string) {
	day, _ := time.Parse("2006-01-02", birthday)
	now := time.Now()

	user := User{
		Name:       name,
		Password:   utils.Md5(password),
		Birthday:   day,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		CreateTime: now,
	}
	db.Create(&user)
}

func ModifyUser(id int, name, tel, addr, desc string) {
	var user User
	if db.First(&user, "id=?", id).Error == nil {
		user.Name = name
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
