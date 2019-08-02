package users

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
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
	json_file   = "users_json.db"
	passwd_file = ".passwd.file"
)

//用户类型的结构体
type User struct {
	ID       int       `json:id`
	Name     string    `json:name`
	Birthday time.Time `json:birthday`
	Addr     string    `json:addr`
	//Tel      string    `json:tel`
	//Desc     string    `json:desc`
}

//密码类型的结构体
type Passwd struct {
	Passwd string
}

var count int
var passwd string

//用户数据的结构体
//var users Users = Users{0, "Admin", time.Now(), "北京"}
var data Data = JsonFile{}
var users = map[int]User{}
var ID int
var user = User{}

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
		fmt.Println("请先设置您的密码，只有正确设置密码后，才可以使用该系统。")
		for {
			if p.SetPasswd() {
				break
			}
		}
	}
	//通过N判断，是否需要验证密码
	if N == true {
		return true
	} else {
		return p.CheckPasswd()
	}
}

//基本方法1，输入信息
func Inputstring(tips string) string {
	//var Ss string
	fmt.Print(tips)
	//fmt.Scan(&Ss)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

//系统方法2，密码文件的md5转换存储和读取

func (p *Passwd) CheckPasswd() bool {
	for {
		if count < 3 {
			fmt.Print("请输入系统登录密码: ")
			bytes, _ := gopass.GetPasswd()
			inputpasswd := fmt.Sprintf("%X", md5.Sum(bytes))
			if inputpasswd == passwd {
				return true
			}
			count++
			fmt.Println("密码错误，请重试: ")
		} else {
			return false
		}
	}
}

func (p *Passwd) ReaderFile() bool {
	if bytes, err := ioutil.ReadFile(passwd_file); err == nil {
		passwd = string(bytes)
	} else {
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

func (p *Passwd) SetPasswd() bool {
	fmt.Print("请设置密码: ")
	bytes, _ := gopass.GetPasswd()
	//fmt.Println(string(bytes))
	first := fmt.Sprintf("%X", md5.Sum(bytes))

	fmt.Print("请确认密码: ")
	bytes, _ = gopass.GetPasswd()
	//fmt.Println(string(bytes))
	confirm := fmt.Sprintf("%X", md5.Sum(bytes))

	if first == confirm {
		passwd = confirm
		p.WriteFile()
		log.Printf("设置密码成功")
		fmt.Println("设置密码成功")
		return true
	} else {
		fmt.Println("两次输入的密码不一致，设置密码失败")
		return false
	}
}

/*
//系统方法4，用户数据的序列化和反序列化
func FileEncode(db_file string, users Users) {
	if file, err := os.Create(db_file); err == nil {
		defer file.Close()
		writer := gob.NewEncoder(file)
		writer.Encode(users)
	} else {
		fmt.Println(err)
	}
}
func FileDecode(db_file string) map[int]Users {
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
*/

type Data interface {
	WriteFile(file string, users map[int]User) error
	LoadFile(file string) (map[int]User, error)
}

type JsonFile struct{}

func (w JsonFile) WriteFile(file string, users map[int]User) error {
	if bytes, err := json.MarshalIndent(users, "", "\t"); err == nil {
		if err := ioutil.WriteFile(file, bytes, os.ModePerm); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (w JsonFile) LoadFile(file string) (map[int]User, error) {
	if file, err := os.Open(file); err == nil {
		defer file.Close()
		if bytes, err := ioutil.ReadAll(file); err == nil {
			json.Unmarshal(bytes, &users)
			return users, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func GetID() int {
	defer data.WriteFile(json_file, users)
	users, _ := data.LoadFile(json_file)

	for k, _ := range users {
		if ID < k {
			ID = k
		}
	}
	return ID + 1
}

func (user *User) SetUser(ID int) {
	user.ID = GetID() - 1
	user.Name = Inputstring("请输入姓名: ")
	user.Birthday = time.Now()
	user.Addr = Inputstring("请输入地址: ")
	//user.Tel = Inputstring("请输入电话: ")
	//user.Desc = Inputstring("请输入备注信息: ")
	//fmt.Println(u)
}

func (user User) String() string {
	return fmt.Sprintf("%-5d|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
}

func Add() {
	defer data.WriteFile(json_file, users)
	users, _ := data.LoadFile(json_file)

	ID = GetID()
	user.SetUser(ID)
	users[ID] = user
	log.Printf("添加用户 %v", users[GetID()].Name)
	user.String()
}

//数据方法2，修改
func Update() {
	defer data.WriteFile(json_file, users)
	users, _ := data.LoadFile(json_file)
	//defer FileEncode(db_file, users)
	//users := FileDecode(db_file)

	if ID, err := strconv.Atoi(Inputstring("请输入想要修改的用户ID: ")); err == nil {
		if user, ok := users[ID]; ok {
			user.String()
			//fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
			user.SetUser(ID)
			users[ID] = user
			log.Printf("更新用户 %v", user.Name)
			user.String()
			//fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
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
	defer data.WriteFile(json_file, users)
	users, _ := data.LoadFile(json_file)
	//defer FileEncode(db_file, users)
	//users := FileDecode(db_file)
	//fmt.Println(users)

	if ID, err := strconv.Atoi(Inputstring("请输入想要删除的用户ID: ")); err == nil {
		if user, ok := users[ID]; ok {
			user.String()
			//fmt.Printf("%-5d|%-10s|%-15s|%-15s\n", ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
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
	defer data.WriteFile(json_file, users)
	users, _ := data.LoadFile(json_file)

	//fmt.Println(users)
	q := Inputstring("请输入想要查询的内容: ")
	if len(q) != 0 {
		for ID, user := range users {
			if q == "all" {
				log.Println("查询用户 all")
				//user.String()
				fmt.Printf("%-5d|%-5d|%-10s|%-15s|%-15s\n", ID, user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
				//fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
			} else if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) {
				log.Printf("查询用户 %v", user.Name)
				//fmt.Println(user)
				//user.String()
				fmt.Printf("%-5d|%-5d|%-10s|%-15s|%-15s\n", ID, user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Addr)
			}
		}
	} else {
		fmt.Println("输入的查询内容为空")
	}
}

func UpdatePasswd() {
	p := Passwd{}
	if p.CheckPasswd() {
		p.SetPasswd()
	} else {
		fmt.Println("您输入的密码不正确，退出修改密码模式")
	}
}
