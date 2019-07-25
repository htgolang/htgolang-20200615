package users

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

// 定义常量maxAuth，为输入密码的最多次数
// 定义常量password，为密码的MD5加密值
const (
	maxAuth   = 3
	passwdgob = "passwd.gob"
	passwdcsv = "passwdfile.csv"
	user_file = "user.gob"
	passwddir = "passwd_dir"
	dir       = "users_dir"
)

var (
	// 定义变量Menu为操作的可选项
	Menu = `1. 显示
2. 查询
3. 添加
4. 修改
5. 删除
6. 退出
*******************************`

	// 定义变量显示所有的Users结构体类型的所有参数
	Sort_menu = `1. ID
2. Name
3. Birthday
4. Addr
5. Tel
6. Desc
*******************************`

	// 定义用户变量
	User          map[int]Users = map[int]Users{}
	passwdfilegob string        = filepath.Join(passwddir, passwdgob)
	passwdfilecsv string        = filepath.Join(passwddir, passwdcsv)
)

func Init() {
	logfile := "user.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		commandfile := os.Args[0]
		log.SetPrefix(commandfile + ":")
		log.SetFlags(log.Flags() | log.Lshortfile)
	}
}

// 在用户系统登录前，显示提示信息
func Title_String() {
	fmt.Println("JevonWei用户系统密码为:danran")
	fmt.Println("")

	// strings.Repeat() 显示特定字符多少次
	fmt.Println(strings.Repeat("*", 30))
	Head := "欢迎进入JevonWei的用户管理系统"
	fmt.Println(Head)
}

// 定义用户结构体类型Users

type Users struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Desc     string
}

// 打印用户信息
func (u Users) String() string {
	log.Printf("打印 %s 的信息\n", u.Name)
	return fmt.Sprintf("ID: %d\n名字: %s\n出生日期: %s\n联系方式: %s\n联系地址: %s\n备注: %s", u.ID, u.Name, u.Birthday.Format("2006/01/02"), u.Addr, u.Tel, u.Desc)
}

// 定义从键盘输入函数，并返回输入的值
func InputString(s string) string {
	var in string
	fmt.Print(s)
	fmt.Scan(&in)
	return strings.TrimSpace(in)
}

// 密码序列化
func Encodepasswd() {
	var password string

	// 判断密码目录是否存在
	_, err := os.Stat(passwddir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(passwddir, 0644)
			log.Printf("%s 目录已创建\n", passwddir)
		}
	}

	// 输入密码
	bytes, _ := gopass.GetPasswd()
	var password_first = fmt.Sprintf("%x", md5.Sum(bytes))

	// 输入确认密码
	fmt.Println("输入确认密码: ")
	bytes, _ = gopass.GetPasswd()
	password_confirm := fmt.Sprintf("%x", md5.Sum(bytes))

	// 当两次密码输入一致时，密码才能设置成功
	if password_first == password_confirm {
		password = password_confirm
	}

	// 创建密码gob文件，并写入
	file, err := os.Create(passwdfilegob)
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(password)
		log.Printf("passwd已写入%s文件\n", passwdfilegob)

		// 将输入的密码写入csv文件中
		csvwpasswd(password)
	}
}

// 密码反序列化
func Decodepasswd() string {
	var password string

	file, err := os.Open(passwdfilegob)
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&password)
		log.Printf("passwd已从%s文件读取\n", passwdfilegob)
		return password
	} else {
		fmt.Errorf("密码获取失败，请重新设置密码")
		log.Println("密码获取失败，请重新设置密码")
		return ""
	}
	return ""
}

// 密码序列化到csv文件
func csvwpasswd(password string) {

	file, err := os.Create(passwdfilecsv)

	if err == nil {
		defer file.Close()

		wirter := csv.NewWriter(file)
		wirter.Write([]string{password})

		wirter.Flush()
		log.Printf("passwd已序列化到%s文件\n", passwdfilecsv)
	}
}

