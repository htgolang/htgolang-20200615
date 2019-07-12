package main

import (
	"fmt"
	"github.com/marksugar/GoStudynotes/useredit"
)
const (
	maxAuth =3
	password = "123"
)
func main(){
	if !GoStudynotes.Auth(password,maxAuth){
		fmt.Printf("密码输入%d次错误，程序退出",maxAuth)
		return
	}
	userMenu()
}
func userMenu(){
	menu:= 	`===========================================
1.查询
2.添加
3.修改
4.删除
5.退出
===========================================
`
	users := map[int]map[string]string{}
	fmt.Println("欢迎登陆用户管理系统")
END:
	for {
		fmt.Println(menu)
		fmt.Print("请输入指令编号:")
		var op string
		fmt.Scan(&op)
		switch op {
		case "1":
			GoStudynotes.QueryUser(users)
		case "2":
			GoStudynotes.AddUser(users)
		case "3":
			GoStudynotes.ModifyUser(users)
		case "4":
			GoStudynotes.DeleteUser(users)
		case "5" :
			break END
		case "q":
			break END
		default:
		    fmt.Println("[err]指令无效!")
		}
	}
}