package main

import "fmt"

type SignalSender interface {
	Send(to, msg string) error
}

type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
	SmtpAddr string
}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给:", to, ", 消息内容是:", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type SmsSender struct {
	SmsAPI string
}

func (s *SmsSender) Send(to, msg string) error {
	fmt.Println("发送短信给:", to, ", 消息内容是:", msg)
	return nil
}

func (s *SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type WechatSender struct {
	ID string
}

func (s WechatSender) Send(to, msg string) error {
	fmt.Println("发送微信给:", to, ", 消息内容是:", msg)
	return nil
}

func (s *WechatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

//func do(sender EmailSender) {
func do(sender Sender) {
	sender.Send("领导", "工作日志")
}

func main() {
	// var sender EmailSender = EmailSender{}
	var sender Sender = EmailSender{"kk@pm.me"}
	// fmt.Println(sender.SmtpAddr)

	// fmt.Printf("%T %v\n", sender, sender)

	// sender.Send("kk", "早上好")
	// sender.SendAll([]string{"祥哥", "烟灰"}, "中午好")

	do(sender)

	// sender = SmsSender{}
	sender = &SmsSender{"juhe"}
	do(sender)

	sender = &EmailSender{"imsilece@pm.me"}
	do(sender)

	//sender = WechatSender{}
	sender = &WechatSender{"testtest"}
	do(sender)

	var ssender SignalSender = sender

	ssender.Send("小凡", "你好")

	// ssender.SendAll([]string{"小贩"}, "你好")

	sender01, ok := ssender.(Sender)

	fmt.Printf("%T, %v\n", sender01, ok)
	sender01.SendAll([]string{"小贩", "祥哥"}, "你好")

	wsender01, ok := ssender.(*WechatSender)
	fmt.Printf("%T, %v\n", wsender01, ok)
	fmt.Println(wsender01.ID)

	esender01, ok := ssender.(*EmailSender)
	fmt.Printf("%T, %v\n", esender01, ok)

	sender = EmailSender{"testtest"}

	switch v := sender.(type) {
	case EmailSender:
		fmt.Println("emailsender", v.SmtpAddr)
	case *SmsSender:
		fmt.Println("smssender", v.SmsAPI)
	case *WechatSender:
		fmt.Println("*wechatsender", v.ID)
	}
}
