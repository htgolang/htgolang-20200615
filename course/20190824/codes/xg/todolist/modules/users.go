package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"todolist/utils"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
	Tel		 string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
}

func (u User) VaildatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func loadUsers() (map[int]User, error) {
	bytes, err := ioutil.ReadFile("datas/user/users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, err
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

func storeUser(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/user/users.json", bytes, os.ModePerm)
}

func GetUsers(q string) []User {
	fmt.Println("q:", q)
	users, err := loadUsers()
	fmt.Println(users)
	if err != nil {
		panic(err)
	}
	rtList := make([]User, 0)

	if q != "" {
		for _, user := range users {
			if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) || strings.Contains(user.Tel, q) {
				rtList = append(rtList, user)
			}
		}
		return rtList
	} else {
		for _, user := range users {
			rtList = append(rtList, user)
		}
		return rtList
	}
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