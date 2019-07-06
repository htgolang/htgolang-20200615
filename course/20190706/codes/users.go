package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxAuth  = 3
	password = "123!@#"
)

func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

// 从命令行输入密码, 并进行验证
// 通过返回值告知验证成功还是失败
func auth() bool {
	var input string
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入密码:")
		fmt.Scan(&input)
		if password == input {
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

func query(users map[int]map[string]string) {
	q := inputString("请输入查询内容:")
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

func getId(users map[int]map[string]string) int {
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
	user["name"] = inputString("请输入名字:")
	user["birthday"] = inputString("请输入出生日期(2000-01-01):")
	user["tel"] = inputString("请输入联系方式:")
	user["addr"] = inputString("请输入联系地址:")
	user["desc"] = inputString("请输入备注:")
	return user
}

func add(users map[int]map[string]string) {
	id := getId(users)
	user := inputUser()
	users[id] = user
	fmt.Println("[+]添加成功")
}

func modify(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputString("请输入修改用户ID:")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			printUser(id, user)
			input := inputString("确定修改(Y/N)?")
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

func del(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputString("请输入删除用户ID:")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			printUser(id, user)
			input := inputString("确定删除(Y/N)?")
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

func main() {
	if !auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", maxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	// id, name, birthday, tel, addr, desc
	// users := map[int][5]string
	// users := map[int][]string
	// users := []map[string]string{}
	// users := [][]string{}
	// users := [][5]string{}
	users := map[int]map[string]string{}

	callbacks := map[string]func(map[int]map[string]string){
		"1": query,
		"2": add,
		"3": modify,
		"4": del,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入KK的用户管理系统")
	// END:
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[inputString("请输入指令:")]; ok {
			callback(users)
		} else {
			fmt.Println("指令错误")
		}
		// switch op {
		// case "1":
		// 	query(users)
		// case "2":
		// 	add(users)
		// case "3":
		// 	modify(users)
		// case "4":
		// 	del(users)
		// case "5":
		// 	break END
		// default:
		// 	fmt.Println("指令错误")
		// }
	}
}
