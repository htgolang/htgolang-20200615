package main

import (
	"fmt"
	"github.com/xiaofan/usermanager/users"
	"os"
)

func main() {
	const (
		maxAuth  = 3
		password = "1fdb7184e697ab9355a3f1438ddc6ef9"
	)

	if !users.Auth(maxAuth, password) {
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

	userMap := map[int]users.User{}

	callbacks := map[string]func(map[int]users.User){
		"1": users.Query,
		"2": users.Add,
		"3": users.Modify,
		"4": users.Del,
		"5": func(userMap map[int]users.User) {
			os.Exit(0)
		},
	}
	fmt.Println("欢迎进入马哥用户管理系统")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[users.InputString("请输入指令:")]; ok {
			callback(userMap)
		} else {
			fmt.Println("指令错误")
		}
	}
}
