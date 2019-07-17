package operate

import (
	"fmt"
	"strings"

	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/userstruct"
)

// 查询函数
func Query() {
	q := inputstring.InputString("请输入查询的信息:")

	title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))

	// 若输入的查询信息包含在Nmae、Desc、Addr任意一个参数中，则返回用户信息
	for _, user := range userstruct.User {
		if strings.Contains(user.Name, q) || strings.Contains(user.Desc, q) || strings.Contains(user.Addr, q) {
			fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
		}
	}

}
