package users

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"github.com/pkg/errors"
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
	MaxAuth = 3
	pfile = ".password"
	info  = "info"
)

type User struct {
	Id       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

func (u User) String() string {
	return fmt.Sprintf("ID:%d\n名字:%s\n出生日期:%s\n联系方式:%s\n联系地址:%s\n备注:%s\n", u.Id, u.Name, u.Birthday.Format("2006/01/02"), u.Tel, u.Addr, u.Desc)
}


func init() {
	// 初始化info目录, 存储用户信息得数据
	err := os.MkdirAll(info, 0644)
	if err != nil {
		fmt.Println("info目录创建失败:", err)
	}
}


func InputString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}


// 认证
func Auth() bool {
	pwdFileExists := passwdFileExists()
	if pwdFileExists {
		// 如果文件存在, 读取文件, 开始认证
		pwd, err := os.Open(pfile)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			defer pwd.Close()
			password, _ := ioutil.ReadAll(pwd)
			// 判断文件内容是否为空， 空的话就需要重新设置密码，非空开始验证
			if string(password) == "" {
				setPasswd("初次使用, 请设置密码：", "请再次输入密码：")
			} else {
				return verifiPwd(string(password))
			}
		}
	} else {
		return setPasswd("初次使用, 请设置密码：", "请再次输入密码：")
	}
	return false
}


// 设置密码
func setPasswd(str1 string, str2 string) bool{
	fmt.Print(str1)
	pass1, _ := gopass.GetPasswd()
	fmt.Print(str2)
	pass2, _ := gopass.GetPasswd()

	if string(pass1) == string(pass2) {
		// 如果文件不存在, 提示初始化密码, 写入文件，开始认证
		pwd, err := os.Create(pfile)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			_, err := pwd.WriteString(fmt.Sprintf("%x", md5.Sum(pass1)))
			if err == nil {
				fmt.Println("密码已经设置完成, 请开始认证.")
				return verifiPwd(fmt.Sprintf("%x", md5.Sum(pass1)))
			} else {
				fmt.Println("密码设置失败： err")
				return false
			}
		}
	} else {
		fmt.Println("[-]两次密码输入不一样, 请重新运行.")
		return false
	}
}


// 判断密码文件是否存在
func passwdFileExists() bool {
	_, err := os.Stat(pfile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	} else {
		return true
	}
	return  false
}


// 验证密码
func verifiPwd(password string) bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码：")
		if inputpwd, err := gopass.GetPasswd(); err == nil {
			if fmt.Sprintf("%x", md5.Sum([]byte(inputpwd))) == string(password) {
				return true
			} else {
				fmt.Printf("密码错误%d次.\n", i+1)
			}
		}
	}
	fmt.Println("系统退出.")
	return false
}


func removeFile(lstfile []int) {
	if len(lstfile) <= 3 {
		return
	} else {
		// 排序 逆序
		sort.Slice(lstfile, func(i, j int) bool {
			return lstfile[i] > lstfile[j]
		})

		// 从第四个文件开始删除
		for i:=0; i<len(lstfile); i++ {
			if i >= 3 {
				os.Remove(filepath.Join(info, strconv.Itoa(lstfile[i])))
			}
		}
	}
}

// 获取用户文件
func getUserfile() string {
	lstfile := []int{}
	fileTimeUnix := 0
	infofile, _ := os.Open(info)
	files, err := infofile.Readdir(-1)
	if err == nil {
		defer infofile.Close()
		// 把文件加入文件列表
		for _, file := range files {
			f, err := strconv.Atoi(file.Name())
			if err != nil {
				fmt.Println(err)
			} else {
				lstfile = append(lstfile, f)
				if f > fileTimeUnix {
					fileTimeUnix = f
				}
			}
		}
	}
	// 保存最新3个文件，其他删除
	removeFile(lstfile)
	return filepath.Join(info, strconv.Itoa(fileTimeUnix))
}


// 反序列化
func gobdecode(file string) (map[int]User, error) {
	users := map[int]User{}
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在, 则创建文件
			ff, _ := os.Open(filepath.Join(info, file))
			defer ff.Close()
			return users, nil
		}
	}
	f, err := os.Open(file)
	if err == nil {
		defer f.Close()
		decoder := gob.NewDecoder(f)
		decoder.Decode(&users)
		return users, nil
	} else {
		return nil, fmt.Errorf("gobdecode 打开文件失败: %s", err)
	}

}


