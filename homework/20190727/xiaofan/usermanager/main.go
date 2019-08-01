package main

import (
	"fmt"
	"github.com/xiaofan/usermanager/users"
	"os"
)

func main() {

	if !users.Auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", users.MaxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 修改密码
6. 退出
*******************************`

	callbacks := map[string]func(){
		"1": users.Query,
		"2": users.Add,
		"3": users.Modify,
		"4": users.Del,
		"5": users.ModifyPassword,
		"6": func() {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入小凡用户管理系统")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[users.InputString("请输入指令:")]; ok {
			callback()
		} else {
			fmt.Println("指令错误")
		}
	}
}
