package auth

import (
	"crypto/md5"
	"fmt"

	"github.com/howeyc/gopass"
)

// 定义常量maxAuth，为输入密码的最多次数
// 定义常量password，为密码的MD5加密值
const (
	maxAuth  = 3
	password = "0a60584d6050943fc98cce83f2052739"
)

// 认证函数
func Auth() bool {
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入JevonWei用户系统密码: ")
		// 将输入的密码隐形显示
		bytes, _ := gopass.GetPasswd()

		// 如果输入密码的MD5值等于password，返回True
		if password == fmt.Sprintf("%x", md5.Sum(bytes)) {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	fmt.Printf("密码输入%d次错误，程序退出\n", maxAuth)
	return false
}
