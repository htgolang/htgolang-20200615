package main

import (
	"fmt"
	"os"

	ulist "github.com/huang19910425/usermanager/users"
)

func main() {
	if !ulist.Auth() {
		fmt.Printf("密码错误%d次, 系统退出.", ulist.MaxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	callbacks := map[string]func(){
		"1": ulist.Query,
		"2": ulist.Add,
		"3": ulist.Modify,
		"4": ulist.Del,
		"5": func() {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入用户管理系统.")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[ulist.InputString("请输入指令:")]; ok {
			callback()
		} else {
			fmt.Println("用户指令错误.")
		}
	}
}
