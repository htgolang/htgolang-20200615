package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//添加用户
func adduser(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Println(id)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	fmt.Print("请输入年龄:")
	fmt.Scan(&age)

	fmt.Print("请输入联系方式:")
	fmt.Scan(&tel)

	fmt.Print("请输入家庭住址:")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"tel":  tel,
		"addr": addr,
	}

}

//查询用户
func query(users map[string]map[string]string) {
	var q string
	fmt.Print("请输入查询信息：")
	fmt.Scan(&q)
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%5s|%20s|%5s|%20s|%50s", user["id"], user["name"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}

//修改用户信息
func change(users map[string]map[string]string) {
	var (
		id           string
		name         string
		age          string
		tel          string
		addr         string
		change_input string
	)
	fmt.Print("请输入要修改用户ID：")
	fmt.Scan(&id)
	//fmt.Println(id)
	if user, ok := users[id]; ok {
		fmt.Printf("用户信息: %5s|%20s|%5s|%20s|%s", user["id"], user["name"], user["tel"], user["addr"])
		fmt.Println()
		fmt.Print("是否确认修改(Y/N):")
		fmt.Scan(&change_input)
		if change_input == "y" || change_input == "Y" {
			fmt.Print("请输入新的用户名：")
			fmt.Scan(&name)

			fmt.Print("请输入新的年龄：")
			fmt.Scan(&age)

			fmt.Print("请输入新的电话号码：")
			fmt.Scan(&tel)

			fmt.Print("请输入新的地址：")
			fmt.Scan(&addr)
			user["name"], user["age"], user["tel"], user["addr"] = name, age, tel, addr
			fmt.Println("用户信息修改成功!")

		}
	} else {
		fmt.Printf("您输入的用户ID,不存在,请重新输入!")
	}
}

//删除用户信息
func delete_user(users map[string]map[string]string) {
	var (
		id           string
		delete_input string
	)
	fmt.Print("请输入要删除用户ID：")
	fmt.Scan(&id)
	if user, ok := users[id]; ok {
		fmt.Printf("用户信息: %5s|%20s|%5s|%20s|%s \n", user["id"], user["name"], user["tel"], user["addr"])
		fmt.Print("是否确认删除(Y/N):")
		fmt.Scan(&delete_input)
		if delete_input == "y" || delete_input == "Y" {
			delete(users, id)
			fmt.Println("用户删除成功!")

		}
	} else {
		fmt.Printf("您输入的用户ID,不存在,请重新输入!")
	}
}

func main() {
	//存储用户信息
	var input_password string
	password := "1234567890"
	for i := 1; i <= 3; i++ {
		fmt.Print("请输入密码:")
		fmt.Scan(&input_password)
		if input_password == password {
			break
		}
		if i == 3 {
			fmt.Print("密码输入三次错误,程序退出")
			os.Exit(1)
		}

	}
	users := make(map[string]map[string]string)
	id := 0
	fmt.Println("==============欢迎使用用户管理系统==============")
	for {
		var op string
		fmt.Printf(`
		1. 新建用户
 		2. 修改用户
		3. 删除用户
 		4. 查询用户
		5. 退出
		请输入指令: `)
		fmt.Scan(&op)
		if op == "1" {
			id++
			adduser(id, users)
			fmt.Println(users)
		} else if op == "2" {
			change(users)
		} else if op == "3" {
			delete_user(users)
		} else if op == "4" {
			query(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("指令错误!!!")
		}
	}
}

/*
评分: 7.5
考虑：
1. add和update有一对堆用户信息输入，是否可以重构为一个函数
2. 思考if else-if else如何使用switch-case替代，更深入思考利用函数类型如何简写
*/
