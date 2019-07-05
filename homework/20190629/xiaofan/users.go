package main

import (
	"fmt"
	"os"
	"strings"
)

func printTitle() {
	title := fmt.Sprintf("%5s|%10s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

func printUser(user map[string]string) {
	fmt.Printf("%5s|%10s|%5s|%20s|%50s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
	fmt.Println()
}

// 添加用户
func add(pk int, users map[string]map[string]string) {
	var (
		id   = fmt.Sprintf("%d", pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	fmt.Print("请输入年龄:")
	fmt.Scan(&age)

	fmt.Print("请输入联系方式:")
	fmt.Scan(&tel)

	fmt.Print("请输入家庭地址:")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"tel":  tel,
		"age":  age,
		"addr": addr,
	}
	fmt.Println(id, name, age, tel, addr)
}

// 删除用户
func del(users map[string]map[string]string) {
	var id, verify string
	fmt.Print("请输入需要删除的用户ID：")
	fmt.Scan(&id)

	if user, ok := users[id]; ok {
		printTitle()
		printUser(user)

		fmt.Print("是否确认删除(Y/N):")
		fmt.Scan(&verify)
		if verify == "Y" || verify == "y" {
			delete(users, id)
			fmt.Println("删除成功")
		}
	} else {
		fmt.Println("你输入的用户ID不正确")
	}
}

// 修改用户
func change(users map[string]map[string]string) {
	var (
		id     string
		name   string
		age    string
		tel    string
		addr   string
		verify string
	)

	fmt.Print("请输入需要修改的用户ID：")
	fmt.Scan(&id)

	if user, ok := users[id]; ok {
		printTitle()
		printUser(user)

		fmt.Print("是否确认修改(Y/N):")
		fmt.Scan(&verify)
		if verify == "Y" || verify == "y" {
			fmt.Print("请输入要修改的姓名:")
			fmt.Scan(&name)

			fmt.Print("请输入要修改的年龄:")
			fmt.Scan(&age)

			fmt.Print("请输入要修改的联系方式:")
			fmt.Scan(&tel)

			fmt.Print("请输入要修改的家庭地址:")
			fmt.Scan(&addr)

			user["name"], user["age"], user["tel"], user["addr"] = name, age, tel, addr
			fmt.Println("修改成功")
			printTitle()
			printUser(user)
		}
	} else {
		fmt.Println("你输入的用户ID不正确")
	}
}

// 查询用户
// q == ""查找全部
// 非空, 名称，电话，住址中包含q内容的显示
func query(users map[string]map[string]string) {
	var q string

	fmt.Print("请输入查询信息：")
	fmt.Scan(&q)

	printTitle()

	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["addr"], q) {
			printUser(user)
		}
	}

}

func main() {
	var pw string
	password := "123abc!@#"

	for i := 1; i <= 3; i++ {
		fmt.Print("请输入密码:")
		fmt.Scan(&pw)
		if pw == password {
			break
		}
		if i == 3 {
			fmt.Println("密码错误3次，程序退出。")
			os.Exit(1)
		}
	}
	id := 0
	users := map[string]map[string]string{}
	fmt.Println("欢迎使用马哥用户系统")
	for {
		var op string
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入指令：`)
		fmt.Scan(&op)
		if op == "1" {
			id++
			add(id, users)
		} else if op == "2" {
			change(users)
		} else if op == "3" {
			del(users)
		} else if op == "4" {
			query(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("指令错误")
		}
	}

}

/*
评分: 7.5
考虑：
1. add和update有一对堆用户信息输入，是否可以重构为一个函数
2. 思考if else-if else如何使用switch-case替代，更深入思考利用函数类型如何简写
*/