// 序列化
func goencode(file string, users map[int]User) {
	infofile, err := os.Create(file)
	if err != nil {
		fmt.Println("create file failed:", file)
		fmt.Println(err)
	} else {
		defer infofile.Close()
		gobencode := gob.NewEncoder(infofile)
		gobencode.Encode(users)
	}
}


// 获取用户id
func getUserID() (id int) {
	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	// 判断是否为空，如果为空，就返回1
	if len(users) == 0 {
		return 1
	}
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}


// 输入用户信息, 并且对用户名进行检查
func inputUser(id int) User {
	user := User{}
	user.Id = id
	user.Name = InputString("请输入姓名：")
	user.Birthday, _ = time.Parse("2006/01/02 ", InputString("请输入出生日期(1991/01/01)："))
	user.Tel = InputString("请输入联系电话：")
	user.Addr = InputString("请输入联系地址：")
	user.Desc = InputString("请输入备注：")
	return user
}


// 验证用户信息
func validUser(user User) error {
	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	if user.Name == "" {
		return fmt.Errorf("输入得用户名为空")
	}

	for _, tuser := range users {
		if user.Name == tuser.Name && user.Id != tuser.Id {
			return errors.New("输入得名称已经存在")
		}
	}
	return nil
}


// 排序
func sortUseLst(user_lst []User) []User {
	sortChoice := `请选择排序属性:
1. ID
2. 用户名
3. 出生日期
4. 联系方式
5. 备注
`
	if len(user_lst) == 0 || len(user_lst) == 1 {
		return user_lst
	} else {
		s := InputString(sortChoice)
		sort.Slice(user_lst, func(i, j int) bool {
			switch s {
			case "1":
				return user_lst[i].Id < user_lst[j].Id
			case "2":
				return user_lst[i].Name < user_lst[j].Name
			case "3":
				return user_lst[i].Birthday.Unix() < user_lst[j].Birthday.Unix()
			case "4":
				return user_lst[i].Addr < user_lst[j].Addr
			case "5":
				return user_lst[i].Desc < user_lst[j].Desc
			default:
				return user_lst[i].Id < user_lst[j].Id
			}
		})
	}
	return user_lst
}


// 查询用户信息
func Query() {
	user_lst := make([]User, 0)
	q := InputString("请输入查询的内容：")

	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	fmt.Println("=================================")
	for _, user := range users {
		if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) {
			user_lst = append(user_lst, user)
		}
	}

	user_lst = sortUseLst(user_lst)
	for _, user := range user_lst {
		fmt.Println(user)
		fmt.Println("-------------------------------")
	}
	fmt.Println("=================================")
}


// 添加用户
func Add() {
	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	id := getUserID()
	user := inputUser(id)
	if err := validUser(user); err == nil {
		users[id] = user
		// 序列化
		goencode(filepath.Join(info, strconv.FormatInt(time.Now().UnixNano(), 10)), users)
		fmt.Println("[+]添加成功")
	} else {
		fmt.Print("[-]添加失败:")
		fmt.Println(err)
	}

}


// 修改用户
func Modify() {
	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	if id, err := strconv.Atoi(InputString("请输入修改的ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改用户信息：")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := validUser(user); err == nil {
					users[id] = user
					goencode(filepath.Join(info, strconv.FormatInt(time.Now().UnixNano(), 10)), users)
					fmt.Println("[+]修改成功")
				} else {
					fmt.Print("[-]修改失败:")
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("[-]输入ID不存在")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}


// 删除用户
func Del() {
	users, err := gobdecode(getUserfile())
	if err != nil {
		fmt.Println("请检查gobdecode函数.")
	}

	if id, err := strconv.Atoi(InputString("请输入删除的ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除用户信息：")
			fmt.Println(user)
			input := InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				goencode(filepath.Join(info, strconv.FormatInt(time.Now().UnixNano(), 10)), users)
				fmt.Println("[+]删除成功")
			}
		} else {
			fmt.Println("[-]用户ID不存在.")
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}


// 修改密码
func MP() {
	fmt.Print("请输入原始密码：")
	if newpass, goerr := gopass.GetPasswd(); goerr == nil {
		f, err := os.Open(pfile)
		if err == nil {
			defer f.Close()
			password, _ := ioutil.ReadAll(f)
			if fmt.Sprintf("%x", md5.Sum(newpass)) != string(password) {
				fmt.Println("密码错误, 请重新输入指令.")
				return
			}
			issuccess := setPasswd("请设置新密码:", "请再次设置新密码:")
			if issuccess {
				fmt.Println("密码修改成功.")
			} else {
				fmt.Println("密码修改失败.")
				return
			}
		}
	}
}