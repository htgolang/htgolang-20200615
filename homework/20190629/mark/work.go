package main

import (
	"fmt"
	"strconv"
	"strings"
)
var (
	id   string
	name string
	age  string
	tel  string
	addr string
	excuse string
)
/*
func auth()(b bool){
	password := "123abc!@#"
	var pass string
	num := 3
	for i := 1; i <= num; i++ {
		rnum := num -i
		fmt.Print("请输入密码:")
		fmt.Scan(&pass)
		if pass == password {
			b = true
		}else {
			fmt.Printf("秘密错误,你还可以尝试%v次\n",rnum)
			continue
		}
		return
	}
	return
}
*/
func auth()(err error){
	defer func() {
		if e :=recover();e != nil{
			err = fmt.Errorf("%v",e)  //  这样以来，如果有错误返回错误信息，如果没有则返回nil
		}
	}()
	password := "123abc!@#"
	var pass string
	num := 3
	for i := 1; i <= num; i++ {
		rnum := num -i
		fmt.Print("请输入密码:")
		fmt.Scan(&pass)
		if pass == password {
			panic("isok")
		}else {
			fmt.Printf("密码错误,你还可以尝试%v次\n",rnum)
			continue
		}
	}
	return
}
func printab(){
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s","id","Name","Age","Tel","Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-",len(title)))
}
func printab2(user map[string]string){
	fmt.Printf("%5s|%20s|%5s|%20s|%50s",user["id"],user["name"],user["age"],user["tel"],user["addr"])
	fmt.Println()
}
func typetext(){
	fmt.Print("Please enter a name:")
	fmt.Scan(&name)
	if len(name) >= 10 {
		fmt.Println("warning:名称过长")
	}
	fmt.Print("Please enter a age:")
	fmt.Scan(&age)
	if len(age) >= 3 {
		fmt.Println("warning:超出范围")
	}
	fmt.Print("Please enter a tel:")
	fmt.Scan(&tel)
	if len(tel) == 7 || len(tel) == 11 {
		fmt.Println("warning:号码非法")
	}
	fmt.Print("Please enter a addr:")
	fmt.Scan(&addr)
	if len(addr) >= 10 {
		fmt.Println("warning:名称过长")
	}
}
func add(pk int,users map[string]map[string]string){ // 添加用户
	var (
		id string = strconv.Itoa(pk)	// 转换字符串
	)
	typetext()
	users[id]= map[string]string{
		"id": id,
		"name":name,
		"tel":tel,
		"age":age,
		"addr":addr,
	}
	fmt.Println(id,name,age,tel,addr)

}
func query(users map[string]map[string]string){
	var q string
	fmt.Print("输入查询信息:")
	fmt.Scan(&q)
	printab()
	for _,user := range users {
		if q == ""|| strings.Contains(user["name"],q)|| strings.Contains(user["tel"],q) || strings.Contains(user["addr"],q){	  // 判断q是否为空，是否存在
		printab2(user)
	}
	}
}
func updateus(users map[string]map[string]string){
	fmt.Print("请输入要修改的用户ID：")
	fmt.Scan(&id)
	if user,ok := users[id];ok {
		printab()
		printab2(user)
		fmt.Print("是否确认修改(Y/N):")
		fmt.Scan(&excuse)
		if excuse == "Y"|| excuse == "y"{
			typetext()
			user["name"], user["age"], user["tel"], user["addr"] = name, age, tel, addr
			fmt.Printf("id%v修改成功\n",id)
			printab()
			printab2(user)
		}

	}else {
		fmt.Printf("键入的id:%v 不存在，请重试",id)
	}
}
func deleteus(users map[string]map[string]string){
	fmt.Print("请输入要删除的用户ID：")
	fmt.Scan(&id)
	if user,ok := users[id];ok {
		printab()
		printab2(user)
		fmt.Print("确定要删除吗(y/Y)：")
		fmt.Scan(&excuse)
		if excuse == "Y"|| excuse == "y"{
			delete(users, id)
			fmt.Printf("id %v已经被成功删除\n", id)
			printab()
			printab2(user)
		}
	}else {
		fmt.Printf("用户id:%v 不存在，删除用户失败\n",id)
	}
}
func main(){
	authok := auth()
	//if authok {
	if authok != nil{


		users:= make(map[string]map[string]string)
		id := 0
		for{
			var op string
			fmt.Print(`
用户管理
		1.create user
		2.update user
		3.delete user
		4.select user
		5.quit
pls input num:`)

			fmt.Scan(&op)

			inpitNm := op
			if inpitNm == "1" {
				id ++
				add(id,users)
			}else 	if inpitNm == "2" {
				updateus(users)
			}else 	if inpitNm == "3" {
				deleteus(users)
			}else 	if inpitNm == "4" {
				query(users)
			}else 	if inpitNm == "5"|| inpitNm == "q"||inpitNm == "exit" ||inpitNm == "quit"{
				fmt.Println("bay bay")
				break
			}else {
				fmt.Println("error")
			}
		}
	}

}