// 密码序列化从csv文件反序列化
func csvrpasswd() []string {
	file, err := os.Open(passwdfilecsv)
	if err == nil {
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
				return []string{}
			} else {
				return line
				log.Printf("passwd已从%s文件反序列化\n", passwdfilecsv)
			}
			return []string{}
		}

	}
	return []string{}
}

// 列出目录下所有的用户文件
func Listfile() []string {
	flies := []string{}

	dirfile, err := os.Open(dir)
	if err == nil {
		defer dirfile.Close()
		childrens, _ := dirfile.Readdir(-1)

		for _, children := range childrens {
			if !children.IsDir() {
				flies = append(flies, children.Name())
			}
		}
	}
	return flies
}

// 获取最新的user文件中的数据及只保留最新的三个文件
func GetUserfile() []string {
	flies := Listfile()

	if len(flies) == 0 {
		return []string{}
	}

	if len(flies) <= 3 {
		//fmt.Printf("%T", flies[len(flies)-1:])
		return flies[len(flies)-1:]
	} else {
		for i := 0; i < len(flies)-2; i++ {
			os.Remove(filepath.Join(dir, flies[i]))
			log.Printf("%s文件已移除\n", filepath.Join(dir, flies[i]))
		}
	}
	return flies[len(flies)-1:]
}

// 用户信息序列化
func Userencode(users map[int]Users) {
	os.Mkdir(dir, 0644)
	file, err := os.Create(filepath.Join(dir, (strconv.FormatInt(time.Now().Unix(), 10) + user_file)))
	if err != nil {
		fmt.Printf("%s文件创建失败\n", file)
		fmt.Println(err)
		log.Printf("%s文件创建失败，err: %v", file, err)
	} else {
		defer file.Close()
		gobencode := gob.NewEncoder(file)
		gobencode.Encode(users)
	}
}

// 用户信息反序列化
func Userdecode() map[int]Users {
	var file *os.File
	var err error
	users := map[int]Users{}

	if len(GetUserfile()) == 0 {
		file, err = os.Create(filepath.Join(dir, (strconv.FormatInt(time.Now().Unix(), 10) + user_file)))
		log.Printf("用户文件不存在，创建了用户空文件: %v\n", file)
		return nil
	} else {
		file, err = os.Open(filepath.Join(dir, GetUserfile()[0]))
	}

	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		return users
	}
	return nil

}

// 键盘输入函数
func stdin() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return str
}

func Setepasswd() {
	_, err := os.Open(passwdfilegob)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("初次登陆或密码已丢失，请重新设置密码.")
			Encodepasswd()
			log.Println("密码设置成功")
		}
	}
}
func resetpasswd() bool {
	os.Stdout.Write([]byte("请输入原始密码："))

	bytes, _ := gopass.GetPasswd()
	//fmt.Printf("%#v", bytes)
	if bytes == nil {
		bytes, _ = gopass.GetPasswd()
	}

	if fmt.Sprintf("%x", md5.Sum(bytes)) == csvrpasswd()[0] {
		fmt.Println("请输入新密码")
		Encodepasswd()
		fmt.Println("密码更新成功")
		log.Println("用户密码更新成功")
		return true
	} else {
		fmt.Println("密码输入错误")
		log.Println("用户密码输入错误")
		return false
	}

}

