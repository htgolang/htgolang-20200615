package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(uid int, users map[string]map[string]string) {
	userInfo(strconv.Itoa(uid), users)
}

// 查询用户
// 非空, 名称、电话、住址中任意一个属性中包含q内容的显示
func query(users map[string]map[string]string) {
	var q string

	fmt.Printf("请输入要查询的信息: ")
	fmt.Scan(&q)
	printCommon("query", q, users)
}

func modify(users map[string]map[string]string) {
	var modifyID, modifyStat string
	fmt.Printf("请输入要修改的用户ID: ")
	fmt.Scan(&modifyID)

	if checkUID(modifyID, users) == true {
		printCommon("modify", modifyID, users)
		fmt.Printf("请确认是否进行修改(yes/no): ")
		fmt.Scan(&modifyStat)
		switch modifyStat {
		case "yes":
			userInfo(modifyID, users)
			fmt.Println("用户修改成功！！！")
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func del(users map[string]map[string]string) {
	var deleteID, deleteStat string
	fmt.Printf("请输入要删除的用户ID: ")
	fmt.Scan(&deleteID)

	if checkUID(deleteID, users) == true {
		printCommon("delete", deleteID, users)
		fmt.Printf("请确认是否删除(yes/no): ")
		fmt.Scan(&deleteStat)
		switch deleteStat {
		case "yes":
			delete(users, deleteID)
			fmt.Println("用户删除成功！！！")
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func checkUID(uid string, users map[string]map[string]string) bool {
	if _, ok := users[uid]; ok {
		return true
	} else {
		fmt.Println("你输入的用户ID不存在, 请确认ID后, 在进行操作")
		return false
	}
}

func userInfo(uid string, users map[string]map[string]string) {
	var (
		name string
		age  string
		tal  string
		addr string
	)

	fmt.Printf("请输入名称: ")
	fmt.Scan(&name)

	fmt.Printf("请输入年龄: ")
	fmt.Scan(&age)

	fmt.Printf("请输入联系方式: ")
	fmt.Scan(&tal)

	fmt.Printf("请输入家庭住址: ")
	fmt.Scan(&addr)

	users[uid] = map[string]string{
		"id":   uid,
		"name": name,
		"age":  age,
		"tal":  tal,
		"addr": addr,
	}
}

func printCommon(action, keyword string, users map[string]map[string]string) {

	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%30s", "ID", "Name", "Age", "Tal", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	switch action {
	case "query":
		for _, user := range users {
			if strings.Contains(user["name"], keyword) || strings.Contains(user["age"], keyword) || strings.Contains(user["tal"], keyword) || strings.Contains(user["addr"], keyword) {
				fmt.Printf("%5s|%20s|%5s|%20s|%30s\n", user["id"], user["name"], user["age"], user["tal"], user["addr"])
			}
		}
	case "modify", "delete":
		user := users[keyword]
		fmt.Printf("%5s|%20s|%5s|%20s|%30s\n", user["id"], user["name"], user["age"], user["tal"], user["addr"])
	}
}

func main() {
	passwd := "123abc!@#"
	var inputPass string
	fmt.Println("欢迎使用马哥用户系统, 请输入管理员密码进入: ")
	for i := 3; i > 0; i-- {
		fmt.Scan(&inputPass)
		if inputPass == passwd {
			break
		} else {
			fmt.Printf("密码输入错误, 你还有%d次机会\n", i-1)
		}
		if i == 1 {
			panic("passowrd error")
		}
	}
	users := make(map[string]map[string]string)
	id := 0

	for {
		var op string
		fmt.Println(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出

请输入指令:`)

		fmt.Scan(&op)

		if op == "1" {
			id++
			add(id, users)
		} else if op == "2" {
			modify(users)
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
评分: 8
考虑：
1. 思考if else-if else如何使用switch-case替代，更深入思考利用函数类型如何简写
2. 考虑代码单一职责，一个函数只提供一个职责 printCommon
*/
