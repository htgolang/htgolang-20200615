package GoStudynotes

import (
	"fmt"
	"time"
)

func GetTimeFromStrDate(date string) (year int) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		fmt.Println("[err]出生日期解析错误！")
		return
	}
	year = d.Year()
	return
}

func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

