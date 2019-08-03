package main

import (
	"flag"
	"fmt"
	"os"

	upkg "github.com/imsilence/usermanager/users"
)

func main() {

	var persistenceType string
	var help bool

	flag.StringVar(&persistenceType, "T", "gob", "Persisitence Type[json, gob]")
	flag.BoolVar(&help, "H", false, "Help")

	flag.Usage = func() {
		fmt.Println("usage: usermangaer.exe -T [gob/json]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if !upkg.Auth() {
		fmt.Printf("[-]密码%d次错误, 程序退出\n", upkg.MaxAuth)
		return
	}

	upkg.SetPersistence(persistenceType)

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	callbacks := map[string]func(){
		"1": upkg.Query,
		"2": upkg.Add,
		"3": upkg.Modify,
		"4": upkg.Del,
		"5": func() {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入KK的用户管理系统")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[upkg.InputString("请输入指令:")]; ok {
			callback()
		} else {
			fmt.Println("指令错误")
		}
	}
}
