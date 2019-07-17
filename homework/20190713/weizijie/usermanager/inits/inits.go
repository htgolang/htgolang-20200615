package inits

import (
	"fmt"
	"strings"
)

// 定义变量Menu为操作的可选项
var Menu = `1. 显示
2. 查询
3. 添加
4. 修改
5. 删除
6. 退出
*******************************`

// 定义变量显示所有的Users结构体类型的所有参数
var Sort_menu = `1. ID
2. Name
3. Birthday
4. Addr
5. Tel
6. Desc
*******************************`

// 在用户系统登录前，显示提示信息
func Title_String() {
	fmt.Println("JevonWei用户系统密码为:danran")
	fmt.Println("")

	// strings.Repeat() 显示特定字符多少次
	fmt.Println(strings.Repeat("*", 30))
	Head := "欢迎进入JevonWei的用户管理系统"
	fmt.Println(Head)
}
