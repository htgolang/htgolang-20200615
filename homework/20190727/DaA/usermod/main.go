package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/xxdu521/usermod/users"
)

func main() {
	//测试
	//users.Test()

	//认证
	auth := flag.Bool("N", false, "no auth")
	flag.Parse()
	//fmt.Println(*auth)
	if !users.Auth(*auth) {
		return
	}

	//功能测试
	menu := `
1. 新建用户
2. 查询用户
3. 删除用户
4. 更新用户
5. 修改密码
6. 退出
*********************************`

	callbacks := map[int]func(){
		1: users.Add,
		2: users.Query,
		3: users.Del,
		4: users.Update,
		5: users.UpdatePasswd,
		6: func() { os.Exit(0) },
	}

	for {
		fmt.Println(menu)
		if ID, err := strconv.Atoi(users.Inputstring("请选择功能项: ")); err == nil {
			if callback, ok := callbacks[ID]; ok {
				//if callback, ok := callbacks[strconv.Atoi(users.Inputstring("请选择功能项: "))]; ok {
				callback()
			} else {
				print("输入错误，请重新输入!!!")
			}
		} else {
			fmt.Println(err)
		}
	}

}
