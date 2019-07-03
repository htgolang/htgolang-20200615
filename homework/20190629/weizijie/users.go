package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 添加用户
func add(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	fmt.Print("请输入年龄:")
	fmt.Scan(&age)

	fmt.Print("请输入电话:")
	fmt.Scan(&tel)

	fmt.Print("请输入地址:")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
}

// 显示所有用户信息
func list_user(users map[string]map[string]string) {
	fmt.Println("当前用户有:")
	for n, user := range users {
		fmt.Printf("%s : %v\n", n, user)
	}
}

// 修改指定key的值
func modify_value(choose string, key, value string, users map[string]map[string]string) {
	if choose == "1" {
		users[key]["name"] = value
	} else if choose == "2" {
		users[key]["age"] = value
	} else if choose == "3" {
		users[key]["tel"] = value
	} else if choose == "4" {
		users[key]["addr"] = value
	} else {
		fmt.Println("请输入你的选择.....")
	}
}

// 修改用户信息
func modify(users map[string]map[string]string) {
	var num string   // 输入用户Key
	var m string     // 选择修改的信息
	var info string  // 输入修改的内容
	var input string // 输入Y/N,是否确认修改

	// 调用list_user函数，打印当前所有用户的信息
	list_user(users)

INPUT:
	fmt.Printf("请输入你要修改的用户ID：\n")
	fmt.Scan(&num)

	// 判断用户是否存在
	if _, ok := users[num]; ok {
		fmt.Println("该用户存在")
		fmt.Print(`请输入你要修改的用户信息：
1. 修改Name
2. 修改Age
3. 修改Tel
4. 修改Addr
请输入你的选择:`)
		// 选择需要修改的选项
		fmt.Scan(&m)
		fmt.Printf("你的选择是: %s\n", m)

		fmt.Println("是否要确认修改(Y/N):")
		fmt.Scan(&input)

		if input == "Y" || input == "y" {
			fmt.Printf("请输入你要修改的内容：\n")
			fmt.Scan(&info)

			// 调用modify_value函数，修改指定key的值
			modify_value(m, num, info, users)
		}
	} else {
		fmt.Println("该用户不存在,请重新输入")
		goto INPUT
	}
}

// 查询用户
// q == "" 查找全部
// 非空，名称 电话 住址任意一个内容包含q的显示
func query(users map[string]map[string]string) {
	var q string
	fmt.Print("请输入查询信息:")
	fmt.Scan(&q)

	title := fmt.Sprintf("%-5s|%-10s|%-5s|%-10s|%-15s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	for _, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%-5s|%-10s|%-5s|%-10s|%-15s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println("")
		}
	}
}

func remove(users map[string]map[string]string) {
	var num string

	// 调用list_user函数，打印当前所有用户的信息
	list_user(users)

INPUT:
	fmt.Printf("请输入你要修改的用户ID:")
	fmt.Scan(&num)

	for k := range users {
		if k == num {
			fmt.Println(users[k])
			var input string
			fmt.Printf("该用户存在,是否确认删除(Y/N):")
			fmt.Scan(&input)
			if input == "Y" || input == "y" {
				delete(users, k)
			}
		} else {
			fmt.Println("该用户不存在,请重新输入")
			goto INPUT
		}
	}
}

func main() {
	const password = "danran"
	var passwd string

	fmt.Println("本系统密码是:danran")
	defer func() {
		fmt.Println("欢迎使用本系统")
	}()
END:
	for i := 0; i < 3; i++ {
		fmt.Println("请输入系统密码：")
		fmt.Scan(&passwd)
		if passwd == password {
			// 存储用户信息
			users := make(map[string]map[string]string)
			id := 0

			fmt.Println("欢迎使用本系统")
			for {
				var op string
				fmt.Print(`1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入你的选择:`)
				fmt.Scan(&op)
				fmt.Printf("你的选择是: %s\n", op)

				if op == "1" {
					id++
					add(id, users)
				} else if op == "2" {
					modify(users)
				} else if op == "3" {
					remove(users)
				} else if op == "4" {
					query(users)
				} else if op == "5" {
					break END
				} else {
					fmt.Println("指令错误")
				}
			}
		} else {
			fmt.Println("密码错误，请再次输入...")
		}
	}
}
