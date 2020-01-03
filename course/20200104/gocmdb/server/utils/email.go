package utils

import (
	"gopkg.in/gomail.v2"
)

type Email struct {
	host string
	port int
	user string
	password string
}


func NewEmail(host string, port int, user string, password string)  *Email {
	return &Email{
		host: host,
		port: port,
		user: user,
		password: password,
	}
}

func (e *Email)Send(to []string, subject string, msg string, attaches[]string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.user)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	for _, attach := range attaches {
		m.Attach(attach)
	}

	d := gomail.NewDialer(e.host, e.port, e.user, e.password)

	return d.DialAndSend(m)
}