package users

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"crypto/md5"

	"github.com/howeyc/gopass"
)

const (
	MaxAuth      = 3
	passwordFile = ".password"
	userFile     = "users.gob"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\n名字:%s\n出生日期:%s\n联系方式:%s\n联系地址:%s\n备注:%s", u.ID, u.Name, u.Birthday.Format("2006-01-02"), u.Tel, u.Addr, u.Desc)
}

func loadUser() map[int]User {
	users := map[int]User{}
	if file, err := os.Open(userFile); err == nil {
		defer file.Close()
		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误: ", err)
		}
	}
	return users
}

func storeUser(users map[int]User) {
	if file, err := os.Create(userFile); err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
	}
}

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

// 从命令行输入密码, 并进行验证
// 通过返回值告知验证成功还是失败
func Auth() bool {
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil && len(password) > 0 {
		// 验证密码
		for i := 0; i < MaxAuth; i++ {
			fmt.Print("请输入密码:")
			// fmt.Scan(&input)
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
			// 初始化密码
			fmt.Print("请输入初始化密码:")
			bytes, _ := gopass.GetPasswd()
			ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			// 其他原因导致错误
			fmt.Println("[-]发生错误", err)
			return false
		}
	}
}

func Query() {
	q := InputString("请输入查询内容:")
	list := make([]User, 0)
	fmt.Println("================================")

	users := loadUser()

	for _, v := range users {
		//name, birthday, tel, addr, desc
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Desc, q) {
			list = append(list, v)
		}
	}
	if len(list) == 0 {
		fmt.Println("查询内容为空")
	} else {
		sortKey := InputString("请输入排序字段(id/name/tel/addr/desc):")
		sort.Slice(list, func(i, j int) bool {
			switch sortKey {
			case "id":
				return list[i].ID < list[j].ID
			case "name":
				return list[i].Name < list[j].Name
			case "tel":
				return list[i].Tel < list[j].Tel
			case "addr":
				return list[i].Addr < list[j].Addr
			case "desc":
				return list[i].Desc < list[j].Desc
			default:
				return list[i].ID < list[j].ID
			}
		})

		for _, user := range list {
			fmt.Println(user)
			fmt.Println("----------------------------")
		}
	}

	fmt.Println("================================")
}

func getId() int {
	var id int
	users := loadUser()
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser(id int) User {
	// user := User{}
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

func validUser(user User) error {
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

func Add() {
	id := getId()
	user := inputUser(id)
	if err := validUser(user); err == nil {
		users := loadUser()
		users[id] = user
		storeUser(users)
		fmt.Println("[+]添加成功")
	} else {
		fmt.Print("[-]添加失败:")
		fmt.Println(err)
	}
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := validUser(user); err == nil {
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

func Del() {
	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			fmt.Println(user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				storeUser(users)
				fmt.Println("[+]删除成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}
