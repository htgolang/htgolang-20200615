package models

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	gorm.Model
	Ip         string    `gorm:"type: varchar(32); not null; default: ''"`
	AccessTime time.Time `gorm:"type: datetime; "`
	Method     string    `gorm:"type: varchar(12); not null; default: ''"`
	Url        string    `gorm:"type: varchar(2048); not null; default: ''"`
	Protocol   string    `gorm:"type: varchar(32); not null; default: ''"`
	StatusCode int       `gorm:"not null; default: 404"`
	Length     int       `gorm:"not null; default 0"`
}

func (l Log) TableName() string {
	return "access_log"
}

func LogStore(log io.Reader) {
	reader := bufio.NewReader(log)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		n := strings.SplitN(string(line), " ", -1)

		timeLeft := strings.TrimLeft(n[3], `[`)
		timeRight := strings.TrimRight(n[4], `]`)
		logtime, err := time.Parse("02/Jan/2006:15:04:05 -0700", timeLeft+" "+timeRight)
		if err != nil {
			fmt.Println(err)
		}
		code, _ := strconv.Atoi(n[8])
		length, _ := strconv.Atoi(n[9])

		//fmt.Println(n[0], logtime, strings.TrimLeft(n[5], `"`), n[6], strings.TrimRight(n[7], `"`), n[8], n[9])
		db.Create(&Log{
			Ip:         n[0],
			AccessTime: logtime,
			Method:     strings.TrimLeft(n[5], `"`),
			Url:        n[6],
			Protocol:   strings.TrimRight(n[7], `"`),
			StatusCode: code,
			Length:     length,
		})
	}
}

func LogQuery() []map[string]string {
	var log Log

	var logs []map[string]string
	rows, err := db.Model(&log).Select("ip, url, count(*)").Group("ip, url").Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		tmp := make(map[string]string)
		var ip string
		var url string
		var count int
		err := rows.Scan(&ip, &url, &count)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tmp["IP"] = ip
		tmp["URL"] = url
		tmp["Count"] = strconv.Itoa(count)
		logs = append(logs, tmp)
	}
	return logs
}
