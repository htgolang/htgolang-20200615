package main

import (
	"flag"
	"fmt"
	"os"
	"5day/usermanagerWork/users"
)

func main()  {
	var noauth bool
	flag.BoolVar(&noauth,"N",false,"no auth")
	flag.Parse()
	if !noauth && !users.Auth(){
		fmt.Printf("密码输入%d次错误，程序退出",users.MaxAuth)
		return
	}
	menu:= 	`===========================================
1.查询
2.添加
3.修改
4.删除
5.退出
===========================================
`
	callbacks := map[string]func(){
		"1": users.QueryUser,
		"2": users.AddUser,
		"3": users.ModifyUser,
		"4": users.DeleteUser,
		"5": func() {
			os.Exit(0)
		},
	}
	fmt.Println("欢迎登陆用户管理系统")
	for {
		fmt.Println(menu)
		if callbacks,ok := callbacks[users.InputString("请输入你的指令:")];ok {
			callbacks()
		}else {
			fmt.Println("指令错误")
		}
	}
}