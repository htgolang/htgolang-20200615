package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	MaxAuth      = 3
	passwordFile = ".password"
)

// 设置序列化格式
var UsersInfo UserSerial = JsonFile{}

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func Auth() bool {
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil && len(password) > 0 {
		for i := 0; i < MaxAuth; i++ {
			fmt.Print("请输入密码:")
			input, _ := gopass.GetPasswd()

			if string(password) == fmt.Sprintf("%x", md5.Sum(input)) {
				return true
			} else {
				fmt.Println("[-]密码错误")
			}
		}
	} else {
		if len(password) == 0 || os.IsNotExist(err) {
			fmt.Print("请输入初始化密码:")
			bytes, _ := gopass.GetPasswd()
			_ = ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			fmt.Println("[-]发生错误", err)
			return false
		}
	}

	return false
}

func PrintUser(id int, user User) {
	fmt.Println("ID:", id)
	fmt.Println("名字:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006-01-02"))
	fmt.Println("联系方式:", user.Tel)
	fmt.Println("联系地址:", user.Addr)
	fmt.Println("备注:", user.Desc)
}

func Query() {
	q := InputString("请输入查询内容:")
	s := InputString("请输入排序属性(id/name/bir/tel)：")
	userSlice := make([]User, 0)
	users := UsersInfo.Load()
	for _, v := range users {
		userSlice = append(userSlice, v)
	}

	sort.Slice(userSlice, func(i, j int) bool {
		switch s {
		case "id":
			return userSlice[i].ID < userSlice[j].ID
		case "name":
			return userSlice[i].Name < userSlice[j].Name
		case "bir":
			return userSlice[i].Birthday.Before(userSlice[i].Birthday)
		case "tel":
			return userSlice[i].Tel < userSlice[j].Tel
		default:
			return userSlice[i].ID < userSlice[j].ID
		}

	})
	fmt.Println("================================")
	for _, v := range userSlice {
		//name, birthday, tel, addr, desc
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Desc, q) {
			PrintUser(v.ID, v)
			fmt.Println("---------------------------------")
		}
	}
	fmt.Println("================================")
}

func GetId() int {
	var id int
	users := UsersInfo.Load()

	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser(id int) User {
	var user User
	user.ID = id
	user.Name = InputString("请输入名字:")
	birthday, _ := time.Parse("2006-01-02", InputString("请输入出生日期(2000-01-01):"))
	user.Birthday = birthday
	user.Tel = InputString("请输入联系方式:")
	user.Addr = InputString("请输入联系地址:")
	user.Desc = InputString("请输入备注:")
	return user
}

func validUser(user User) error {
	if user.Name == "" {
		return fmt.Errorf("输入的用户名为空")
	}
	var load UserSerial = JsonFile{}
	users := load.Load()
	for _, tuser := range users {
		if user.Name == tuser.Name && user.ID != tuser.ID {
			return errors.New("输入的名字已经存在")
		}
	}
	return nil
}

func Add() {
	id := GetId()
	user := inputUser(id)
	if err := validUser(user); err == nil {
		users := UsersInfo.Load()
		users[id] = user
		UsersInfo.Store(users)
		fmt.Println("[+]添加成功")
	} else {
		fmt.Print("[-]添加失败:")
		fmt.Println(err)
	}
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		users := UsersInfo.Load()
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := validUser(user); err == nil {
					users[id] = user
					UsersInfo.Store(users)
					fmt.Println("[+]修改成功")
				} else {
					fmt.Print("[-]修改失败:")
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

func Del() {
	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		users := UsersInfo.Load()
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			fmt.Println(user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				UsersInfo.Store(users)
				fmt.Println("[+]删除成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}
func ModifyPassword() {
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil {
		// 验证密码
		fmt.Print("请输入密码:")
		bytes, _ := gopass.GetPasswd()
		if string(password) == fmt.Sprintf("%x", md5.Sum(bytes)) {
			fmt.Print("请输入新密码:")
			bytes, _ := gopass.GetPasswd()
			_ = ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			fmt.Println("[+]密码修改成功")
		} else {
			fmt.Println("[-]密码错误")
		}
	}
}
