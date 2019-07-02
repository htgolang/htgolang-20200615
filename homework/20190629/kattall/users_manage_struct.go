package main

import (
	"fmt"
	"strings"
)

// 认证函数
func auth(passwd string) bool {
	var (
		inputpwd string
		count    int = 0
		max      int = 3
	)

	for {
		fmt.Print("请输入密码：")
		fmt.Scan(&inputpwd)
		if inputpwd != passwd {
			count++
			if count >= max {
				fmt.Println("密码错误, 退出系统。")
				return false
			}
		} else {
			return true
		}
	}
}

// 用户输入信息, 单独拎出来
func scanData() (name, age, tel, addr string) {
	fmt.Println()
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入电话号码：")
	fmt.Scan(&tel)
	fmt.Print("请输入家庭住址：")
	fmt.Scan(&addr)
	fmt.Println()

	return
}

// 用户选择
func scanFunc() (str string) {
	fmt.Scan(&str)
	fmt.Println()
	return str
}

// 用户输入id
func scanId() (id int) {
	fmt.Scan(&id)
	fmt.Println()
	return id
}

// 判断users是否为空, 为空就不让查询，删除等操作。 返回 users长度和是否为空
func isEmptyUser(users []user) (flag bool) {
	if len(users) == 0 {
		return true
	}
	return false
}

// 打印头信息
func titleString() {
	title := fmt.Sprintf("%5s|%20s|%5s|%10s|%50s", "ID", "NAME", "AGE", "TEL", "ADDR")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

// 通过id, 查询用户信息. 打印信息
func queryId(id int, users []user) {
	for _, u := range users {
		if id == u.id {
			formatUser(u)
		}
	}
}

// 传入user, 输出打印信息
func formatUser(user user) {
	fmt.Printf("%5d|%20s|%5s|%10s|%50s\n", user.id, user.name, user.age, user.tel, user.addr)
	fmt.Println()
}

// 添加用户
func adduser(id int, users *[]user) {
	fmt.Println(users)
	var (
		name   string
		age    string
		tel    string
		addr   string
		choice string
	)

	// 调用用户输入函数，直接返回数据
	name, age, tel, addr = scanData()

	fmt.Println("你输入的信息如下：")
	titleString()
	fmt.Printf("%5d|%20s|%5s|%10s|%50s\n", id, name, age, tel, addr)

	fmt.Print("请确定是否插入数据?(y/n) ")
	choice = scanFunc()
	if choice == "y" || choice == "Y" {
		*users = append(*users, user{id, name, age, tel, addr})
		fmt.Println("添加成功,添加信息如下：")
		titleString()
		queryId(id, *users)
		return
	}
}

// 查询用户信息
func queryUser(users *[]user) {
	fmt.Println(users)
	var (
		queryStr string
		isEmpty  bool
	)

	// 判断users是否是空
	if isEmpty = isEmptyUser(*users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来查询.")
		return
	}

	fmt.Print("请输入想要查询信息：")
	queryStr = scanFunc()

	titleString()
	for _, user := range *users {
		if strings.Contains(user.name, queryStr) || strings.Contains(user.tel, queryStr) || strings.Contains(user.addr, queryStr) {
			formatUser(user)
		}
	}
}

// 修改用户
func modifyUser(users []user) {
	fmt.Println(users)
	var (
		id      int
		choice  string
		isEmpty bool
		name    string
		age     string
		tel     string
		addr    string
	)

	// 判断users是否是空
	if isEmpty = isEmptyUser(users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来修改.")
		return
	}

	fmt.Print("请输入你想修改的ID: ")
	id = scanId()

	for i, u := range users {
		if u.id == id {
			titleString()
			formatUser(u)
			fmt.Print("请确定是否修改此用户(y/n):")
			choice = scanFunc()
			if choice == "y" || choice == "Y" {
				name, age, tel, addr = scanData()
				users[i].name= name
				users[i].age= age
				users[i].tel= tel
				users[i].addr= addr
				fmt.Println("信息修改如下：")
				titleString()
				queryId(id, users)
				return
			}
		}
	}
	fmt.Println("你输入的id不存在")
}

// 删除用户
func deleteUser(users []user) []user {
	fmt.Println(users)
	var (
		id      int
		choice  string
		isEmpty bool
	)

	if isEmpty = isEmptyUser(users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来删除.")
		return nil
	}

	fmt.Print("请输入你想删除的ID: ")
	id = scanId()
	for i, u := range users {
		if u.id == id {
			titleString()
			formatUser(u)
			fmt.Print("请确定是否需要删除此用户(y/n): ")
			choice = scanFunc()
			if choice == "y" || choice == "Y" {
				users = append(users[:i], users[i+1:]...)
				fmt.Println("用户删除成功.")
				return users
			}
		}
	}
	fmt.Println("用户id不存在.")
	return users
}

type user struct {
	id   int
	name string
	age  string
	tel  string
	addr string
}

func main() {
	// 列表结构体
	var (
		users  []user = []user{}
		op     string
		id     int
		passwd string = "123456"
	)

	// 用户认证
	if !auth(passwd) {
		return
	}

	fmt.Println("欢迎使用马哥用户系统.")
	for {
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入你想要操作的指令：`)

		op = scanFunc()
		fmt.Println("你输入的指令为：", op)

		if op == "1" {
			id++
			adduser(id, &users)
		} else if op == "2" {
			modifyUser(users)
		} else if op == "3" {
			users = deleteUser(users)
		} else if op == "4" {
			queryUser(&users)
		} else if op == "5" {
			fmt.Println("欢迎下次再来")
			break
		} else {
			fmt.Println("指令错误.")
			continue
		}
	}
}
