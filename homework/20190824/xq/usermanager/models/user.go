package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type User struct {
	ID int
	Name string
	Brithday string
	Tel string
	Addr string
	Desc string
	Pass string
}


func md5Pass(pass string) string{
	ctx := md5.New()
	ctx.Write([]byte(pass))
	return hex.EncodeToString(ctx.Sum(nil))
}

func (u User) ValidatePassword(password string) bool {
	return md5Pass(password) == u.Pass
}
func loadUsers() ([]User, error) {

	if bytes, err := ioutil.ReadFile("datas/users.json"); err != nil{
		if os.IsNotExist(err){
			return []User{}, nil
		}
		return nil, err
	}else {
		var users []User
		if err:= json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		}else {
			return nil, err
		}
	}

}

func storeUsers(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {

		return err
	}
	return ioutil.WriteFile("datas/users.json", bytes, 0777)

}

func GetUsers(args ...string) []User {

	users, err := loadUsers()

	if err == nil {

		if args != nil {
			newUsers := make([]User, 0)

			for _, user := range users {

				for _, arg := range args {

					if strings.Contains(user.Name, arg) ||
						strings.Contains(user.Tel, arg) ||
						strings.Contains(user.Addr, arg) ||
						strings.Contains(user.Desc, arg) {
						newUsers = append(newUsers, user)
					}

				}
			}

			return newUsers

		}else {
			//fmt.Println(users)
			return users

		}
	}else {
		panic(err)
	}


}

func AutoUsers(name string) (*User, error) {

	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	for _,user := range users {
		if user.Name == name {
			return &user, nil
		}
	}

	return nil, nil

}

func GetUserId() (int, error)  {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for _, user := range users {
		if id < user.ID {
			id = user.ID
		}

	}
	return id +1 , nil
}

func ValidateCreateUser(name, brithday, tel, addr, desc, passwd string) map[string]string {

	errors := map[string]string{}
	if len(name) >12 || len(name) < 4 {
		errors["name"] = "名字长度在4-12之间"
	}else if err := GetUsers(name); err == nil {
		errors["name"] = "用户已存在"
	}

	if len(passwd) >20 || len(passwd) < 6 {
		errors["passwd"] = "密码长度在6-20之间"
	}

	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	}

	if len(addr) == 0 {
		errors["addr"] = "地址不能为空"
	}

	return errors

}



func CreateUsers(name, brithday, tel, addr, desc, passwd string){

	id, err := GetUserId()
	if err != nil {
		panic(err)
	}

	user := User{
		ID: id,
		Name: name,
		Brithday: brithday,
		Tel: tel,
		Addr: addr,
		Desc:desc,
		Pass:md5Pass(passwd),
	}
	users, err:= loadUsers()
	if err != nil {
		panic(err)
	}

	users = append(users, user)

	fmt.Println(user)

	storeUsers(users)
}

func GetUserById(id int) (User, error){

	users, err:= loadUsers()
	if err != nil {
		return User{}, err
	}

	for _, user := range users {
		if id == user.ID {
			return user, nil
		}
	}
	return User{}, errors.New("Not found")

}

func ModifyUsers(id int, name string, brithday ,tel, addr , desc, pass string )  {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	newUsers := make([]User, len(users))

	for i, user := range  users {
		if id == user.ID {
			user.Name = name
			user.Desc = desc
			user.Tel = tel
			user.Brithday = brithday
			user.Addr = addr
			user.Pass = pass
		}
		newUsers[i] = user
	}
	fmt.Println(newUsers)
	storeUsers(newUsers)
}

func ModifyUserPass(id int, name, brithday,tel, addr, desc, oldpass, newpass string) (error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	newUsers := make([]User, len(users))
	for i, user := range  users {
		if id == user.ID && md5Pass(oldpass) == user.Pass && newpass != " "{
			user.Name = name
			user.Brithday = brithday
			user.Tel = tel
			user.Addr = addr
			user.Desc = desc
			user.Pass = md5Pass(newpass)
		}
		newUsers[i] = user
	}
	fmt.Println(users)
	fmt.Println(newUsers)
	storeUsers(newUsers)
	return nil
}

func DeleteUsers(id int)  {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	//fmt.Println(id)

	newUsers := make([]User, 0)
	for _, user := range users {
		if id != user.ID {
			newUsers = append(newUsers, user)
		}else {
			fmt.Println(user)
		}
	}
	fmt.Println(storeUsers(newUsers))


}