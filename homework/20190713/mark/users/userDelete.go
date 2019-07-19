package users

import (
	"fmt"
	"strconv"
)

func DeleteUser(){
	if id,err := strconv.Atoi(InputString("请输入你即将删除的用户id，输入5退出:"));err == nil {
		// 判断id在不在
		if user,ok :=  users[id];ok{
			printUser(user)
			fmt.Printf("[warning]即将删除的用户信息:%v",user)
			input := InputString("\n你确定删除吗？（Y/y）")
			quits(input)
			if input == "y"|| input == "Y" {
				delete(users,id)
				fmt.Printf("[ok]删除%v成功\n",id)
			}
		}else {
			fmt.Println("[err]用户id不存在\n")
		}
	}
}