func Auth() bool {
	Setepasswd()

	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入JevonWei用户系统密码: ")
		// 将输入的密码隐形显示
		bytes, _ := gopass.GetPasswd()

		// 如果输入密码的MD5值等于password，返回True
		if Decodepasswd() == fmt.Sprintf("%x", md5.Sum(bytes)) {
			in := InputString("是否要修改密码(Y/N):")
			if in == "Y" || in == "y" {
				for n := 0; n < maxAuth; n++ {
					if !resetpasswd() {
						if n == maxAuth-1 {
							fmt.Printf("密码输入%d次错误，密码修改程序退出\n", maxAuth)
						}

					} else {
						break
					}

				}
			}
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	fmt.Printf("密码输入%d次错误，程序退出\n", maxAuth)
	return false
}

// 将输入的字符串数据转换为时间类型，格式为2006-01-02，返回时间类型的值和错误信息
func Birthday_time(s string) (time.Time, error) {

	T, Err := time.Parse("2006-01-02", s)
	return T, Err

}

// 获取用户的最大ID，且返回ID+1
func GetId() int {
	Id := 0
	users := Userdecode()

	if len(users) == 0 {
		return 1
	}

	for k := range users {
		if Id < k {
			Id = k
		}
	}
	return Id + 1

}

// 定义用户输入函数，将从键盘输入的每个值对应传入Users结构体的元素中
func Inputuser(num int) {
	// 定义结构体类型Users的变量User_input
	var User_input Users = Users{}
	User := Userdecode()
	// 从键盘输入name(用户名)
	name := InputString("请输入名字:")

	// 若name为空，则不赋值，直接返回
	if name == "" {
		fmt.Println("输入的Name不能为空")
		goto END
	}

	// 若输入的name，在系统中已存在，则提示用户已存在，不能新建用户
	for _, user := range Userdecode() {
		if name == user.Name && user.ID != num {
			fmt.Printf("用户名%s已存在,不能新增/修改\n", name)
			log.Printf("用户名%s已存在,不能新增/修改\n", name)
			goto END
		}
	}

	// 将输入的name，赋值给Users结构体的Name
	User_input.Name = name
	// 将User映射的key值，赋值给结构体Users的ID
	User_input.ID = num

	// 将输入的字符串类型的Birthday转换为时间类型
	// 判断输入的Birthday格式是否正确，若输入格式有误，则打印提示信息，并重新输入
	for {
		birthday_time, err := Birthday_time(InputString("请输入出生日期(2019-07-07):"))
		if err == nil {
			User_input.Birthday = birthday_time
			break
		} else {
			fmt.Println(errors.New("请输入正确认格式"))
		}
	}

	// 将输入的其他值依次赋值为Tel，Addr，Desc,并返回用户信息
	User_input.Tel = InputString("请输入联系方式:")
	User_input.Addr = InputString("请输入地址:")
	User_input.Desc = InputString("请输入描述信息:")
	fmt.Println("*******************************")
	fmt.Printf("ID为%d的用户已添加/修改\n", num)

	User[num] = User_input
	Userencode(User)
	log.Printf("%v用户信息已添加或修改\n", User_input.Name)

END:
}

// 定义Listuser函数，打印用户系统中所有的用户
func Listuser() {
	user_slice := []string{}
	title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))

	// 遍历所有的用户，并打印
	for _, user := range Userdecode() {
		user_slice = append(user_slice, user.Name)
		fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
	}
	log.Printf("%v用户信息已遍历\n", user_slice)
}

// 添加函数
func Add() {
	// 获取用户ID
	id := GetId()

	// 调用用户函数，新增用户
	Inputuser(id)
	log.Printf("ID为 %d 的用户已添加\n", id)
}

// 删除用户
func Deluser() {
	// 当有用户数据文件时，方能删除用户
	// 当没有用户数据时，返回主菜单
	if len(GetUserfile()) > 0 {
		// 调用函数，显示系统中的所有用户
		Listuser()

		// 从gob文件中获取用户信息
		users := Userdecode()

		// 按照输入的用户ID，删除用户
		idString := InputString("请输入删除用户ID:")
		if id, err := strconv.Atoi(idString); err == nil {
			if user, ok := users[id]; ok {
				log.Printf("要删除的用户: %s\n", user.Name)
				fmt.Println("将要删除的用户信息为:")
				fmt.Println("================================")
				// 打印用户信息
				fmt.Println(user)

				// 确认是否删除用户
				in := InputString("是否确定删除(Y/N)?")
				if in == "Y" || in == "y" {
					delete(users, id)
					// 将users信息重新写入gob文件中
					Userencode(users)
					fmt.Printf("ID为%d的用户已删除\n", id)
					log.Printf("%s 用户已删除\n", user.Name)
				}
			} else {
				fmt.Println("输入的用户ID不存在")
				log.Printf("输入的ID不存在，用户删除失败")
			}
		} else {
			fmt.Println("输入的ID不正确")
			log.Printf("输入的ID不存在，用户删除失败")
		}
	} else {
		fmt.Println("当前没有用户数据，请先添加用户.")
	}

}

