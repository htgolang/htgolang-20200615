package users

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
	"github.com/pkg/errors"
)

const (
	MaxAuth      = 3
	passwordFile = ".password"
	userdir      = "userdata"
	userFile     = "users"
)

type User struct {
	ID       int       `json:"id""`
	Name     string    `json:"name""`
	Birthday time.Time `json:"birthday"`
	Tel      string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
}

func init() {
	_, err := os.Stat(userdir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(userdir, 0755)
		} else {
			fmt.Println("[-]发生错误：", err)
		}
	}
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\n名字:%s\n出生日期:%s\n联系方式:%s\n联系地址:%s\n备注:%s", u.ID, u.Name, u.Birthday.Format("2006-01-02"), u.Tel, u.Addr, u.Desc)
}

//var persistence Persistence = NewJSONFile(filepath.Join(userdir, userFile))
var persistence Persistence

func SetPersistence(type_ string) {
	fmt.Println("type_:", type_)
	switch type_ {
	case "json":
		persistence = NewJSONFile(filepath.Join(userdir, userFile))
	case "gob":
		persistence = NewGOBFile(filepath.Join(userdir, userFile))
	default:
		persistence = NewGOBFile(filepath.Join(userdir, userFile))
	}
}

func InputString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func Auth() bool {
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil && len(passwordFile) > 0 {
		for i := 0; i < MaxAuth; i++ {
			fmt.Print("请输入密码：")
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
			fmt.Print("请输入初始化密码：")
			bytes, _ := gopass.GetPasswd()
			ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			fmt.Println("[-]发生错误:", err)
			return false
		}
	}
}

func Query() {
	q := InputString("请输入查询内容:")
	list := make([]User, 0)
	fmt.Println("============================")

	users, err := persistence.Load()
	if err != nil {
		fmt.Println("[-]发生错误: ", err)
		return
	}

	for _, v := range users {
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
	}
}

func getId() (int, error) {
	var id int
	users, err := persistence.Load()
	if err != nil {
		fmt.Println("[-]发生错误:", err)
		return -1, err
	}
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1, nil
}

func inputUser(id int) User {
	var user User
	user.ID = id
	user.Name = InputString("请输入名字：")
	birthday, _ := time.Parse("2006-01-02", InputString("请输入出生日期(2000-01-01):"))
	user.Birthday = birthday
	user.Tel = InputString("请输入联系方式:")
	user.Addr = InputString("请输入联系地址:")
	user.Desc = InputString("请输入备注:")
	return user
}

func vaildUser(user User) error {
	if user.Name == "" {
		return fmt.Errorf("输入的用户名为空.")
	}
	users, err := persistence.Load()
	if err != nil {
		fmt.Println("[-]发生错误：", err)
		return err
	}

	for _, tuser := range users {
		if user.Name == tuser.Name && user.ID != tuser.ID {
			return errors.New("输入的名字已经存在.")
		}
	}
	return nil
}

func Add() {
	id, err := getId()
	if err != nil {
		fmt.Println("[-]发生错误:", err)
		return
	}

	user := inputUser(id)
	if err := vaildUser(user); err == nil {
		users, err := persistence.Load()
		if err != nil {
			fmt.Println("[-]发生错误: ", err)
			return
		}
		users[id] = user
		if err := persistence.Store(users); err == nil {
			fmt.Println("[+]添加成功")
		} else {
			fmt.Println("[-]添加失败, 发生错误: ", err)
		}
	} else {
		fmt.Print("[-]添加失败:")
		fmt.Println(err)
	}
}

func Modify() {
	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		users, err := persistence.Load()
		if err != nil {
			fmt.Println("[-]发生错误: ", err)
			return
		}
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息:")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := vaildUser(user); err == nil {
					users[id] = user
					if err := persistence.Store(users); err == nil {
						fmt.Println("[+]修改成功")
					} else {
						fmt.Println("[-]修改失败, 发生错误: ", err)
					}
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
		users, err := persistence.Load()
		if err != nil {
			fmt.Println("[-]发生错误: ", err)
			return
		}
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息:")
			fmt.Println(user)
			input := InputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				if err := persistence.Store(users); err == nil {
					fmt.Println("[+]删除成功")
				} else {
					fmt.Println("[-]删除失败, 发生错误: ", err)
				}
			}
		} else {
			fmt.Println("[-]用户ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

