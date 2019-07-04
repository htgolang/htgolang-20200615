package main

import (
	"fmt"
	"strings"
)

func auth() bool {
	var password string = "123abc!@#"
	var count int
	var passwd_input string
	for {
		//fmt.Println(count)
		if count < 3 {
			fmt.Print("请输入密码：")
			fmt.Scan(&passwd_input)
			if passwd_input == password {
				return true
			}
			count++
			fmt.Println("密码错误，请重试")
		} else {
			fmt.Println("尝试次数过多, 退出管理系统。")
			return false
		}
	}
}

//add,update,del,query
func add(pk int, users map[string]map[string]string) {
	var (
		id   string = fmt.Sprintf("%d", pk)
		name string
		age  string
		tel  string
		addr string
	)
	//fmt.Println(id, name, age, tel, addr)
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入电话：")
	fmt.Scan(&tel)
	fmt.Print("请输入住址：")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
	fmt.Printf("\n\n用户添加成功：\n%5s|%20s|%5s|%20s|%50s", users[id]["id"], users[id]["name"], users[id]["age"], users[id]["tel"], users[id]["addr"])
}

func del(users map[string]map[string]string) {
	var (
		id   string
		exec string
	)
	fmt.Print("请输入要删除的用户ID [string|all]：")
	fmt.Scan(&id)
	if id == "all" {
		fmt.Print("确定要清空全部数据吗(y/Y)：")
		users = make(map[string]map[string]string)
		fmt.Println(users)
		fmt.Printf("\n\n已经清空全部数据。")

	} else if user, ok := users[id]; ok {
		fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
		fmt.Print("确定要删除吗(y/Y)：")
		fmt.Scan(&exec)
		if exec == "y" || exec == "Y" {
			delete(users, id)
			fmt.Printf("删除 %s 成功", users[id]["name"])
		} else {
			fmt.Printf("\n\n退出删除用户状态")
		}
	} else {
		fmt.Printf("\n\n用户ID不存在，删除用户失败")
	}
}

func update(users map[string]map[string]string) {
	var (
		id   string
		name string
		age  string
		tel  string
		addr string
		exec string
	)

	fmt.Print("请输入要修改的用户ID：")
	fmt.Scan(&id)
	if user, ok := users[id]; ok {
		fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
		fmt.Print("确认要修改吗(y/Y)：")
		fmt.Scan(&exec)
		if exec == "y" || exec == "Y" {
			fmt.Print("请输入姓名：")
			fmt.Scan(&name)
			fmt.Print("请输入年龄：")
			fmt.Scan(&age)
			fmt.Print("请输入电话：")
			fmt.Scan(&tel)
			fmt.Print("请输入住址：")
			fmt.Scan(&addr)

			users[id] = map[string]string{
				"id":   id,
				"name": name,
				"age":  age,
				"tel":  tel,
				"addr": addr,
			}
			fmt.Println("\n\n用户信息修改成功!")

		} else {
			fmt.Println("\n\n退出更新信息状态")
		}
	} else {
		fmt.Println("\n\n用户ID不存在，更新信息失败")
	}
}

func query(users map[string]map[string]string) {
	var q string
	fmt.Print("请输入查询信息 [string|all] ：")
	fmt.Scan(&q)

	//先打印title
	title := fmt.Sprintf("\n\n%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))

	for _, user := range users {
		if q == "all" || strings.Contains(user["name"], q) || strings.Contains(user["addr"], q) || strings.Contains(user["tel"], q) {
			fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
		}

	}
}

func main() {
	users := make(map[string]map[string]string)
	id := 0
	fmt.Println("马哥用户管理系统")

	if !auth() {
		fmt.Println(".............密码错误")
		return
	}

	for {
		var op string
		fmt.Print(`
----------------
1.add
2.update
3.del
4.query
5.exit

请输入指令：`)

		fmt.Scan(&op)
		if op == "1" {
			//fmt.Println("add")
			id++
			add(id, users)
		} else if op == "2" {
			//fmt.Println("update")
			update(users)
		} else if op == "3" {
			//fmt.Println("del")
			del(users)
		} else if op == "4" {
			//fmt.Println("query")
			query(users)
		} else if op == "5" {
			fmt.Println("\n\n退出管理系统，下次再见！")
			break
		} else {
			fmt.Printf("\n\n输入错误，请重新输入")
		}
	}
}
