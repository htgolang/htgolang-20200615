package main

import (
	"fmt"
	"github.com/xlotz/usersmem"
	"os"
	"flag"
    //"github.com/xlotz/usermanagerv3/usersmem"

)

func main() {

	var noauth bool
	flag.BoolVar(&noauth, "N", false, "no auth")

	flag.Parse()

	if !noauth && !usersmem.Auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", usersmem.MaxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`


	//users := map[int]map[string]string{}

	callbacks := map[string]func(){
		"1": usersmem.Query,
		"2": usersmem.Add,
		"3": usersmem.Modify,
		"4": usersmem.Delete,
		"5": func() {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入KK的用户管理系统")
	// END:
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[usersmem.InputString("请输入指令:")]; ok {
			callback()
		} else {
			fmt.Println("指令错误")
		}
	}
}
