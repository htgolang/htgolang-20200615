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

/*
评分: 7
建议：
1. 包名使用全小写英文字母，并且与所在文件名一致
2. 注意代码组织方式，按代码按照操作对象或逻辑按文件存放
*/
