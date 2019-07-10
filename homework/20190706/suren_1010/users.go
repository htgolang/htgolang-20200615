package gopkg

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const PASSWD string = "123abc!@#"

func Auth() bool {
	var inputPass string
	passwd := md5.Sum([]byte(PASSWD))
	fmt.Println("欢迎使用马哥用户系统, 请输入管理员密码进入: ")

	for i := 3; i > 0; i-- {
		fmt.Scan(&inputPass)
		if md5.Sum([]byte(inputPass)) == passwd {
			return true
		} else if i == 1 {
			fmt.Println("密码输入错误, 请重试")
		} else {
			fmt.Printf("密码输入错误, 你还有%d次机会\n", i-1)
		}
	}
	return false
}

func Add(users map[int]map[string]string) {
	userInfo(getUid(users), users)
}

func Query(users map[int]map[string]string) {
	var keyword string

	fmt.Printf("请输入要查询的信息: ")
	fmt.Scan(&keyword)
	pirntTitle()
	for _, user := range users {
		if strings.Contains(user["name"], keyword) || strings.Contains(user["age"], keyword) || strings.Contains(user["tal"], keyword) || strings.Contains(user["addr"], keyword) {
			printUserInfo(user)
		}
	}
}

func Modify(users map[int]map[string]string) {
	var modifyID int
	var modifyStat string
	fmt.Printf("请输入要修改的用户ID: ")
	fmt.Scan(&modifyID)

	if checkUID(modifyID, users) == true {
		pirntTitle()
		printUserInfo(users[modifyID])
		fmt.Printf("请确认是否进行修改(yes/no): ")
		fmt.Scan(&modifyStat)
		switch modifyStat {
		case "yes":
			userInfo(modifyID, users)
			fmt.Println("用户修改成功！！！")
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func Del(users map[int]map[string]string) {
	var deleteID int
	var deleteStat string
	fmt.Printf("请输入要删除的用户ID: ")
	fmt.Scan(&deleteID)

	if checkUID(deleteID, users) == true {
		pirntTitle()
		printUserInfo(users[deleteID])
		fmt.Printf("请确认是否删除(yes/no): ")
		fmt.Scan(&deleteStat)
		switch deleteStat {
		case "yes":
			delete(users, deleteID)
			fmt.Println("用户删除成功！！！")
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func checkUID(uid int, users map[int]map[string]string) bool {
	if _, ok := users[uid]; ok {
		return true
	} else {
		fmt.Println("你输入的用户ID不存在, 请确认ID后, 在进行操作")
		return false
	}
}

func userInfo(uid int, users map[int]map[string]string) {
	var (
		name string
		age  string
		tal  string
		addr string
	)

	fmt.Printf("请输入名称: ")
	fmt.Scan(&name)

	fmt.Printf("请输入年龄: ")
	fmt.Scan(&age)

	fmt.Printf("请输入联系方式: ")
	fmt.Scan(&tal)

	fmt.Printf("请输入家庭住址: ")
	fmt.Scan(&addr)

	users[uid] = map[string]string{
		"id":   strconv.Itoa(uid),
		"name": name,
		"age":  age,
		"tal":  tal,
		"addr": addr,
	}
}

func pirntTitle() {
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%30s", "ID", "Name", "Age", "Tal", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

func printUserInfo(user map[string]string) {
	fmt.Printf("%5s|%20s|%5s|%20s|%30s\n", user["id"], user["name"], user["age"], user["tal"], user["addr"])
}

func getUid(users map[int]map[string]string) int {
	var uid int
	for k := range users {
		if k > uid {
			uid = k
		}
	}
	return uid + 1
}
