package usersmem

import (
	"crypto/md5"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	MaxAuth  = 3
	//password = "e10adc3949ba59abbe56e057f20f883e"
    pwdfile = ".password"
    //userfile = "userinfo.gob"
)
var users = map[int]User{}

type User struct {
	ID int
	Name string
	Birthday time.Time
	tel string
	addr string
	desc string
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\n名字:%s\n出生日期:%s\n联系方式:%s\n联系地址:%s\n备注:%s", u.ID, u.Name, u.Birthday.Format("2006-01-02"), u.tel, u.addr, u.desc)
}

func init(){
	users = make(map[int]User)
}

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func checkFile(file string) bool {

	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false

		}
	}
	return true

}

// 从命令行输入密码, 并进行验证
// 通过返回值告知验证成功还是失败
func Auth() bool {

	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码进行认证: ")
		// fmt.Scan(&input)
		bytes, _ := gopass.GetPasswd()

		file := checkFile(pwdfile)


		if file {

			passfile, err := os.Open(pwdfile)

			if err != nil {
				fmt.Println(err)

			}else {

				defer passfile.Close()

				password, _ := ioutil.ReadAll(passfile)

				//		pwd := strings.Replace(string(password),"\n", "", -1)
				pwd := string(password)


				if pwd == "" {
					setPassword("","")

				}else {

					if pwd == fmt.Sprintf("%x", md5.Sum(bytes)) {

						return true

					} else {
						fmt.Println("密码错误")
					}

				}
			}
		}else {
			fmt.Println("密码不存在，请初始化密码！")
			setPassword("","")
		}


	}
	return true


}

// 设置密码
func setPassword(pass1, pass2 string) bool {
	fmt.Println("请设置密码: ")

	p1, _ := gopass.GetPasswd()
	fmt.Println("请再次输入密码: ")
	p2,_ := gopass.GetPasswd()

	if string(p1) == string(p2) {
		passfile, err := os.Create(pwdfile)
		if err != nil {
			fmt.Println(err)

		}else {
			_, err := passfile.WriteString(fmt.Sprintf("%x", md5.Sum(p1)))
			if err != nil {
				//fmt.Println("密码设置成功.")
				return true

			}

		}
		return false
	}else {
		fmt.Println("两次输入不一致，请重新输入！")
		return false
	}

	return false

}

// 获取用户信息文件
func getUserFile() string {

	filelist := []string{}

	gobpath,_ := filepath.Glob("*.gob")


	for _,file := range gobpath{
		filelist = append(filelist, file)
	}
	if len(filelist) >0 {
		sort.Sort(sort.Reverse(sort.StringSlice(filelist)))

		if len(filelist) > 3 {
			for i, file := range filelist[3:] {
				os.Remove(file)
				filelist = append(filelist[:i], file[i+1:])

			}

		}
		file := string(filelist[0])
		fmt.Println(file)
		return file
	}else {
		filename := filepath.Join(strconv.FormatInt(time.Now().UnixNano(), 10))
		filename = filename + ".gob"

		return filename
	}


}


// 反序列化
func gobdecode(file string) (map[int]User, error) {
	users := map[int]User{}

	tmpfile := checkFile(file)

	if tmpfile {
		f, err := os.Open(file)

		if err == nil{
			defer f.Close()

			decoder := gob.NewDecoder(f)
			decoder.Decode(&users)
			return users, nil
		}else {
			return nil, fmt.Errorf("打开文件失败")
		}

	}else {
		ff, _:= os.Open(file)
		defer ff.Close()
		return users, nil
	}


	return users, nil
}


// 序列化
func goencode(file string, users map[int]User) {

	tmpfile, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
	}else{
		defer tmpfile.Close()
		gobdecode := gob.NewEncoder(tmpfile)
		gobdecode.Encode(users)
	}

}


func printUser(user User) {
	fmt.Println("ID:", user.ID)
	fmt.Println("名字:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006-01-02"))
	fmt.Println("联系方式:", user.tel)
	fmt.Println("联系地址:", user.addr)
	fmt.Println("备注:", user.desc)
}

