package main

import (
	"fmt"

	"github.com/suren00/gopkg"
)

func main() {

	if !gopkg.Auth() {
		return
	}

	users := make(map[int]map[string]string)

	methods := map[string]func(map[int]map[string]string){
		"1": gopkg.Add,
		"2": gopkg.Modify,
		"3": gopkg.Del,
		"4": gopkg.Query,
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
			method(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("输入的选项不存在！！！！")
		}
	}
}

/*
评分: 6
建议：
1. 项目使用module方式组织
2. 模块名使用功能进行命名，全小写，见名知意
3. 自己的代码自己负责，进行测试
*/
