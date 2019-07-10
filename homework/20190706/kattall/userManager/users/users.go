package users

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"strconv"
	"strings"
)

const (
	MaxAuth  = 3
	// 密码 123456
	password = "e10adc3949ba59abbe56e057f20f883e"
)

func InputString(prompt string) (input string) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func Auth() bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码：")
		if inputpwd, err := gopass.GetPasswd(); err == nil {
			if fmt.Sprintf("%x", md5.Sum([]byte(inputpwd))) == password {
				return true
			} else {
				fmt.Println("密码错误.")
			}
		}
	}
	return false
}

func getUserID(users_lst map[int]map[string]string) (id int) {
	// 判断是否为空，如果为空，就返回1
	if len(users_lst) == 0 {
		return 1
	}
	for k := range users_lst {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = InputString("请输入姓名：")
	user["age"] = InputString("请输入年龄：")
	user["tel"] = InputString("请输入联系电话：")
	user["addr"] = InputString("请输入联系地址：")
	user["desc"] = InputString("请输入备注：")
	return user
}

func printUser(pk int, user map[string]string) {
	fmt.Println("ID:", pk)
	fmt.Println("名字:", user["name"])
	fmt.Println("年龄:", user["age"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("联系地址:", user["addr"])
	fmt.Println("备注:", user["desc"])
}

func Query(users_lst map[int]map[string]string) {
	q := InputString("请输入查询的内容：")
	fmt.Println("=================================")
	for idx, user := range users_lst {
		if strings.Contains(user["name"], q) || strings.Contains(user["addr"], q) || strings.Contains(user["desc"], q) {
			printUser(idx, user)
		}
	}
	fmt.Println("=================================")
}

func Add(users_lst map[int]map[string]string) {
	id := getUserID(users_lst)
	user := inputUser()
	users_lst[id] = user
	fmt.Println("[+]用户添加成功")
}

func Modify(users_lst map[int]map[string]string) {
	if id, err := strconv.Atoi(InputString("请输入修改的ID：")); err == nil {
		if user, ok := users_lst[id]; ok {
			fmt.Println("将修改用户信息：")
			printUser(id, user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser()
				users_lst[id] = user
				fmt.Println("[+]修改成功")
			} else {
				fmt.Println("[-]用户ID不存在.")
			}
		} else {
			fmt.Println("[-]输入ID不正确")
		}

	}
}

func Del(users_lst map[int]map[string]string) {
	if id, err := strconv.Atoi(InputString("请输入删除的ID：")); err == nil {
		if user, ok := users_lst[id]; ok {
			fmt.Println("将删除用户信息：")
			printUser(id, user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users_lst, id)
				fmt.Println("[+]删除成功")
			} else {
				fmt.Println("[-]用户ID不存在.")
			}
		} else {
			fmt.Println("[-]输入ID不正确")
		}

	}
}

/*
func UserStart(){
	if !Auth() {
		fmt.Printf("密码错误%d次, 系统退出.", MaxAuth)
		return
	}

	menu := `*******************************
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	users_lst := map[int]map[string]string{}
	callbacks := map[string]func(map[int]map[string]string){
		"1": Query,
		"2": Add,
		"3": Modify,
		"4": Del,
		"5": func(users_list map[int]map[string]string){
			os.Exit(0)
		},
	}

	fmt.Println("欢迎进入用户管理系统.")
	for {
		fmt.Println(menu)
		if callback, ok := callbacks[InputString("请输入指令:")]; ok {
			callback(users_lst)
		} else {
			fmt.Println("用户指令错误.")
		}
	}
}
*/