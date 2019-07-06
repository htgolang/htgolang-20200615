package main

import (
	"fmt"

	lgopkg "gopkg"

	"github.com/howeyc/gopass"
	"github.com/imsilence/gopkg"
)

func main() {
	fmt.Println(gopkg.VERSION)
	fmt.Println(lgopkg.VERSION)
	// fmt.Println(lgopkg.name)
	// lgopkg.PrintName()

	fmt.Print("请输入密码：")
	if bytes, err := gopass.GetPasswd(); err == nil {
		fmt.Println(string(bytes))
	}

}
