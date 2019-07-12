package GoStudynotes

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/howeyc/gopass"
)

func salt(pwd string) string {
	saltkey := []byte("dsadiojijml;mxklzcjioank;mdasjq")
	//salt := time.Now().Unix()
	m5 := md5.New()
	m5.Write([]byte(pwd))
	m5.Write([]byte(string(saltkey)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
func Auth(password string,maxAuth int)bool{
	stats := salt(password)
	for i:=0;i< maxAuth;i++{
		fmt.Print("请输入密码:")
		//Timeout("请输入密码:")
		if pass,err := gopass.GetPasswd();err == nil {
			if stats != salt(string(pass)) {
				fmt.Println("[err]密码错误,请重试\n")
			}else {
				return true
			}
		}
	}
	return false
}