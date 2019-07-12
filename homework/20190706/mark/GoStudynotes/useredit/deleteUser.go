package GoStudynotes

import (
	"fmt"
	"strconv"
)

func DeleteUser(users map[int]map[string]string){
	if id,err := strconv.Atoi(inputString("请输入你即将删除的用户id:"));err == nil {
		// 判断id在不在
		if user,ok :=  users[id];ok{
			printUser(id,user)
			fmt.Printf("[warning]即将删除的用户信息:%v",user)
			input := inputString("\n你确定删除吗？（Y/y）")
			if input == "y"|| input == "Y" {
				delete(users,id)
				fmt.Printf("[ok]删除%v成功\n",id)
			}
		}else {
			fmt.Println("[err]用户id不存在\n")
		}
	}
}
