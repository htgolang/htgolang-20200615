package GoStudynotes

import (
	"fmt"
	"strconv"
)

func ModifyUser(users map[int]map[string]string){
	if id,err := strconv.Atoi(inputString("[info]请输入你即将修改的用户id:"));err == nil {
		// 判断id在不在
		if user,ok :=  users[id];ok{
			printUser(id,user)
			fmt.Printf("[warning]即将修改的用户信息:%v",user)
			input := inputString("\n你确定修改吗？（Y/y）")
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