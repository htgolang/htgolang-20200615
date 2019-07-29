package main

import (
	"fmt"
	"os"

	upkg "github.com/xlotz/usersmem"
)

func main() {
	if !upkg.Auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", upkg.MaxAuth)
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
		"1": upkg.Query,
		"2": upkg.Add,
		"3": upkg.Modify,
		"4": upkg.Del,
		"5": upkg.ModifyPass,
		"6": func() {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入KK的用户管理系统")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[upkg.InputString("请输入指令:")]; ok {
			callback()
		} else {
			fmt.Println("指令错误")
		}
	}
}
