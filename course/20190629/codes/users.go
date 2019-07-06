package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 添加用户
func add(users map[string]map[string]string) {
	pk := 0
	for k := range users {
		if v, err := strconv.Atoi(k); err == nil && v > pk {
			pk = v
		}
	}
	var (
		id   string = strconv.Itoa(pk + 1)
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
}

// 查询用户
// 非空， 名称，电话，住址任意一个属性中包含q内容的显示
func query(users map[string]map[string]string) {
	var q string

	fmt.Print("请输入查询信息:")
	fmt.Scan(&q)
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	for _, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%5s|%20s|%5s|%20s|%50s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}

func main() {
	// 存储用户信息
	users := make(map[string]map[string]string)
	fmt.Println("欢迎使用马哥用户系统")

	callbacks := map[string]func(map[string]map[string]string){
		"1": add,
		"4": query,
	}

	for {
		var op string
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入指令:`)
		fmt.Scan(&op)

		if callback, ok := callbacks[op]; ok {
			callback(users)
		} else {
			fmt.Println("指令错误")
		}
	}
}
