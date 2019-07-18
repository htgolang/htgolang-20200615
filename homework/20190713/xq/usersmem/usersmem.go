package usersmem

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"crypto/md5"
	"github.com/howeyc/gopass"
	"time"
)

const (
	MaxAuth  = 3
	password = "e10adc3949ba59abbe56e057f20f883e"
)
var users = map[int]User{}

type User struct {
	ID int
	Name string
	//Birthday string
	Birthday time.Time
	tel string
	addr string
	desc string
}

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

// 从命令行输入密码, 并进行验证
// 通过返回值告知验证成功还是失败
func Auth() bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码:")
		// fmt.Scan(&input)
		bytes, _ := gopass.GetPasswd()
		if password == fmt.Sprintf("%x", md5.Sum(bytes)) {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func printUser(user User) {
	fmt.Println("ID:", user.ID)
	fmt.Println("名字:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006-01-02"))
	fmt.Println("联系方式:", user.tel)
	fmt.Println("联系地址:", user.addr)
	fmt.Println("备注:", user.desc)
}

func sortUser(user_list []User) []User {
	var op string
	fmt.Print(`请输入排序字段:
	1: ID
	2: Name
	3: Birthday
	4: tel
`)

	fmt.Scan(&op)

	sort.Slice(user_list, func(i, j int) bool {
		switch op {
		case "1":
			return user_list[i].ID < user_list[j].ID
		case "2":
			return user_list[i].Name < user_list[j].Name
		case "3":
			return user_list[i].Birthday.Before(user_list[i].Birthday)
		case "4":
			return user_list[i].tel < user_list[j].tel
		default:
			return user_list[i].ID < user_list[j].ID
		}
	})
	return user_list
}

func Query() {
	tmp := make([]User, 0)
	q := InputString("请输入查询内容:")
	fmt.Println("================================")
	for _, v := range users {
		//name, birthday, tel, addr, desc
		if strings.Contains(v.Name, q) || strings.Contains(v.tel, q) || strings.Contains(v.addr, q) || strings.Contains(v.desc, q) {
			tmp = append(tmp, v)

		}
	}
	tmp = sortUser(tmp)
	for _, v := range tmp {
		printUser(v)
		fmt.Println("++++++++++++++++++++++++++")

	}
	fmt.Println("================================")
}

func getId() int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser() User {
	user := User{}
	END:
	user.Name = InputString("请输入名字:")
	for _, u :=range users{
		if user.Name == "" || user.Name == u.Name{
			fmt.Println("该用户已存在，请重新输入。")
			goto END
		}
	}
	user.Birthday, _ = time.Parse("2006-01-02",InputString("请输入出生日期(2006-01-02):"))
	user.tel = InputString("请输入联系方式:")
	user.addr = InputString("请输入联系地址:")
	user.desc = InputString("请输入备注:")
	return user
}

func Add() {
	id := getId()
	user := inputUser()
	user.ID = id
	users[id] = user
	fmt.Println("[+]添加成功")
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		for _, v := range users {
			if v.ID == id {
				fmt.Println("将修改的用户信息:")
				printUser(v)
				input := InputString("确定修改(Y/N)?")
				if input == "y" || input == "Y" {
					user := inputUser()
					users[id] = user
					fmt.Println("[+]修改成功")
				}
			}
			}

	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

func Delete() {

	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		for _, v := range users {
			fmt.Println(id)
			if v.ID == id {
				fmt.Println("将删除的用户信息:")
				printUser(v)
				input := InputString("确定删除(Y/N)?")
				if input == "y" || input == "Y" {
					delete(users, id)
					fmt.Println("[+]删除成功")
				}
			} else {
				fmt.Println("[-]用户ID不存在")
			}

			}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

/*
bug: 修改后 ID 会变成0
问题：本地环境 go mod 一直会提示 循环加载crypto, 待解决
*/