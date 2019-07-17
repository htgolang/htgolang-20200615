package usersort

import (
	"fmt"
	"sort"
	"strings"

	"github.com/JevonWei/usermanager/inits"
	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/userstruct"
)

// 定义排序函数 ，将用户按照指定的参数排序，并返回用户数组

func User_sort() []userstruct.Users {

	// 定义空数组，存储用户信息
	Users_array := []userstruct.Users{}

	// 输入按照哪个参数排序，并将输入的参数赋值给变量o
	o := inputstring.InputString("请输入需要排序的键值:")

	//将系统中的用户保存在Users_array数组中
	for _, user := range userstruct.User {
		Users_array = append(Users_array, user)
	}

	// 调用sort.Slice()函数，根据输入的排序参数，排序系统的所有用户
	switch o {
	case "1":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].ID < Users_array[j].ID
		})
	case "2":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Name < Users_array[j].Name
		})
	case "3":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Birthday.Format("2006/01/02") < Users_array[j].Birthday.Format("2006/01/02") // 将time类型的Birthday值转换为字符串排序
		})
	case "4":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Addr < Users_array[j].Addr
		})
	case "5":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Tel < Users_array[j].Tel
		})
	case "6":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Desc < Users_array[j].Desc
		})
	case "7":
		break
	}
	return Users_array
}

// 定义函数将排序后的用户打印

func Print_sort() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(inits.Sort_menu)

	// 将排序后的用户数组赋值给list变量
	list := User_sort()

	// 打印排序后的用户
	fmt.Println(list)
	for _, v := range list {
		fmt.Printf("用户ID: %d\n", v.ID)
		fmt.Printf("用户Name: %s\n", v.Name)
		fmt.Printf("用户Birthday: %s\n", v.Birthday.Format("2006/01/02"))
		fmt.Printf("用户Addr: %s\n", v.Addr)
		fmt.Printf("用户Tel: %s\n", v.Tel)
		fmt.Printf("用户Desc: %s\n", v.Desc)
		fmt.Println(strings.Repeat("=", 30))
	}
}
