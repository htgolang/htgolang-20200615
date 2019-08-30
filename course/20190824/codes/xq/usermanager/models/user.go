package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
			fmt.Println(newUsers)

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

func CreateUsers(name, brithday, tel, addr, desc, pass string) {

}