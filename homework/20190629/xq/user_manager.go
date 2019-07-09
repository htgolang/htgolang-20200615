/*

// 命令行的用户管理
// 存储到内存， 结构 可以用 映射
// 用户的信息 id name age tel addr (string)
*/
package main

import (
	"fmt"
	"strings"
	"github.com/howeyc/gopass"
)

// 用户信息输出格式
func outFormat(i string, users map[string]map[string]string) {
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s\n", "id", "name", "age", "tel", "addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", users[i]["id"], users[i]["name"], users[i]["age"], users[i]["tel"], users[i]["addr"])

}

// 检查用户是否存在
func checkUser(d string, users map[string]map[string]string) int {
	if _, ok := users[d]; !ok {
		return 0
	}
	return 1
}

// 添加 修改 删除 查询
func addUser(pk int, users map[string]map[string]string) {
	var (
		id string = fmt.Sprintf("%d", pk)
		// id string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)

	fmt.Println(id)
	fmt.Print("姓名: ")
	fmt.Scan(&name)
	fmt.Print("年龄: ")
	fmt.Scan(&age)
	fmt.Print("手机: ")
	fmt.Scan(&tel)
	fmt.Print("地址: ")
	fmt.Scan(&addr)

	//fmt.Println(id, name, age, tel, addr)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
	fmt.Println(users)
}

// 查询用户
// q 如果是空 查找全部
// 如果非空，必须 name, tel, age, addr
func selectUser(users map[string]map[string]string) {
	fmt.Print("输入查询关键字: ")
	var q string
	fmt.Scan(&q)

	for i, v := range users {

		if len(q) == 0 || strings.Contains(v["name"], q) ||
			strings.Contains(v["age"], q) || strings.Contains(v["tel"], q) ||
			strings.Contains(v["addr"], q) || strings.Contains(v["id"], q) {

			outFormat(i, users)
		}
	}

}

// 修改用户
//}
// 不在，提示不正确
// 在， 打印用户信息
// 提示用户是否确认修改（y/n）
// y: 提示用户输入修改后内容， name， age,xxx
//var d string

func modifyUser(users map[string]map[string]string) {

	var (
		d    string
		name string
		age  string
		tel  string
		addr string
		ch   string
	)
	fmt.Println("请输入需要修改的用户ID：")
	fmt.Scan(&d)

	if checkUser(d, users) == 1 {
		outFormat(d, users)
		fmt.Println("是否确定修改(y/n): ")
		fmt.Scan(&ch)
		if ch == "y" {
			fmt.Print("姓名: ")
			fmt.Scan(&name)
			fmt.Print("年龄: ")
			fmt.Scan(&age)
			fmt.Print("手机: ")
			fmt.Scan(&tel)
			fmt.Print("地址: ")
			fmt.Scan(&addr)

			fmt.Println()

			users[d] = map[string]string{
				"id":   d,
				"name": name,
				"age":  age,
				"tel":  tel,
				"addr": addr,
			}
			fmt.Println("修改后的内容为：")
			outFormat(d, users)

		} else if ch == "n" {
			fmt.Println("你已放弃修改。")
		} else {
			fmt.Println("非法输入，请重新输入。")
		}

	} else {
		fmt.Println("用户不存在！")
	}

}

// 删除用户
// 不在， 提示不争取
// 在，打印
// 提示用户是否删除（y/n）
// y 删除
//
func deleteUser(users map[string]map[string]string) {
	var d, ch string
	fmt.Println("请输入需要删除用户的ID：")
	fmt.Scan(&d)
	if checkUser(d, users) == 1 {
		outFormat(d, users)
		fmt.Println("是否确定删除(y/n): ")
		fmt.Scan(&ch)
		if ch == "y" {
			delete(users, d)
			fmt.Println("已删除！")
		} else if ch == "n" {
			fmt.Println("你已放弃删除！")
		} else {
			fmt.Println("非法输入，请重新输入。")
		}
	} else {
		fmt.Println("用户不存在！")
	}
}

// 用户管理系统界面
func initUser() {
	// 存储用户的信息
	users := make(map[string]map[string]string)
	id := 0

	// 第一种结构 if else
	//for {
	//	var op string
	//	fmt.Print(`请输入操作指令:
   //1：add
   //2: modify
   //3: delete
   //4: select
   //5: exit
   //      `)
	//	fmt.Scan(&op)
   //
	//	if op == "1" {
	//		id++
	//		addUser(id, users)
	//	} else if op == "2" {
	//		modifyUser(users)
	//	} else if op == "3" {
	//		deleteUser(users)
	//	} else if op == "4" {
	//		selectUser(users)
	//	} else if op == "5" {
	//		break
	//	} else {
	//		fmt.Println("指令错误!!!")
	//	}
   //
	//}

	// 第二种结构 switch case

	END:
	for {
		var op string
		fmt.Print(`请输入操作指令:
		1：add
		2: modify
		3: delete
		4: select
		5: exit
		     `)

		fmt.Scan(&op)
		switch op {
		case "1":
			id++
			addUser(id, users)
		case "2":
			modifyUser(users)
		case "3":
			deleteUser(users)
		case "4":
			selectUser(users)
		case "5":
			break END
		default:
			fmt.Println("指令错误！")
		}
	}
}

// 从命令行输入密码，并进行验证
// 通过返回值告知验证是否成功
const (
	maxAuth = 3
	pass = "123abc!@#"
)

func auth() bool {
	// 验证密码
	// 定义用户密码，让用户提示输入密码，passowrd=123abc!@#
	// 3 次失败，提示失败并退出
	// 成功，用户管理操作

	var p string

	for i := 0; i < maxAuth; i++ {

		fmt.Print(`你现在登陆的是用户管理系统，请输入密码: `)

		fmt.Scan(&p)
		if pass == p {
			fmt.Println("密码正确，已进入用户管理系统。")

			return true
		}else {
				fmt.Println("密码错误")
		}
	}

	return false
}

func main() {


	if !auth(){
		fmt.Printf("密码%d次错误， 程序退出\n", maxAuth)
		return
	}

	initUser()

}

/*
bug： 检索时 条件为空，不会显示全部用户信息
bug： 操作 5 退出后，外面密码验证依然会循环2次
*/

/*
评分: 7.5
考虑：
1. add和update有一对堆用户信息输入，是否可以重构为一个函数
2. 思考if else-if else如何使用switch-case替代，更深入思考利用函数类型如何简写
3. 减少不必要的函数定义, 如checkUser
*/
