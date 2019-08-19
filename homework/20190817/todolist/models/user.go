package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func GetUsers() []User {
	users, err := loadUsers()
	if err == nil {
		return users
	}
	panic(err)
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
		Password: pwd,
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
			if user.Password == password {
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
			users[i].Password = password
		}
	}

	storeUsers(users)
}
