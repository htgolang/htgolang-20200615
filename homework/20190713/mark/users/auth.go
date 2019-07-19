package users

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
)
const (
	MaxAuth =3
	password = "34f85ca80ec353d3052b8a2d3973a0c5"		// abc123
)
func Auth()bool{
	for i:=0;i< MaxAuth;i++{
		fmt.Print("请输入密码:")
		if pass,err := gopass.GetPasswd();err == nil {
			if password != fmt.Sprintf("%x",md5.Sum(pass)) {
				fmt.Println("[err]密码错误,请重试\n")
			}else {
				return true
			}
		}
	}
	return false
}