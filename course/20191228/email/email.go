package main

import (
	"gopkg.in/gomail.v2"
)

func main() {

	server := "smtp.qq.com"
	port := 465

	user := "782874382@qq.com"
	password := "bsmujneemttzbbfe"

	to := []string{"786725806@qq.com", "370143965@qq.com", "imsilence@outlook.com", "imsilence@pm.me"}

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "CMDB告警测试邮件")
	m.SetBody("text/html", "<div style='color:red;'>CMDB告警测试邮件</div>CMDB告警测试邮件<img src='cid:test.png' height=200 width=300 />")
	m.Attach("./email.go")
	m.Attach("./test.png")

	d := gomail.NewDialer(server, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}