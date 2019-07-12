package main

import (
	"fmt"
)

const (
	maxAuth  = 3
	password = "123"
)

func main() {
	if !GoStudynotes.Auth(password, maxAuth) {
		fmt.Printf("密码输入%d次错误，程序退出", maxAuth)
		return
	}
	userMenu()
}
func userMenu() {
	menu := `===========================================
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
		case "5":
			break END
		case "q":
			break END
		default:
			fmt.Println("[err]指令无效!")
		}
	}
}

/*
评分: 7
建议：
1. 包名使用全小写英文字母，并且与所在文件名一致
2. 注意代码组织方式，按代码按照操作对象或逻辑按文件存放
*/
