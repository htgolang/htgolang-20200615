package users

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	MaxAuth      = 3
	passwordFile = ".password"
	userFile     = "user.json"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

// String方法
func (u User) String() string {
	return fmt.Sprintf("ID: %d\n名字:%s\n出生日期:%s\n联系方式:%s\n联系地址:%s\n备注:%s", u.ID, u.Name, u.Birthday.Format("2006-01-02"), u.Tel, u.Addr, u.Desc)
}

// 获取user
func inputUser(id int) User {
	var user User
	user.ID = id
	user.Name = InputString("请输入名字:")
	birthday, _ := time.Parse("2006-01-02", InputString("请输入出生日期(2000-01-01):"))
	user.Birthday = birthday
	user.Tel = InputString("请输入联系方式:")
	user.Addr = InputString("请输入联系地址:")
	user.Desc = InputString("请输入备注:")
	return user
}

// 验证用户
func vaildUser(user User) error {
	if user.Name == "" {
		return fmt.Errorf("输入的用户名为空")
	}
	users := loadUser()
	for _, tuser := range users {
		if user.Name == tuser.Name && user.ID != tuser.ID {
			return errors.New("输入的名字已经存在")
		}
	}
	return nil
}

// json反序列化 加载用户信息
func loadUser() map[int]User {
	users := map[int]User{}
	if file, err := os.Open(userFile); err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&users)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误：", err)
		}
	}
	return users
}

// 存储用户信息
func storeUser(users map[int]User) {
	if _, err := os.Stat(userFile); err == nil {
		os.Rename(userFile, string(strconv.FormatInt(time.Now().Unix(), 10)+".user.json"))
	} else {
		fmt.Println(err)
	}

	if names, err := filepath.Glob("*.user.json"); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}
	}

	if file, err := os.Create(userFile); err == nil {
		defer file.Close()
		encode := json.NewEncoder(file)
		encode.Encode(users)
	}
}

// 输入
func InputString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// 获取用户id
func getid() int {
	var id int
	users := loadUser()
	if len(users) == 0 {
		return 1
	}

	for k := range users {
		if k > id {
			id = k
		}
	}
	return id + 1
}

// 认证
func Auth() bool {
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil && len(password) > 0 {
		for i := 0; i < MaxAuth; i++ {
			fmt.Print("请输入密码:")
			bytes, _ := gopass.GetPasswd()
			if string(password) == fmt.Sprintf("%x", md5.Sum(bytes)) {
				return true
			} else {
				fmt.Println("[-]密码错误")
			}
		}
		return false
	} else {
		if len(password) == 0 || os.IsNotExist(err) {
			fmt.Println("请初始化密码：")
			bytes, _ := gopass.GetPasswd()
			ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			fmt.Println("[-]发生错误:", err)
			return false
		}
	}
}

// 查询用户
func Query() {
	q := InputString("请输入查询内容:")
	lst := make([]User, 0)
	fmt.Println("=======================")

	users := loadUser()
	for _, v := range users {
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Desc, q) {
			lst = append(lst, v)
		}
	}

	if len(lst) == 0 {
		fmt.Println("查询内容为空")
	} else {
		sortkey := InputString("请输入排序字段(id/name/tel/addr/desc):")
		sort.Slice(lst, func(i, j int) bool {
			switch sortkey {
			case "id":
				return lst[i].ID < lst[j].ID
			case "name":
				return lst[i].Name < lst[j].Name
			case "tel":
				return lst[i].Tel < lst[j].Tel
			case "addr":
				return lst[i].Addr < lst[j].Addr
			case "desc":
				return lst[i].Desc < lst[j].Desc
			default:
				return lst[i].ID < lst[j].ID
			}
		})

		for _, user := range lst {
			fmt.Println(user)
			fmt.Println("-----------------------------")
		}
	}
}

// 添加用户
func Add() {
	id := getid()
	user := inputUser(id)
	if err := vaildUser(user); err == nil {
		users := loadUser()
		users[id] = user
		storeUser(users)
		fmt.Println("[+]添加成功")
	} else {
		fmt.Print("[-]添加失败:")
		fmt.Println(err)
	}
}

// 修改用户
func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := vaildUser(user); err == nil {
					users[id] = user
					storeUser(users)
					fmt.Println("[+]修改成功")
				} else {
					fmt.Print("[-]修改失败:")
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

// 删除用户
func Del() {
	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			fmt.Println(user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := vaildUser(user); err == nil {
					delete(users, id)
					storeUser(users)
					fmt.Println("[+]删除成功")
				} else {
					fmt.Print("[-]删除失败:")
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

//修改密码
func MP() {
	if password, err := ioutil.ReadFile(passwordFile); err == nil {
		fmt.Print("请输入原密码:")
		bytes, _ := gopass.GetPasswd()
		if string(password) == fmt.Sprintf("%x", md5.Sum(bytes)) {
			fmt.Print("请输入新密码:")
			bytes, _ := gopass.GetPasswd()
			ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			fmt.Println("[+]密码修改成功")
		} else {
			fmt.Println("[-]密码错误")
		}
	}
}
