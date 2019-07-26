package main

import (
	"crypto/md5"
	"fmt"

	"github.com/suren00/usermanager/users"
)

func main() {

	err := users.PasswdFileCheck()
	if err != nil {
		fmt.Println(err)
		var inputPass string
		fmt.Scan(&inputPass)
		inputPass = fmt.Sprintf("%x", md5.Sum([]byte(inputPass)))
		users.Passwd2File(inputPass)
	} else {
		if !users.Auth() {
			return
		}
	}

	/* 	if !users.Auth() {
		return
	} */

	methods := map[string]func(){
		"1": users.Add,
		"2": users.Modify,
		"3": users.Del,
		"4": users.Query,
	}

	msg := `
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出

请输入指令:
`
	for {
		fmt.Println(msg)
		var op string
		fmt.Scan(&op)
		if method, ok := methods[op]; ok {
			method()
		} else if op == "5" {
			break
		} else {
			fmt.Println("输入的选项不存在！！！！")
		}
	}
}
