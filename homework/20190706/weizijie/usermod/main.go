package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JevonWei/user/auth"
	"github.com/JevonWei/user/input"
	"github.com/JevonWei/user/operate"
	"github.com/JevonWei/usermod/title"
)

func print(s string) {
	fmt.Println(s)
}

func main() {
	title.Title_String()

	if !auth.Auth() {
		return
	}

	users := map[int]map[string]string{}

	callbacks := map[string]func(map[int]map[string]string){
		"1": operate.Query,
		"2": operate.Add,
		"3": operate.Modify,
		"4": operate.DelUser,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}
	//END:
	for {
		print(strings.Repeat("-", 30))
		print(title.Menu)

		if callback, ok := callbacks[input.InputString("请输入你选择的操作:")]; ok {
			callback(users)
		} else {
			print("选择无效，请重新输入!!!")
		}
	}
}
