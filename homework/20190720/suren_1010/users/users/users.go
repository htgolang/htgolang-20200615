package users

import (
	"crypto/md5"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tal      string
	Addr     string
}

var users = make(map[int]User)

//const PASSWD string = "123abc!@#"
//const PASSWD string = "8c13b5750412d922b01b2da95d24f8b6"
const PasswdFile string = "passwd.gob"

func Auth() bool {
	var inputPass string
	passwd := file2passwd()
	fmt.Println("欢迎使用马哥用户系统, 请输入管理员密码进入: ")

	for i := 3; i > 0; i-- {
		fmt.Scan(&inputPass)
		inputPass = fmt.Sprintf("%x", md5.Sum([]byte(inputPass)))
		if inputPass == passwd {
			return true
		} else if i == 1 {
			fmt.Println("密码输入错误, 请重试")
		} else {
			fmt.Printf("密码输入错误, 你还有%d次机会\n", i-1)
		}
	}
	return false
}

func Add() {
	uid := getUid()
	user := userInfo(uid)
	if err := checkName(user); err != nil {
		fmt.Println(err)
	} else {
		users[uid] = user
		fmt.Println("添加成功")
	}
}

func Query() {
	keyword := InputString("请输入要查询的信息: ")
	list := make([]User, 0)
	for _, user := range users {
		if strings.Contains(user.Name, keyword) || strings.Contains(user.Tal, keyword) || strings.Contains(user.Addr, keyword) {
			list = append(list, user)
		}
	}
	if len(list) == 0 {
		fmt.Println("查询内容为空")
	} else {
		sortkey := InputString("请输入要排序的字段: id/name/tal/addr: ")
		sort.Slice(list, func(i, j int) bool {
			switch sortkey {
			case "id":
				return list[i].ID > list[j].ID
			case "name":
				return list[i].Name > list[j].Name
			case "tal":
				return list[i].Tal > list[j].Tal
			case "addr":
				return list[i].Addr > list[j].Addr
			default:
				return list[i].ID < list[j].ID
			}
		})
		pirntTitle()
		for _, user := range list {
			fmt.Println(user)
		}
	}
}

func Modify() {
	uid, _ := strconv.Atoi(InputString("请输入要修改的用户ID: "))
	if checkUID(uid) == true {
		pirntTitle()
		printUserInfo(users[uid])
		switch InputString("请确认是否进行修改(yes/no): ") {
		case "yes":
			if err := checkName(users[uid]); err != nil {
				fmt.Println(err)
			} else {
				users[uid] = userInfo(uid)
				fmt.Println("用户修改成功！！！")
			}
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func Del() {
	uid, _ := strconv.Atoi(InputString("请输入要删除的用户ID: "))
	if checkUID(uid) == true {
		pirntTitle()
		printUserInfo(users[uid])
		switch InputString("请确认是否删除(yes/no): ") {
		case "yes":
			delete(users, uid)
			fmt.Println("用户删除成功！！！")
		case "no":
			break
		default:
			fmt.Println("请输入yes or no")
		}
	}
}

func checkName(user User) error {
	if user.Name == "" {
		return fmt.Errorf("输入的用户名不能为空")
	}
	for _, u := range users {
		if user.Name == u.Name && user.ID != u.ID {
			return errors.New("输入的名字已存在")
		}
	}
	return nil
}

func checkUID(uid int) bool {
	if _, ok := users[uid]; ok {
		return true
	} else {
		return false
	}
}

func userInfo(uid int) User {
	var user User
	user.ID = uid
	user.Name = InputString("请输入名称: ")
	user.Birthday, _ = time.Parse("2006-01-02", InputString("请输入生日(年-月-日): "))
	user.Tal = InputString("请输入联系方式: ")
	user.Addr = InputString("请输入家庭住址: ")
	return user
}

func InputString(pam string) string {
	var input string
	fmt.Printf(pam)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func pirntTitle() {
	title := fmt.Sprintf("%5s|%20s|%20s|%20s|%30s", "ID", "Name", "Tal", "Birthday", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

func (u User) String() string {
	return fmt.Sprintf("%5d|%20s|%20s|%20s|%30s\n", u.ID, u.Name, u.Tal, u.Birthday.Format("2006-01-02"), u.Addr)

}

func printUserInfo(user User) {
	fmt.Printf("%5d|%20s|%20s|%20s|%30s\n", user.ID, user.Name, user.Tal, user.Birthday.Format("2006-01-02"), user.Addr)
}

func getUid() int {
	uid := 0
	for k := range users {
		if k > uid {
			uid = k
		}
	}
	return uid + 1
}

func PasswdFileCheck() error {
	fInfo, err := os.Stat(PasswdFile)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("文件不存在, 请初始化密码")
		}
	} else if fInfo.Size() == 0 {
		return errors.New("密码为空, 请初始化密码")
	}
	return nil
}

func Passwd2File(passwd string) {
	file, err := os.Create(PasswdFile)
	if err == nil {
		defer file.Close()
		encode := gob.NewEncoder(file)
		encode.Encode(passwd)
	}
}

func file2passwd() string {
	var passwd string
	file, err := os.Open(PasswdFile)
	if err == nil {
		defer file.Close()
		decode := gob.NewDecoder(file)
		decode.Decode(&passwd)
	}
	return passwd
}
