package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type User struct {
	ID       int    `json:id`
	Name     string `json:name`
	Birthday string `json:birthday`
	Addr     string `json:addr`
	Tel      string `json:tel`
	Desc     string `json:desc`
	Password string `json:password`
}

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

func loadUsers() (map[int]User, error) {
	if bytes, err := ioutil.ReadFile("datas/user.json"); err != nil {
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
	return ioutil.WriteFile("datas/user.json", bytes, 0X066)
}

func GetUsers(q string) []User {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	rt_list := make([]User, 0)

	for _, user := range users {
		if q == "" || strings.Contains(user.Name, q) ||
			strings.Contains(user.Tel, q) || strings.Contains(user.Addr, q) ||
			strings.Contains(user.Desc, q) {
			rt_list = append(rt_list, user)
		}

	}
	return rt_list
}

func GetUserByName(name string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if name == user.Name {
			return user, nil
		}
	}
	log.Printf("Name == %s: Users Get Success", name)
	return User{}, errors.New("Not Found")
}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "Name长度需在4-12位之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "Name已存在"
	}

	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "password长度需在6-30位之间"
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

	user := User{
		ID:       id,
		Name:     name,
		Password: password,
		Birthday: birthday,
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

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if id == user.ID {
			return user, nil
		}
	}
	log.Printf("ID == %d: Users Get Success", id)
	return User{}, errors.New("Not Found")
}

func ModifyUser(id int, name, birthday, addr, tel, desc string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make(map[int]User, len(users))
	for i, user := range users {
		if id == user.ID {
			user.Name = name
			user.Desc = desc
			user.Birthday = birthday
			user.Addr = addr
			user.Tel = tel
		}
		newUsers[i] = user
	}
	storeUsers(newUsers)
	log.Printf("%s user Create Success", name)

}

func DeleteUser(id int) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	i := 0
	newUsers := make(map[int]User, 0)
	for _, user := range users {
		if id != user.ID {
			newUsers[i] = user
			i++
		} else {
			log.Printf("%s user will Delete", user)
		}
	}
	log.Printf("ID=%d user Delete Success", id)
	storeUsers(newUsers)
}

func ModifyPasswd(name, passwd string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make(map[int]User, len(users))
	for i, user := range users {
		if name == user.Name {
			user.Password = passwd
		}
		newUsers[i] = user
	}
	storeUsers(newUsers)
	log.Printf("Name为 %s 的Passwd Modidy Success", name)

}
