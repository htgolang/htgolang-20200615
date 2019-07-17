package operate

import (
	"fmt"
	"strconv"

	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/listuser"
	"github.com/JevonWei/usermanager/userstruct"
)

// 删除用户
func Deluser() {
	// 调用函数，显示系统中的所有用户
	listuser.Listuser()

	// 按照输入的用户ID，删除用户
	idString := inputstring.InputString("请输入删除用户ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := userstruct.User[id]; ok {
			fmt.Println("将要删除的用户信息为:")
			fmt.Println("================================")
			fmt.Println("ID:", user.ID)
			fmt.Println("Name:", user.Name)
			fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
			fmt.Println("联系方式:", user.Tel)
			fmt.Println("地址:", user.Addr)
			fmt.Println("描述:", user.Desc)

			// 确认是否删除用户
			in := inputstring.InputString("是否确定删除(Y/N)?")
			if in == "Y" || in == "y" {
				delete(userstruct.User, id)
				fmt.Printf("ID为%d的用户已删除\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}
