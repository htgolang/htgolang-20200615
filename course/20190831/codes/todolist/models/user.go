package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"todolist/utils"
)

type User struct {
	Id         int
	Name       string
	Password   string
	Birthday   time.Time
	Sex        bool
	Tel        string
	Addr       string
	Desc       string
	CreateTime time.Time
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
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	rtList := make([]User, 0)
	for _, user := range users {
		if q == "" || strings.Contains(user.Name, q) ||
			strings.Contains(user.Tel, q) || strings.Contains(user.Addr, q) ||
			strings.Contains(user.Desc, q) {
			rtList = append(rtList, user)
		}
	}
	return rtList
}

func GetUserByName(name string) (User, error) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user, err
	}

	if err := db.Ping(); err != nil {
		return user, err
	}
	defer db.Close()

	row := db.QueryRow("select id,name,password,birthday,sex,tel,addr,create_time from todolist_user where name=?", name)
	err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.CreateTime)
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
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}

	day, _ := time.Parse("2006-01-02", birthday)

	user := User{
		Id:       id,
		Name:     name,
		Password: utils.Md5(password),
		Birthday: day,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users[id] = user
	storeUsers(users)
}
