package users

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func InputTime(prompt string) time.Time {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		return time.Time{}
	}
	return t
}

func Auth(maxAuth int, password string) bool {
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入密码:")
		input, _ := gopass.GetPasswd()

		if password == fmt.Sprintf("%x", md5.Sum(input)) {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func PrintUser(id int, user User) {
	fmt.Println("ID:", id)
	fmt.Println("名字:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006-01-02"))
	fmt.Println("联系方式:", user.Tel)
	fmt.Println("联系地址:", user.Addr)
	fmt.Println("备注:", user.Desc)
}

func Query(users map[int]User) {
	q := InputString("请输入查询内容:")
	s := InputString("请输入排序属性(id/name/bir/tel)：")
	userSlice := []User{}
	for _, v := range users {
		userSlice = append(userSlice, v)
	}

	sort.Slice(userSlice, func(i, j int) bool {
		switch s {
		case "id":
			return userSlice[i].ID < userSlice[j].ID
		case "name":
			return userSlice[i].Name < userSlice[j].Name
		case "bir":
			return userSlice[i].Birthday.Before(userSlice[i].Birthday)
		case "tel":
			return userSlice[i].Tel < userSlice[j].Tel
		default:
			return userSlice[i].ID < userSlice[j].ID
		}

	})
	fmt.Println("================================")
	for _, v := range userSlice {
		//name, birthday, tel, addr, desc
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Desc, q) {
			PrintUser(v.ID, v)
			fmt.Println("---------------------------------")
		}
	}
	fmt.Println("================================")
}

func GetId(users map[int]User) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func InputUser(users map[int]User, id int) User {
	user := User{}

	for {
		loop := false
		user.Name = InputString("请输入名字:")
		if user.Name == "" {
			fmt.Println("用户名不能为空")
			continue
		}

		for _, info := range users {
			if info.Name == user.Name {
				if users[id].Name != user.Name {
					fmt.Printf("%s 用户已存在\n", info.Name)
					loop = true
				}

			}
		}

		if !loop {
			break
		}
	}
	user.ID = id
	user.Birthday = InputTime("请输入出生日期(2000-01-01):")
	user.Tel = InputString("请输入联系方式:")
	user.Addr = InputString("请输入联系地址:")
	user.Desc = InputString("请输入备注:")
	return user
}

func Add(users map[int]User) {
	id := GetId(users)
	user := InputUser(users, id)
	users[id] = user
	fmt.Println("[+]添加成功")
}

func Modify(users map[int]User) {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			PrintUser(id, user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := InputUser(users, id)
				users[id] = user
				fmt.Println("[+]修改成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

func Del(users map[int]User) {
	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			PrintUser(id, user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("[+]删除成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}
