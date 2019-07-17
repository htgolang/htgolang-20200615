package listuser

import (
	"fmt"
	"strings"

	"github.com/JevonWei/usermanager/userstruct"
)

// 定义Listuser函数，打印用户系统中所有的用户
func Listuser() {

	title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))

	// 遍历所有的用户，并打印
	for _, user := range userstruct.User {
		fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
	}
}
