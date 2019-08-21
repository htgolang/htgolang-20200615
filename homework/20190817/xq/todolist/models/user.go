package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func md5Pass(pass string) string{

	ctx := md5.New()
	ctx.Write([]byte(pass))
	return hex.EncodeToString(ctx.Sum(nil))
}

func AutoUsers(name, password string ) bool {

	users, err := loadUsers()

	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if user.Name == name && user.Pass == md5Pass(password) {
			return true
		}
	}
	return false
}

func GetUsers() []User {

	users, err := loadUsers()

	if err == nil {
		return users
	}else {
		panic(err)
	}

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

func CreateUsers(name, brithday, tel, addr, desc, passwd string) (error){
	id, err := GetUserId()

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
	for _, u := range users{
		if string(u.Name) == string(name) {
			return errors.New("用户已存在")
		}
	}
	users = append(users, user)
	storeUsers(users)
	return nil
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

func ReginsterUsers(name, brithday, tel, addr, desc, passwd string) (error){

	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	newuser := User{
		ID: id,
		Name: name,
		Brithday: brithday,
		Tel: tel,
		Addr: addr,
		Desc:desc,
		Pass:md5Pass(passwd),
	}
	for _, v := range users {
		if string(v.Name) == string(name) {
			return errors.New("用户已存在")
		}
	}
	users = append(users, newuser)
	storeUsers(users)
	return nil
}