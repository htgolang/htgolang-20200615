package inputuser

import (
	"errors"
	"fmt"

	"github.com/JevonWei/usermanager/birthday"
	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/userstruct"
)

// 定义用户输入函数，将从键盘输入的每个值对应传入Users结构体的元素中
func Inputuser(num int) {
	// 定义结构体类型Users的变量User_input
	var User_input userstruct.Users = userstruct.Users{}

	// 从键盘输入name(用户名)
	name := inputstring.InputString("请输入名字:")

	// 若name为空，则不赋值，直接返回
	if name == "" {
		fmt.Println("输入的Name不能为空")
		goto END
	}

	// 若输入的name，在系统中已存在，则提示用户已存在，不能新建用户
	for _, user := range userstruct.User {
		if name == user.Name {
			fmt.Printf("用户名%s已存在,不能新增/修改\n", name)
			goto END
		}
	}

	// 将输入的name，赋值给Users结构体的Name
	User_input.Name = name
	// 将User映射的key值，赋值给结构体Users的ID
	User_input.ID = num

	// 将输入的字符串类型的Birthday转换为时间类型
	// 判断输入的Birthday格式是否正确，若输入格式有误，则打印提示信息，并重新输入
	for {
		birthday_time, err := birthday.Birthday_time(inputstring.InputString("请输入出生日期(2019-07-07):"))
		if err == nil {
			User_input.Birthday = birthday_time
			break
		} else {
			fmt.Println(errors.New("请输入正确认格式"))
		}
	}

	// 将输入的其他值依次赋值为Tel，Addr，Desc,并返回用户信息

	//User_input.Birthday = inputstring.InputString("请输入出生日期(2019-07-07):")
	User_input.Tel = inputstring.InputString("请输入联系方式:")
	User_input.Addr = inputstring.InputString("请输入地址:")
	User_input.Desc = inputstring.InputString("请输入描述信息:")
	fmt.Println("*******************************")
	fmt.Printf("ID为%d的用户已添加/修改\n", num)

	userstruct.User[num] = User_input
END:
}
