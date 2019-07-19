package users

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var users map[int]UserDescribe = map[int]UserDescribe{}
type UserDescribe struct {
	id int
	name,addr,desc,tel,birthday string
}
func newInput(name,addr,desc,tel,birthday string,id int)UserDescribe{
	return UserDescribe{
		id:id,
		name:name,
		addr:addr,
		desc:desc,
		tel:tel,
		birthday:birthday,
	}
}
func quits(input string){
	if input == "q" || input == string(5) {
		func() {
			os.Exit(0)
		}()
	}
}
func getid() int{
	var id int
	for k,_ := range users {
		if id < k {
			id = k
		}
	}
	return id +1
}
func inputName(prompt string)string{
	var input string
	for {
	EOF:
		fmt.Print(prompt)
		fmt.Scan(&input)
		quits(input)
		for _,checkuser := range users {
			if input == checkuser.name {
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
func inputUser() UserDescribe{
	user := UserDescribe{}
	fmt.Println("[info]立即退出请输入5或者q")
	user.name = inputName("请输入名字:")
	user.birthday = inputYMD("请按照提示输入出生年月日!")
	user.tel = inputTel("请输入联系方式:")
	user.addr = InputString("请输入联系地址:")
	user.desc  = InputString("请输入备注:")
	return user
}
func printUser(user UserDescribe){
	fmt.Println("ID:",user.id)
	fmt.Println("名称:",user.name)
	fmt.Println("年龄:",(getAge(getTimeFromStrDate(string(user.birthday)))),string('岁'))
	fmt.Println("联系方式:",user.tel)
	fmt.Println("地址:",user.addr)
	fmt.Println("出生日期:",user.birthday)
	fmt.Println("备注信息:",user.desc)
}