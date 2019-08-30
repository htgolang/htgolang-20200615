package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"usertask/utils"
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

func (u User) TrimSpace(txt string) string {
	return strings.TrimSpace(txt)
}

func (u User) VaildatePassword(password string) bool {
	return u.Password == utils.Md5(password)
}

func loadUser() (map[int]User, error) {
	bytes, err := ioutil.ReadFile("datas/user/users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	} else {
		var users map[int]User
		if err := json.Unmarshal(bytes, &users); err != nil {
			return nil, err
		} else {
			return users, nil
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
	users, err := loadUser()
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

func GetUserByName(name string) (User, bool) {
	users, err := loadUser()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		if user.Name == name {
			return user, true
		}
	}
	return User{}, false
}

func ValidateCreateUser(name, password, password2, birthday, tel, desc, addr string) map[string]string {
	fmt.Println("validateCreateUser:", name, password, password2, birthday, tel, desc, addr)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	if password != password2 {
		errors["password"] = "两次密码输入不一样"
	} else if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6-30之间"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}

	return errors
}

func GetUserkId() (int, error) {
	var id int
	users, err := loadUser()
	if err != nil {
		return -1, err
	}
	for _, user := range users {
		if id < user.Id {
			id = user.Id
		}
	}
	return id + 1, nil
}

func CreateUser(name, password, birthday, tel, desc, addr string) {
	id, err := GetUserkId()
	if err != nil {
		panic(err)
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	user := User{
		Id: id,
		Name: name,
		Password: utils.Md5(password),
		Birthday: bir,
		Tel: tel,
		Addr: addr,
		Desc: desc,
	}

	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	users[id] = user
	storeUser(users)
}

func GetUserById(id int) (User, bool) {
	fmt.Println("GetUserById:", id)
	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if user.Id == id {
			return user, false
		}
	}

	return User{}, true
}


func ValidateModifyUser(name string, birthday string) map[string]string {
	// name, birthday, tel, desc, addr
	fmt.Println("ValidateModifyUser:", name, birthday)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}
	return errors
}


func ModifyUser(id int, name, birthday, tel, desc, addr string) {
	users, err := loadUser()
	if err != nil {
		panic(err)
	}

	new_user := make(map[int]User, len(users))
	for i, user := range users {
		if user.Id == id {
			user.Name = name
			bir, _ := time.Parse("2006-01-02", birthday)
			user.Birthday = bir
			user.Tel = tel
			user.Desc = desc
			user.Addr = addr
		}
		new_user[i] = user
	}
	storeUser(new_user)
}

func DeleteUser(id int) {
	fmt.Println("delete user:", id)
	users, err :=loadUser()
	if err != nil {
		panic(err)
	}

	new_user := make(map[int]User, 0)
	for i, user := range users {
		if user.Id != id {
			new_user[i] = user
		}
	}
	storeUser(new_user)
}