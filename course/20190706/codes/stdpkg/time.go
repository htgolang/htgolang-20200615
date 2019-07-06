package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)

	// 2000/01/02 08:10:01
	// 2006 年
	// 01 月
	// 02 天
	// 24进制小时 15
	// 分钟 04
	// 秒 05
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	t, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 03:04:05")

	fmt.Println(err, t)

	t = time.Unix(0, 0)
	fmt.Println(t)

	d := t.Sub(now)

	fmt.Printf("%T, %v", d, d)

	// time.Second
	// time.Minute
	// time.Hour
	// 3h2m4s

	fmt.Println(time.Now())
	time.Sleep(time.Second * 5)
	fmt.Println(time.Now())
	t = now.Add(-3 * time.Hour)
	fmt.Println(t)

	d, err = time.ParseDuration("-3h2m4s")
	fmt.Println(err, d)
	fmt.Println(d.Hours())
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())

}
