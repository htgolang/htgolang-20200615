package operate

import (
	"fmt"
	"strconv"

	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/inputuser"
	"github.com/JevonWei/usermanager/listuser"
	"github.com/JevonWei/usermanager/userstruct"
)

// 修改函数
func Modify() {
	// 调用函数，显示系统中的所有用户
	listuser.Listuser()

	// 根据输入的用户ID，修改用户信息
	idString := inputstring.InputString("请输入修改用户ID:")

	// 将输入的字符串类型的值转换为int类型
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := userstruct.User[id]; ok {
			fmt.Println("")
			fmt.Println("将要修改的用户信息为:")
			fmt.Println("================================")
			fmt.Println("ID:", user.ID)
			fmt.Println("Name:", user.Name)
			fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
			fmt.Println("联系方式:", user.Tel)
			fmt.Println("地址:", user.Addr)
			fmt.Println("描述:", user.Desc)

			in := inputstring.InputString("是否确定修改(Y/N)?: ")
			if in == "Y" || in == "y" {
				inputuser.Inputuser(id)

			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}
