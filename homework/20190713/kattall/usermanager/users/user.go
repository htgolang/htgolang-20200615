package users

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	MaxAuth = 3
	// 密码 123456
	password = "e10adc3949ba59abbe56e057f20f883e"
)

var users map[int]User

type User struct {
	Id int
	Name string
	Birthday time.Time
	Tel string
	Addr string
	Desc string
}

func init() {
	// 默认初始化数据， 省的总要输入
	users = map[int]User{
		1: {1, "kk", time.Now().AddDate(1, 1, 1), "18777777777", "hangzhou", "kk"},
		2: {2, "ll", time.Now().AddDate(2,2,2), "18888888888", "shanghai", "ll"},
		3: {3, "kl", time.Now().AddDate(3,3,3), "18999999999", "beijing", "kl"},
		4: {4, "ko", time.Now().AddDate(4,4,4), "19000000000", "jiangxi", "ko"},
	}
}

func InputString(prompt string) (input string) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func Auth() bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码：")
		if inputpwd, err := gopass.GetPasswd(); err == nil {
			if fmt.Sprintf("%x", md5.Sum([]byte(inputpwd))) == password {
				return true
			} else {
				fmt.Println("密码错误.")
			}
		}
	}
	return false
}

func getUserID() (id int) {
	// 判断是否为空，如果为空，就返回1
	if len(users) == 0 {
		return 1
	}
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

// 输入用户信息, 并且对用户名进行检查
func inputUser() User {
	user := User{}

INPUTSTART:
	user.Name = InputString("请输入姓名：")
	// 对用户名进行检查, 输入不通过则跳转至LABEL INPUTSTART
	for _, u := range users {
		if user.Name == "" || user.Name == u.Name {
			fmt.Println("用户已存在, 请重新输入.")
			goto INPUTSTART
		}
	}
	user.Birthday, _ = time.Parse("2006/01/02 ", InputString("请输入出生日期(1991/01/01)："))
	user.Tel = InputString("请输入联系电话：")
	user.Addr = InputString("请输入联系地址：")
	user.Desc = InputString("请输入备注：")
	return user
}

func printUser(user User) {
	fmt.Println("ID:", user.Id)
	fmt.Println("名字:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
	fmt.Println("联系方式:", user.Tel)
	fmt.Println("联系地址:", user.Addr)
	fmt.Println("备注:", user.Desc)
}

// 排序
func sortUseLst(user_lst []User) []User {
	sortChoice := `请选择排序属性:
1. ID
2. 用户名
3. 出生日期
4. 联系方式
`
	s := InputString(sortChoice)
	if len(user_lst) == 1 {
		// 如果只是匹配一条数据，直接返回
		return user_lst
	} else {
		sort.Slice(user_lst, func(i, j int) bool {
			if s == "1" {
				return user_lst[i].Id < user_lst[j].Id
			} else if sortChoice  == "2" {
				return user_lst[i].Name < user_lst[j].Name
			} else if sortChoice == "3" {
				return user_lst[i].Birthday.Unix() < user_lst[j].Birthday.Unix()
			} else if sortChoice == "4" {
				return user_lst[i].Addr < user_lst[j].Addr
			} else {
				return false
			}
		})
		//fmt.Println(user_lst)
	}
	return user_lst
}


func Query() {
	user_lst := make([]User, 0)
	q := InputString("请输入查询的内容：")

	fmt.Println("=================================")
	for _, user := range users {
		if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) {
			user_lst = append(user_lst, user)
		}
	}

	//fmt.Println(user_lst)
	user_lst = sortUseLst(user_lst)
	for _, user := range user_lst {
		printUser(user)
	}
	fmt.Println("=================================")
}

func Add() {
	id := getUserID()
	user := inputUser()
	// 写入Id
	user.Id = id
	users[id] = user
	fmt.Println("[+]用户添加成功")
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改的ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改用户信息：")
			printUser(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser()
				// 写入Id
				user.Id = id
				users[id] = user
				fmt.Println("[+]修改成功")
			} else {
				fmt.Println("[-]用户ID不存在.")
			}
		} else {
			fmt.Println("[-]输入ID不正确")
		}

	}
}

func Del() {
	if id, err := strconv.Atoi(InputString("请输入删除的ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除用户信息：")
			printUser(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("[+]删除成功")
			} else {
				fmt.Println("[-]用户ID不存在.")
			}
		} else {
			fmt.Println("[-]输入ID不正确")
		}

	}
}
