package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)
const (
	MaxAuth =3
	passwordFile = ".password"
	userFile = "users.json"
)
type UserDescribe struct {
	ID int
	Name,Addr,Desc,Tel string
	Birthday time.Time
}
func (user UserDescribe) String() string{
	age := (getAge(getTimeFromStrDate(user.Birthday.Format("2006-01-02"))))
	return fmt.Sprintf("\nID:%d\n名称:%v\n年龄:%d\n联系方式:%s\n地址:%s\n出生日期:%v\n备注信息:%s\n\n",
		user.ID,user.Name,age,user.Tel,user.Addr,user.Birthday.Format("2006-01-02"),user.Desc)
}
func quits(input string){
	if input == "q" || input == "5" {
		func() {
			os.Exit(0)
		}()
	}
}
// 带缓冲的IO
func InputString1(prompt string)string{
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
func InputString(prompt string)string{
	c := make(chan string, 1)
	var input string
	go func() {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		c <- input
	}()
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 60)
	select {
	case <-ctx.Done():
		fmt.Print("\ntime out")
		os.Exit(1)
	case <-c:
		return strings.TrimSpace(input)
	}
	return strings.TrimSpace(input)
}
func getid() (id int){
	users := loadUser()
	for k,_ := range users {
		if id < k {
			id = k
		}
	}
	return id +1
}
// 用户为空和用户名称重复验证
func validUser(user UserDescribe) error {
	users := loadUser()
	if user.Name == "" {
		return fmt.Errorf("输入用户名为空")
	}
	for _,tuser := range  users{
		if user.Name == tuser.Name && user.ID != tuser.ID{
			return errors.New("键入的名称已经存在!")
		}
	}
	return nil
}
func AddUser(){
	id := getid()
	setname := inputUser(id)
	if err := validUser(setname);err == nil{
		users := loadUser()
		users[id] = setname
		stroreUser(users)
		fmt.Printf("[ok]添加%v成功\n",id)
	}else {
		fmt.Println("[err]添加失败!")
	}
}
func inputUser(id int) UserDescribe{
	user := UserDescribe{}
	user.ID = id
	fmt.Println("[info]立即退出请输入5或者q")
	user.Name = inputName("请输入名字:",id)
	birthday,_ := time.Parse("2006-01-02",inputYMD("请按照提示输入出生年月日!"))
	user.Birthday = birthday
	user.Tel = inputTel("请输入联系方式:")
	user.Addr = InputString("请输入联系地址:")
	user.Desc  = InputString("请输入备注:")
	return user
}
func inputName(prompt string,id int)string{
	users := loadUser()
	var input string
	for {
	EOF:
		fmt.Print(prompt)
		fmt.Scan(&input)
		quits(input)
		for _,checkuser := range users {
			if input == checkuser.Name && checkuser.ID != id {
				fmt.Printf("[err]%v用户名存在,请重新输入!\n", input)
				goto EOF
			}
		}
		count := utf8.RuneCountInString(input)
		if count > 6 {
			fmt.Printf("[err]你输入的字符串是%v位，请重新输入一个长度为6的字符串!\n", count)
		} else {
			return input
		}
	}
	return input
}
func ModifyUser(){
	if id,err := strconv.Atoi(InputString("[info]请输入你即将修改的用户id，输入5退出:"));err == nil {
		users := loadUser()
		// 判断id在不在
		quits(string(id))
		if user,ok :=  users[id];ok{
			fmt.Printf("[warning]即将修改的用户信息:%v",user)
			input := InputString("\n你确定修改吗？（Y/y）")
			quits(input)
			if input == "y"|| input == "Y" {
				user := inputUser(id)
				if err := validUser(user);err == nil{
					users[id] = user
					stroreUser(users)
					fmt.Printf("[ok]修改%v成功\n",id)
				}else {
					fmt.Println("[err]修改失败!")
					fmt.Println(err)
				}
			}
		}else {
			fmt.Println("[err]用户id不存在\n")
		}
	}
}
func DeleteUser(){
	users := loadUser()
	if id,err := strconv.Atoi(InputString("请输入你即将删除的用户id，输入5退出:"));err == nil {
		// 判断id在不在
		if user,ok :=  users[id];ok{
			fmt.Printf("[warning]即将删除的用户信息:%v",user)
			input := InputString("\n你确定删除吗？（Y/y）")
			quits(input)
			if input == "y"|| input == "Y" {
				delete(users,id)
				stroreUser(users)
				fmt.Printf("[ok]删除%v成功\n",id)
			}
		}else {
			fmt.Println("[err]用户id不存在\n")
		}
	}
}
func QueryUser() {
	users := loadUser()
	q := InputString("请输入查询内容:")
	quits(q)
	list := make([]UserDescribe,0)
	for _, v := range users {
		if strings.Contains(v.Name, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Tel, q)  ||  strings.Contains(v.Desc, q) {
			list = append(list,v)
		}
	}
	if len(list) == 0{
		fmt.Println("查询内容为空!")
	}else {
		in := InputString("\n请输入排序属性:[id/name/]，立即退出输入5:")
		quits(in)
		sort.Slice(list, func(i, j int) bool {
			switch in {
			case "id":
				return list[i].ID < list[j].ID
			case "name":
				return list[i].Name < list[j].Name
			default:
				return false
			}
		})
		for _,user := range list{
			fmt.Print(user)
			fmt.Println("```````````````````````")
		}
	}
}
func loadUser() map[int]UserDescribe{
	users := map[int]UserDescribe{}
	if file,err := os.Open(userFile);err == nil{
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&users)
	}else {
		if !os.IsNotExist(err){
			fmt.Println("[err]loadUser 文件打开失败",err)
		}
	}
	return users
}
func fileBackup(){
	// 进行COPY文件
	srcfile,err := os.Open(userFile)
	if err != nil{
		fmt.Println("[info]首次运行将会创建文件!")
	}else {
		defer srcfile.Close()
		dsfName := fmt.Sprint(time.Now().Format("2006-01-02_15-04-05")+".user.json") // 以时间为格式
		destfile, err:= os.Create(dsfName)
		if err != nil{
			fmt.Println(err)
		}
		reader := bufio.NewReader(srcfile)
		writer := bufio.NewWriter(destfile)
		bytes := make([]byte, 1024*1024*1)
		for {
			n, err := reader.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			writer.Write(bytes[:n])
			writer.Flush()
		}
	}
	names,err := filepath.Glob("*.user.json")
	if err != nil{
		fmt.Println("没有备份文件")
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(names))) // 对过滤结果进行排序
		if len(names) >= 4 {
			for _, name := range names[3:] { // 保留三个文件
				os.Remove(name)
			}
		}
	}
}
func stroreUser(users map[int]UserDescribe){
	// 拷贝函数
	fileBackup()
	// 开始逻辑
	if file,err := os.Create(userFile); err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(users)
	}else {
		fmt.Println(err)
	}
}
func Auth()bool{
	passwd,err := ioutil.ReadFile(passwordFile)
	if err == nil && len(passwd) > 0 {  // 密码长度大于0进行验证
		// 验证密码
		for i:=0;i< MaxAuth;i++{
			fmt.Print("请输入密码:")
			if pass,err := gopass.GetPasswd();err == nil {
				if string(passwd) != fmt.Sprintf("%x",md5.Sum(pass)) {
					fmt.Println("[err]密码错误,请重试\n")
				}else {
					return true
				}
			}
		}
		return false
	}else {
		if os.IsNotExist(err){
			fmt.Print("请输入初始化密码:")
			bytes,_ := gopass.GetPasswd()
			ioutil.WriteFile(passwordFile,[]byte(fmt.Sprintf("%x",md5.Sum(bytes))),os.ModePerm)
			return true
		}else {
			fmt.Println(err)
			return false
		}
	}
}
func ModifyPass(){
	passwd,err := ioutil.ReadFile(passwordFile)
	if err == nil && len(passwd) > 0 { // 密码长度大于0进行验证
		// 验证密码
		fmt.Print("请输入密码:")
		if pass, err := gopass.GetPasswd(); err == nil {
			if string(passwd) == fmt.Sprintf("%x", md5.Sum(pass)) {
				fmt.Print("请输入新密码:")
				bytes,_ := gopass.GetPasswd()
				ioutil.WriteFile(passwordFile,[]byte(fmt.Sprintf("%x",md5.Sum(bytes))),os.ModePerm)
				fmt.Println("[ok]密码修改完成!")
			} else {
				fmt.Println("[err]密码错误!")
			}
		}
	}
}
func inputTel(prompt string)string{
	var a string
	for {
		fmt.Print(prompt)
		fmt.Scan(&a)
		quits(a)
		count := utf8.RuneCountInString(a)
		_,err := strconv.ParseInt(a,11,0)
		if err != nil {
			fmt.Println("[err]请确定你输入的是11位手机号码或者7位数的电话号码!",err)
			continue
		}else if int(count) == 7 || int(count) == 11 {
			return a
		}else {
			fmt.Printf("[err]你输入了%v位，请确定你输入的是11位手机号码或者7位数的电话号码!\n", count)
		}
	}
	return a
}
func inputYMD(prompt ...string)(string) {
	fmt.Println(prompt)
	var year, month, day string
	for {
		fmt.Print("请输入年份:")
		fmt.Scan(&year)
		yearlen := len(year)
		if year, err := strconv.Atoi(year); err != nil {
			fmt.Printf("[err]请重新输入，并确定你输入的是一个年份!如:[%v]\n", time.Now().Year())
			continue
		} else if year >= time.Now().Year() || yearlen != 4 {
			fmt.Printf("[err]年份错误，请重新输入年份且不能大于%v\n", time.Now().Year())
			continue
		}
		for {
			fmt.Print("请输入月份:")
			fmt.Scan(&month)
			if month, err := strconv.Atoi(month); err != nil {
				fmt.Print("[err]请重新输入，并确定你输入的是一个月份!如:[12]\n")
				continue
			} else if month >= 13 {
				fmt.Print("[err]月份错误，请重新输入月份!不能大于12月\n")
				continue
			}
			for {
				fmt.Print("请输入日期:")
				fmt.Scan(&day)
				if day, err := strconv.Atoi(day); err != nil {
					fmt.Print("[err]请重新输入，并确定你输入是一个日期!如:[31]\n")
					continue
				} else if day >= 32 {
					fmt.Print("[err]日期错误，请重新输入日期!不能大于31号\n")
					continue
				}
				break
			}
			break
		}
		monthlen := len(month)
		daylen := len(day)
		if monthlen == 1 {
			month = fmt.Sprintf("0%v", month)
		}
		if daylen == 1 {
			day = fmt.Sprintf("0%v", day)
		}
		Date := strings.Join([]string{year, month, day}, "-")
		return Date
	}
}
func getTimeFromStrDate(date string) (year int) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		fmt.Println("[err]出生日期解析错误！")
		return
	}
	year = d.Year()
	return
}

func getAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

func main()  {
	var noauth bool
	flag.BoolVar(&noauth,"N",false,"no auth")
	flag.Parse()
	if !noauth && !Auth(){
		fmt.Printf("密码输入%d次错误，程序退出",MaxAuth)
		return
	}
	menu:= 	`===========================================
1.查询
2.添加
3.修改
4.删除
5.退出
0.修改密码
===========================================
`
	callbacks := map[string]func(){
		"1": QueryUser,
		"2": AddUser,
		"3": ModifyUser,
		"4": DeleteUser,
		"5": func() {
			os.Exit(0)
		},
		"0": ModifyPass,
	}
	fmt.Println("欢迎登陆用户管理系统")
	for {
		fmt.Println(menu)
		if callbacks,ok := callbacks[InputString1("请输入你的指令:")];ok {
			callbacks()
		}else {
			fmt.Println("指令错误")
		}
	}
}