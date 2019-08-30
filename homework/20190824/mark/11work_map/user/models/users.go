package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"user/utils"
)

type UserDescribe struct {
	Id       int
	Name     string
	Addr     string
	Desc     string
	Tel      string
	Birthday time.Time
	Password string
}

func (u UserDescribe) ValidatePassword(password string) bool {
	fmt.Println("utils.Md5New(password):", utils.Md5New(password))
	fmt.Println("u.Passwd:", u.Password)
	return utils.Md5New(password) == u.Password
}
func loadUser() (map[int]UserDescribe, error) {
	if bytes, err := ioutil.ReadFile("datas/tasks.json"); err != nil {
		if os.IsNotExist(err) {
			return map[int]UserDescribe{}, nil
		}
		return nil, nil
	} else {
		var users map[int]UserDescribe
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}
func storeUser(users map[int]UserDescribe) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/tasks.json", bytes, 0X066)
}
func GetUsers(q string) []UserDescribe {
	users, err := loadUser()
	//fmt.Println(users)
	if err != nil {
		panic(err)
	}
	rtList := make([]UserDescribe, 0)
	for _, user := range users {
		if q == "" ||
			strings.Contains(user.Name, q) || // q不为0，或者
			strings.Contains(user.Tel, q) ||
			strings.Contains(user.Addr, q) ||
			strings.Contains(user.Desc, q) {
			rtList = append(rtList, user)
		}
	}
	return rtList
}
func GetUserByName(name string) (UserDescribe, error) {
	users, err := loadUser()
	if err != nil {
		return UserDescribe{}, err
	}
	for _, user := range users {
		if user.Name == name {
			return user, nil
		}
	}
	return UserDescribe{}, errors.New("Not Found")
}

//func GetUserByName(name string) (*UserDescribe){
//	users,err := loadUser()
//	if err != nil{
//		panic(err)
//	}
//	for _,user := range users {
//		if user.Name == name{
//			return &user
//		}
//	}
//	return nil
//}

//func GetUserByName(name string) (UserDescribe,bool){
//	users,err := loadUser()
//	if err != nil{
//		panic(err)
//	}
//	for _,user := range users {
//		if user.Name == name{
//			return user,true
//		}
//	}
//	return UserDescribe{},false
//}

func ValidateCreateUser(name, password string, birthday, tel, desc, addr string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6~30位之间"
	}
	_, err := strconv.ParseInt(tel, 11, 0)
	if err != nil {
		errors["tel"] = "请确定你输入的是11位手机号码或者7位数的电话号码"
	} else if len(tel) != 7 && len(tel) != 11 {
		fmt.Println(len(tel))
		errors["tel"] = "电话号码必须是7位或者11位"
	}
	return errors
}
func GetUserId() (int, error) {
	users, err := loadUser()
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
func CreateUser(name, password string, birthday, tel, addr, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	day, _ := time.Parse("2006-01-02", birthday)
	user := UserDescribe{
		Id:       id,
		Name:     name,
		Password: utils.Md5New(password),
		Birthday: day,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc, // name, password, birthday, tel, desc, addr
	}

	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	//users = append(users, user)
	users[id] = user
	fmt.Println("CreateUser:", users)
	storeUser(users)
}

// (id, name, birthday, tel, desc, addr, password)
// name, password, birthday, tel, desc, addr
func ModifyUser(id int, name, password, birthday, tel, addr, desc string) {
	day, _ := time.Parse("2006-01-02", birthday)
	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	fmt.Println("password:", password)
	password = utils.Md5New(password)
	fmt.Println("password:", password)
	newUsers := make(map[int]UserDescribe, len(users))
	for i, user := range users {
		if id == user.Id && password != "" {
			user.Name = name
			user.Password = password
			user.Birthday = day
			user.Tel = tel
			user.Addr = addr
			user.Desc = desc
		}
		newUsers[i] = user
		fmt.Println("password:", password)
		//users = append(users, user)
	}
	storeUser(newUsers)
}
func GetUserByID(id int) (UserDescribe, error) {
	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	//newUser := make([]UserDescribe, len(users))
	for i, user := range users {
		fmt.Println(user.Id, id)
		if user.Id == id {
			//newUser[i] = user
			//fmt.Println("%#v:", newUser)
			//fmt.Println(id)
			//fmt.Println("%#v:", user)
			fmt.Println(i, user)
			return user, nil
		}
		fmt.Println("%#v:", user)
	}
	return UserDescribe{}, errors.New("not found")
}
func DeleteUser(id int) {
	users, err := loadUser()
	if err != nil {
		panic(err)
	}
	newUser := make(map[int]UserDescribe, len(users))
	for i, user := range users {
		if id != user.Id {
			//user = newUser[id]
			newUser[i] = user
			//fmt.Println("%#v:", newUser)
		}
	}
	//fmt.Println("%#v:", newUser)
	storeUser(newUser)
}
