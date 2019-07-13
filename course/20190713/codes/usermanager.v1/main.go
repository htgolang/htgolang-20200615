package main

import (
	"fmt"
	"os"

	"flag"

	upkg "github.com/imsilence/usermanager/users"
)

func main() {

	var noauth bool
	flag.BoolVar(&noauth, "N", false, "no auth")

	flag.Parse()

	if !noauth && !upkg.Auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", upkg.MaxAuth)
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
		"1": upkg.Query,
		"2": upkg.Add,
		"3": upkg.Modify,
		"4": upkg.Del,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入KK的用户管理系统")
	// END:
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[upkg.InputString("请输入指令:")]; ok {
			callback(users)
		} else {
			fmt.Println("指令错误")
		}
	}
}
