package main

import (
	"fmt"
	"strconv"
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
	return
}

// 用户输入id
func scanId() (id string) {
	fmt.Print("请输入ID: ")
	fmt.Scan(&id)
	fmt.Println()
	return id
}

// 判断users是否为空, 为空就不让查询，删除等操作。 返回 users长度和是否为空
func isEmptyUser(users map[string]map[string]string) (flag bool) {
	if len(users) == 0 {
		return true
	}
	return false
}

// 打印信息头信息
func titleString() {
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "id", "name", "age", "tel", "addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

// 通过id, 查询用户信息. 打印信息
func queryId(id string, users map[string]map[string]string) {
	for _, user := range users {
		if id == user["id"] {
			formatUser(user)
		}
	}
}

// 传入user, 输出打印信息
func formatUser(user map[string]string) {
	fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
	fmt.Println()
}

// 添加用户
func adduser(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
		choice string
	)

	// 调用用户输入函数，直接返回数据
	name, age, tel, addr = scanData()

	fmt.Println()
	fmt.Println("你输入的信息如下：")
	titleString()
	fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", id, name, age, tel, addr)

	fmt.Println()
	fmt.Print("请确定是否插入数据?(y/n) ")
	fmt.Scan(&choice)

	if choice == "y" || choice == "Y" {
		users[id] = map[string]string{
			"id":   id,
			"name": name,
			"age":  age,
			"tel":  tel,
			"addr": addr,
		}
		fmt.Println("添加成功,添加信息如下：")
		titleString()
		queryId(id, users)
	}
}

// 查询用户信息
func queryUser(users map[string]map[string]string) {
	var (
		queryStr string
		isEmpty  bool
	)

	// 判断users是否是空
	if isEmpty = isEmptyUser(users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来查询.")
		return
	}

	fmt.Print("请输入想要查询信息：")
	fmt.Scan(&queryStr)

	// 打印titile
	titleString()
	for _, user := range users {
		if strings.Contains(user["name"], queryStr) || strings.Contains(user["tel"], queryStr) || strings.Contains(user["addr"], queryStr) {
			// 调用格式化输出函数
			formatUser(user)
		}
	}
}

// 修改用户
func modifyUser(users map[string]map[string]string) {
	var (
		id string
		choice string
		isEmpty bool
		name string
		age  string
		tel  string
		addr string
	)

	// 判断users是否是空
	if isEmpty = isEmptyUser(users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来修改.")
		return
	}

	// 用户输入id
	id = scanId()
	if user, ok := users[id]; ok {
		titleString()
		formatUser(user)
		fmt.Print("请确定是否修改此用户(y/n):")
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			name, age, tel, addr = scanData()
			titleString()
			fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", id, name, age, tel, addr)
			fmt.Print("请确定是否修改数据?(y/n) ")
			fmt.Scan(&choice)
			if choice == "y" || choice == "Y" {
				users[id] = map[string]string{
					"id":   id,
					"name": name,
					"age":  age,
					"tel":  tel,
					"addr": addr,
				}
			}
			fmt.Println("信息修改如下：")
			titleString()
			queryId(id, users)
			return
		}
	}
}

// 删除用户
func deleteUser(users map[string]map[string]string) {
	var (
		id      string
		choice  string
		isEmpty bool
	)

	if isEmpty = isEmptyUser(users); isEmpty {
		fmt.Println("暂无用户信息, 请先录入用户信息, 再来删除.")
		return
	}

	id = scanId()
	if user, ok := users[id]; ok {
		titleString()
		formatUser(user)
		fmt.Print("请确定是否需要删除此用户(y/n): ")
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			delete(users, id)
			fmt.Println("用户删除成功.")
		}
	} else {
		fmt.Println("暂无此用户")
	}
}

// 主函数
func main() {
	// 用户列表
	var (
		users  map[string]map[string]string = map[string]map[string]string{}
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

		fmt.Scan(&op)
		fmt.Println("你输入的指令为：", op)

		if op == "1" {
			id++
			adduser(id, users)
		} else if op == "2" {
			modifyUser(users)
		} else if op == "3" {
			deleteUser(users)
		} else if op == "4" {
			queryUser(users)
		} else if op == "5" {
			fmt.Println("欢迎下次再来")
			break
		} else {
			fmt.Println("指令错误.")
			continue
		}
	}
}
