package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
	Tel      string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u User) ValidatePassword(password string) bool {
	return u.Password == password
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

func GetUsers() map[int]User {
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
	for uid, _ := range users {
		if id < uid {
			id = uid
		}
	}
	return id + 1, nil
}

func CreateUser(name, password string, birthday time.Time, tel, addr, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	user := User{
		Id:       id,
		Name:     name,
		Birthday: birthday,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	user.SetPassword(password)
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users[id] = user
	storeUsers(users)
}

func GetUserByName(name string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New("Not Found")
}

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("Not Found")
}

func ModifyUser(id int, name string, birthday time.Time, tel, addr, desc string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	user, ok := users[id]
	if !ok {
		return
	}

	users[id] = User{
		Id:       id,
		Name:     name,
		Password: user.Password,
		Birthday: birthday,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	storeUsers(users)
}

func ModifyPassword(id int, password string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	user, ok := users[id]
	if !ok {
		return
	}
	user.SetPassword(password)

	users[id] = user
	storeUsers(users)
}

func DeleteUser(id int) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	delete(users, id)
	storeUsers(users)
}
