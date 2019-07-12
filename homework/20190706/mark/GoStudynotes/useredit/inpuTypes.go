package GoStudynotes

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func inputName(prompt string)string{
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scan(&input)
		count := utf8.RuneCountInString(input)
		if count > 6 {
			fmt.Printf("[err]你输入的字符串是%v位，请重新输入一个长度为6的字符串!\n", count)
		} else {
			return input
		}
	}
	return input
}
func InputTel(prompt string)string{
	var a string
	for {
		fmt.Print(prompt)
		fmt.Scan(&a)
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
func InputYMD(prompt ...string)(string){
	fmt.Println(prompt)
	var year,month,day string
	for {
		fmt.Print("请输入年份:")
		fmt.Scan(&year)
		//timeNw,_ := strconv.Atoi(time.Now().Format("2006"))
		yearlen := len(year)
		if year,err := strconv.Atoi(year); err != nil {
			fmt.Printf("[err]请重新输入，并确定你输入的是一个年份!如:[%v]\n", time.Now().Year())
			continue
		} else if year >= time.Now().Year() || yearlen != 4 {
			fmt.Printf("[err]年份错误，请重新输入年份!不能大于%v\n",time.Now().Year())
			continue
		}
		for {
			fmt.Print("请输入月份:")
			fmt.Scan(&month)
			monthlen := len(month)
			if month, err := strconv.Atoi(month); err != nil {
				fmt.Print("[err]请重新输入，并确定你输入的是一个月份!如:[12]")
				continue
			} else if monthlen > 2 || month >= 13 {
				fmt.Print("[err]月份错误，请重新输入月份!不能大于12月\n")
				continue
			}
			for {
				fmt.Print("请输入日期:")
				fmt.Scan(&day)
				daylen := len(day)
				if day,err := strconv.Atoi(day);err != nil {
					fmt.Print("[err]请重新输入，并确定你输入是一个日期!如:[31]")
					continue
				}else if daylen >= 3 || day >= 32{
					fmt.Print("[err]日期错误，请重新输入日期!不能大于%v号\n")
					continue
				}
				break
			}
			break
		}
		break
	}
	Date := strings.Join([]string{year,month,day},"-")
	return Date
}
func InputDa(prompt string)string{
	var a string
	for {
		fmt.Print(prompt)
		fmt.Scan(&a)
		alen := len(a)
		if _,err := strconv.Atoi(a);err != nil {
			fmt.Println("[err]请确定你输入的是一组数字!",err)
			continue
		}else if alen >= 3 {
			fmt.Println("[err]请输入100以内的数字\n")
		}else {
			return a
		}
	}
}