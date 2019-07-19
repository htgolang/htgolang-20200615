package users

import (
	"fmt"
	"strconv"
)

func ModifyUser(){
	if id,err := strconv.Atoi(InputString("[info]请输入你即将修改的用户id，输入5退出:"));err == nil {
		// 判断id在不在
		quits(string(id))
		if user,ok :=  users[id];ok{
			printUser(user)
			fmt.Printf("[warning]即将修改的用户信息:%v",user)
			input := InputString("\n你确定修改吗？（Y/y）")
			quits(input)
			if input == "y"|| input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Printf("[ok]修改%v成功\n",id)
			}
		}else {
			fmt.Println("[err]用户id不存在\n")
		}
	}
}