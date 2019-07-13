package users

import (
	"fmt"
	"strconv"
	"strings"

	"crypto/md5"

	"github.com/howeyc/gopass"
)

const (
	MaxAuth  = 3
	password = "1fdb7184e697ab9355a3f1438ddc6ef9"
)

var users map[int]map[string]string

func init() {
	users = make(map[int]map[string]string)
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

func printUser(pk int, user map[string]string) {
	fmt.Println("ID:", pk)
	fmt.Println("名字:", user["name"])
	fmt.Println("出生日期:", user["birthday"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("联系地址:", user["addr"])
	fmt.Println("备注:", user["desc"])
}

func Query() {
	q := InputString("请输入查询内容:")
	fmt.Println("================================")
	for k, v := range users {
		//name, birthday, tel, addr, desc
		if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["addr"], q) || strings.Contains(v["desc"], q) {
			printUser(k, v)
			fmt.Println("---------------------------------")
		}
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

func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = InputString("请输入名字:")
	user["birthday"] = InputString("请输入出生日期(2000-01-01):")
	user["tel"] = InputString("请输入联系方式:")
	user["addr"] = InputString("请输入联系地址:")
	user["desc"] = InputString("请输入备注:")
	return user
}

func Add() {
	id := getId()
	user := inputUser()
	users[id] = user
	fmt.Println("[+]添加成功")
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			printUser(id, user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Println("[+]修改成功")
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
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			printUser(id, user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("[+]删除成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}
