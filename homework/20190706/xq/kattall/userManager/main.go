package main

import (
	"fmt"
	"os"
	"userManager/users"
)

func main() {
	if !users.Auth() {
		fmt.Printf("密码错误%d次, 系统退出.", users.MaxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	users_lst := map[int]map[string]string{}
	callbacks := map[string]func(map[int]map[string]string){
		"1": users.Query,
		"2": users.Add,
		"3": users.Modify,
		"4": users.Del,
		"5": func(users_list map[int]map[string]string) {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入用户管理系统.")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[users.InputString("请输入指令:")]; ok {
			callback(users_lst)
		} else {
			fmt.Println("用户指令错误.")
		}
	}
}
