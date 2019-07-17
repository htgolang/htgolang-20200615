package operate

import (
	"github.com/JevonWei/usermanager/getid"
	"github.com/JevonWei/usermanager/inputuser"
)

// 添加函数
func Add() {
	id := getid.GetId()

	// 调用用户函数，新增用户

	inputuser.Inputuser(id)
	//userstruct.User[id] = inputuser.Inputuser(id)

	//userstruct.User.ID++
	//userstruct.User.Name = inputstring.InputString("请输入名字:")
	//userstruct.User.Birthday = inputstring.InputString("请输入出生日期(2019-07-07):")
	//userstruct.User.Tel = inputstring.InputString("请输入联系方式:")
	//userstruct.User.Addr = inputstring.InputString("请输入地址:")
	//userstruct.User.Desc = inputstring.InputString("请输入描述信息:")
	//fmt.Println("*******************************")

}
