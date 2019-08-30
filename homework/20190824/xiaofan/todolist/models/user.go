package models

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Tel      string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
	Password string    `json:"password"`
}

func loadUsers() ([]User, error) {

	if bytes, err := ioutil.ReadFile("datas/users.json"); err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	} else {
		var users []User
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}

}

func storeUsers(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/users.json", bytes, 0x066)
}

func GetUsers(query string) []User {
	users, err := loadUsers()
	if err == nil {
		newUser := make([]User, 0)
		for _, user := range users {
			if query == "" || strings.Contains(user.Name, query) || strings.Contains(user.Addr, query) ||
				strings.Contains(user.Desc, query) || strings.Contains(user.Tel, query) {
				newUser = append(newUser, user)
			}
		}
		return newUser
	}
	return []User{}
}

func GetUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for _, user := range users {
		if id < user.Id {
			id = user.Id
		}
	}
	return id + 1, nil
}

func CreateUser(name, pwd, bir, tel, addr, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}
	user := User{
		Id:       id,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(pwd))),
		Name:     name,
		Birthday: birParse,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users = append(users, user)
	storeUsers(users)
}

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if id == user.Id {
			return user, nil
		}

	}
	return User{}, errors.New("not found")
}

func ModifyUser(id int, name, bir, tel, addr, desc string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}

	for i, user := range users {
		if user.Id == id {
			users[i].Name = name
			users[i].Birthday = birParse
			users[i].Tel = tel
			users[i].Addr = addr
			users[i].Desc = desc

		}
	}
	storeUsers(users)
}

func DeleteUser(id int) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	newUser := make([]User, 0)
	for _, user := range users {
		if user.Id == id {
			continue
		}
		newUser = append(newUser, user)
	}

	storeUsers(newUser)
}

func CheckUser(username, password string) bool {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		if user.Name == username {
			if user.Password == fmt.Sprintf("%x", md5.Sum([]byte(password))) {
				return true
			}
			return false
		}
	}
	return false
}

func ModifyPassword(name, password string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	for i, user := range users {
		if user.Name == name {
			users[i].Password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
		}
	}
	storeUsers(users)
}

func GetUserByName(name string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		return User{}, err
	}

	for _, user := range users {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New("Not Found")
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
