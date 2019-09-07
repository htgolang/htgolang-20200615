package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"time"

	"github.com/xlotz/todolist/utils"
)

type User struct {
	gorm.Model
	//Id         int
	Name       string `gorm:"type:varchar(32);not null;default:''"`
	Password   string `gorm:"type:varchar(1024);not null;default:''"`
	Birthday   time.Time `gorm:"type:date;not null"`
	Sex        bool	`gorm:"not null, default:false"`
	Tel        string `gorm:"type:varchar(32);not null;default:''"`
	Addr       string `gorm:"type:varchar(100);not null;default:''"`
	Descs       string `gorm:"type:varchar(512); not null;default:''"`
	CreateTime *time.Time `gorm:"type:datetime"`
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

func GetUserId(id int) (User, error) {

	var users User
	err := db.Find(&users, "id=?", id).Error
	return users, err
}

func CreateUser(name, password, birthday, tel, addr, desc string) {

	var now = time.Now()
	day, _ := time.Parse("2006-01-02", birthday)

	user := User{
		//ID:       id,
		Name:     name,
		Password: utils.Md5(password),
		Birthday: day,
		Tel:      tel,
		Addr:     addr,
		Descs:     desc,
		CreateTime: &now,
	}

	db.Create(&user)
}

func ModifyUserFromDB(id int, name, tel, addr, desc string){

	var users User

	//day, _ := time.Parse("2006-01-02", birthday)

	//fmt.Println(day)
	if db.First(&users, "id=?", id).Error == nil {
		users.Name = name
		//users.Birthday = day
		users.Addr = addr
		users.Tel = tel
		users.Descs = desc

		fmt.Println(&users)
		db.Save(&users)

		fmt.Println(db)

	}

}

func ModifyPassFromDB(id int, password string){

	var users User
	if db.First(&users, "id=?", id).Error == nil {
		users.Password = md5Pass(password)

		db.Save(&users)
	}

}
func DeleteUserFromDB(id int){

	var users User
	if db.Delete(&users, "id=?", id).Error == nil {
		db.Delete(&users)
	}

}


func md5Pass(pass string) string{
	ctx := md5.New()
	ctx.Write([]byte(pass))
	return hex.EncodeToString(ctx.Sum(nil))
}
