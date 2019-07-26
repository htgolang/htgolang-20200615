package users

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

//初始化常量变量和结构体
const (
	log_file    = "users.log"
	db_file     = "users.db"
	passwd_file = ".passwd.file"
)

//用户类型的结构体
type User struct {
	Name     string
	Birthday time.Time
	Addr     string
	//Tel      string
	//Desc     string
}

//密码类型的结构体
type Passwd struct {
	Passwd string
}

var count int
var passwd = "123abc!@#"

//用户数据的结构体
var users = map[int]User{0: User{"Admin", time.Now(), "北京"}}

//系统方法3，记录日志,放在初始化里了
func init() {
	if file, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm); err == nil {
		log.SetOutput(file)
		log.SetPrefix("users:")
		log.SetFlags(log.Flags() | log.Lshortfile)
	} else {
		fmt.Println(err)
	}
}

//输出测试，编写过程中，调试用
func Test() {
	//fmt.Printf("users映射数据类型: %v\n", users)
	//p.SetPasswd()
	fmt.Println(passwd)
	//Add()
	//Update()
	//Query()
	Del()
	//fmt.Printf("下次新增ID为: %v ", GetID())
	//FileEncode(db_file,users)
	//fmt.Println(users)
}

//系统方法1，认证，N为true时，跳过密码。
func Auth(N bool) bool {
	p := Passwd{}
	if !p.ReaderFile() {
		p.SetPasswd()
	}

	if N != true {
		for {
			if count < 3 {
				input_passwd := fmt.Sprintf("%X", md5.Sum([]byte(Inputstring("请输入密码: "))))
				if input_passwd == passwd {
					return true
				}
				count++
				fmt.Println("密码错误，请重试: ")
			} else {
				fmt.Println("尝试次数过多，退出管理系统")
				return false
			}
		}
	} else {
		return true
	}
}

//基本方法1，输入信息
func Inputstring(tips string) string {
	var Ss string
	fmt.Print(tips)
	fmt.Scan(&Ss)
	return strings.TrimSpace(Ss)
}

//系统方法2，密码文件的md5转换存储和读取

func (p *Passwd) ReaderFile() bool {
	if bytes, err := ioutil.ReadFile(passwd_file); err == nil {
		passwd = string(bytes)
		fmt.Println(passwd)
	} else {
		fmt.Println(err)
		return false
	}
	return true
}

func (p *Passwd) WriteFile() {
	if file, err := os.Create(passwd_file); err == nil {
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(passwd)
		writer.Flush()
	} else {
		fmt.Println(err)
	}
}

func (p *Passwd) SetPasswd() {
	fmt.Print("请设置密码: ")
	bytes, _ := gopass.GetPasswd()
	fmt.Println(string(bytes))
	first := fmt.Sprintf("%X", md5.Sum(bytes))

	fmt.Print("请确认密码: ")
	bytes, _ = gopass.GetPasswd()
	fmt.Println(string(bytes))
	confirm := fmt.Sprintf("%X", md5.Sum(bytes))

	if first == confirm {
		passwd = confirm
		p.WriteFile()
		log.Printf("设置密码成功")
		fmt.Println("设置密码成功")
	} else {
		fmt.Println("两次输入的密码不一致，设置密码失败")
	}
}

//系统方法2，排序，

//系统方法4，用户数据的序列化和反序列化
func FileEncode(db_file string, users map[int]User) {
	if file, err := os.Create(db_file); err == nil {
		defer file.Close()
		writer := gob.NewEncoder(file)
		writer.Encode(users)
	} else {
		fmt.Println(err)
	}
}

func FileDecode(db_file string) map[int]User {
	if file, err := os.Open(db_file); err == nil {
		defer file.Close()
		reader := gob.NewDecoder(file)
		reader.Decode(&users)
		return users
	} else {
		fmt.Println(err)
	}
	return nil
}

//系统方法6，欢迎菜单
//系统方法7，获取最大ID
func GetID() int {
	ID := 0
	for k, _ := range users {
		if ID < k {
			ID = k
		}
	}
	return ID + 1
}

//用户方法1，写入用户数据
func (user *User) SetUser() {
	user.Name = Inputstring("请输入姓名: ")
	user.Birthday = time.Now()
	user.Addr = Inputstring("请输入地址: ")
	//user.Tel = Inputstring("请输入电话: ")
	//user.Desc = Inputstring("请输入备注信息: ")
	//fmt.Println(user)
}

//数据方法1，增加,
func Add() {
	defer FileEncode(db_file, users)
	users := FileDecode(db_file)
	//fmt.Println(users)
	user := User{}
	user.SetUser()
	ID := GetID()
	users[ID] = user
	log.Printf("添加用户 %v", user.Name)
	fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)

}

//数据方法2，修改
func Update() {
	defer FileEncode(db_file, users)
	users := FileDecode(db_file)

	if ID, err := strconv.Atoi(Inputstring("请输入想要修改的用户ID: ")); err == nil {
		if user, ok := users[ID]; ok {
			fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
			user.SetUser()
			users[ID] = user
			log.Printf("更新用户 %v", user.Name)
			fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
		} else {
			fmt.Println("用户ID不存在.")
		}
	} else {
		fmt.Println("请输入正确的用户的ID.")
	}
	//fmt.Println(users)
}

//数据方法3，删除
func Del() {
	defer FileEncode(db_file, users)
	users := FileDecode(db_file)
	fmt.Println(users)

	if ID, err := strconv.Atoi(Inputstring("请输入想要删除的用户ID: ")); err == nil {
		if user, ok := users[ID]; ok {
			fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
			exec := Inputstring("请问你确定要删除这个用户吗？ Y/N: ")
			if exec == "Y" || exec == "y" {
				delete(users, ID)
				log.Printf("删除用户 %v", user.Name)
				fmt.Println("用户已经删除.")
			} else {
				fmt.Println("退出用户删除操作.")
			}
		} else {
			fmt.Println("输入的用户ID不存在.")
		}
	} else {
		fmt.Println(err)
	}
}

//数据方法4，查询
func Query() {
	defer FileEncode(db_file, users)
	users := FileDecode(db_file)

	q := Inputstring("请输入想要查询的内容: ")
	if len(q) != 0 {
		for ID, user := range users {
			if q == "all" {
				log.Println("查询用户 all")
				fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
				//fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
			} else if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) {
				log.Printf("查询用户 %v", user.Name)
				fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
			}
		}
	} else {
		fmt.Println("输入的查询内容为空")
	}
}