func sortUser(user_list []User) []User {
	var op string
	fmt.Print(`请输入排序字段:
	1: ID
	2: Name
	3: Birthday
	4: tel
`)

	fmt.Scan(&op)

	sort.Slice(user_list, func(i, j int) bool {
		switch op {
		case "1":
			return user_list[i].ID < user_list[j].ID
		case "2":
			return user_list[i].Name < user_list[j].Name
		case "3":
			return user_list[i].Birthday.Before(user_list[i].Birthday)
		case "4":
			return user_list[i].tel < user_list[j].tel
		default:
			return user_list[i].ID < user_list[j].ID
		}
	})
	return user_list
}

// 获取ID
func getId() int {
	var id int

	users, err := gobdecode(getUserFile())
	if err != nil {
		fmt.Println(err)
	}
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

func validUser(user User) error {
	users, err := gobdecode(getUserFile())
	if err != nil {
		fmt.Println(err)
	}
	if user.Name == "" {
		return fmt.Errorf("输入的用户名为空")
	}
	for _, tuser := range users {
		if user.Name == tuser.Name && user.ID != tuser.ID {
			return errors.New("输入的名字已经存在")
		}
	}
	return nil
}

// 用户输入
func inputUser(id int) User {
	var user User
	user.ID = id
	user.Name = InputString("请输入名字: ")
	user.Birthday, _ = time.Parse("2006-01-02",InputString("请输入出生日期(2006-01-02): "))
	user.tel = InputString("请输入联系方式: ")
	user.addr = InputString("请输入联系地址: ")
	user.desc = InputString("请输入备注: ")
	return user
}

// 添加用户
func Add() {

	users, err := gobdecode(getUserFile())
	fmt.Println(users)

	fmt.Println(users)
	if err != nil {
		fmt.Println(err)
	}

	id := getId()
	user := inputUser(id)

	if err := validUser(user); err == nil {
		users[id] = user
		filename := filepath.Join(strconv.FormatInt(time.Now().UnixNano(), 10))
		filename = filename + ".gob"
		goencode(filename, users)
		fmt.Println(users)

		fmt.Println("[+]添加成功")
	}

}

// 查询用户
func Query() {
	tmp := make([]User, 0)
	q := InputString("请输入查询内容:")
	users, err := gobdecode(getUserFile())


	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("================================")
	for _, v := range users {
		//name, birthday, tel, addr, desc
		if strings.Contains(v.Name, q) || strings.Contains(v.tel, q) || strings.Contains(v.addr, q) || strings.Contains(v.desc, q) {
			tmp = append(tmp, v)

		}
	}
	tmp = sortUser(tmp)
	for _, v := range tmp {
		printUser(v)
		fmt.Println("++++++++++++++++++++++++++")

	}
	fmt.Println("================================")
}

func Modify() {

	users, err:= gobdecode(getUserFile())

	if err != nil {
		fmt.Println(err)
	}

	if id, err := strconv.Atoi(InputString("请输入修改用户ID:")); err == nil {
		for _, v := range users {
			if v.ID == id {
				fmt.Println("将修改的用户信息:")
				printUser(v)
				input := InputString("确定修改(Y/N)?")
				if input == "y" || input == "Y" {
					user := inputUser(id)
					users[id] = user
					filename := filepath.Join(strconv.FormatInt(time.Now().UnixNano(), 10))
					filename = filename + ".gob"
					goencode(filename, users)

					fmt.Println("[+]修改成功")
				}
			}
			}

	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

func Delete() {

	users, err:= gobdecode(getUserFile())
	if err != nil {
		fmt.Println(err)
	}

	if id, err := strconv.Atoi(InputString("请输入删除用户ID:")); err == nil {
		for _, v := range users {
			fmt.Println(id)
			if v.ID == id {
				fmt.Println("将删除的用户信息:")
				printUser(v)
				input := InputString("确定删除(Y/N)?")
				if input == "y" || input == "Y" {
					delete(users, id)

					filename := filepath.Join(strconv.FormatInt(time.Now().UnixNano(), 10))
					filename = filename + ".gob"
					goencode(filename, users)

					fmt.Println("[+]删除成功")
				}
			} else {
				fmt.Println("[-]用户ID不存在")
			}

			}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