// 修改函数
func Modify() {
	// 当有用户数据文件时，方能修改用户
	// 当没有用户数据时，返回主菜单
	if len(GetUserfile()) > 0 {
		// 调用函数，显示系统中的所有用户
		Listuser()
		users := Userdecode()
		// 根据输入的用户ID，修改用户信息
		idString := InputString("请输入修改用户ID:")

		// 将输入的字符串类型的值转换为int类型
		if id, err := strconv.Atoi(idString); err == nil {
			if user, ok := users[id]; ok {
				log.Printf("要修改的用户为: %s\n", user.Name)
				fmt.Println("")
				fmt.Println("将要修改的用户信息为:")
				fmt.Println("================================")
				fmt.Println(user)

				in := InputString("是否确定修改(Y/N)?: ")
				if in == "Y" || in == "y" {
					// 输入用户信息
					Inputuser(id)
					log.Printf("%s 用户修改成功\n", user.Name)
				}
			} else {
				fmt.Println("输入的用户ID不存在")
				log.Printf("输入的ID不存在，用户修改失败\n")
			}
		} else {
			fmt.Println("输入的ID不正确")
			log.Printf("输入的ID不存在，用户修改失败\n")
		}
	} else {
		fmt.Println("当前没有用户数据，请先添加用户.")
	}

}

// 查询函数
func Query() {
	// 当有用户数据文件时，方能查询用户
	// 当没有用户数据时，返回主菜单
	if len(GetUserfile()) > 0 {
		q := InputString("请输入查询的信息:")

		title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
		fmt.Println(title)
		fmt.Println((strings.Repeat("-", len(title))))

		// 若输入的查询信息包含在Nmae、Desc、Addr任意一个参数中，则返回用户信息
		for _, user := range Userdecode() {
			if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) {
				log.Printf("查询的用户为: %s\n", user.Name)
				fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
			}
		}
	} else {
		fmt.Println("当前没有用户数据，请先添加用户.")
	}

}

// 定义排序函数 ，将用户按照指定的参数排序，并返回用户数组
func User_sort() []Users {

	// 定义空切片，存储用户信息
	Users_array := []Users{}

	// 输入按照哪个参数排序，并将输入的参数赋值给变量o
	o := InputString("请输入需要排序的键值:")

	//将系统中的用户保存在Users_array数组中
	for _, user := range Userdecode() {
		Users_array = append(Users_array, user)
	}

	if len(Users_array) == 0 || len(Users_array) == 1 {
		return Users_array
	} else {
		// 调用sort.Slice()函数，根据输入的排序参数，排序系统的所有用户
		sort.Slice(Users_array, func(i, j int) bool {
			switch o {
			case "1":
				return Users_array[i].ID < Users_array[j].ID
			case "2":
				return Users_array[i].Name < Users_array[j].Name
			case "3":
				return Users_array[i].Birthday.Format("2006/01/02") < Users_array[j].Birthday.Format("2006/01/02") // 将time类型的Birthday值转换为字符串排序
			case "4":
				return Users_array[i].Addr < Users_array[j].Addr
			case "5":
				return Users_array[i].Tel < Users_array[j].Tel
			case "6":
				return Users_array[i].Desc < Users_array[j].Desc
			default:
				return Users_array[i].ID < Users_array[j].ID
			}
		})
	}
	return Users_array
}

// 定义函数将排序后的用户打印
func Print_sort() {
	if len(GetUserfile()) > 0 {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println(Sort_menu)

		// 将排序后的用户数组赋值给list变量
		list := User_sort()

		// 打印排序后的用户
		//fmt.Println(list)
		for _, v := range list {
			fmt.Println(v)
			fmt.Println(strings.Repeat("*", 30))
		}
	} else {
		fmt.Println("当前没有用户数据，请先添加用户.")
	}

}
