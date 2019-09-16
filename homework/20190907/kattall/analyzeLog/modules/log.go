package modules

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	gorm.Model
	Ip string	`gorm: "type: varchar(32)"; not null; default: ''`
	AccessTime *time.Time `gorm:"type: datetime"; not null; default: ''`
	Method string	`gorm: "type: varchar(12)"; not null; default: ''`
	Url string	`gorm: "type: varchar(2048); not null; default: ''"`
	Protocol string	`gorm: "type: varchar(32); not null; default: ''"`
	StatusCode int	`gorm: "not null; default: 404"`
	Length int	`gorm: "not null; default 0"`
}



func (l Log) TableName() string {
	return "todolist_log"
}

func StoreLog(logfile io.Reader) error {
	var log Log
	reader := bufio.NewReader(logfile)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		access_time, _ := time.Parse("02/Jan/2006:15:04:05 -0700", strings.Split(strings.Split(string(line), "]")[0], "[")[1])
		code, _ := strconv.Atoi(strings.Fields(string(line))[8])
		length, _ := strconv.Atoi(strings.Fields(string(line))[9])
		fmt.Println(log.Ip, log.AccessTime, log.Method, log.Url, log.Protocol, log.StatusCode, log.Length)
		db.Create(&Log{
			Ip:         strings.Fields(string(line))[0],
			AccessTime: &access_time,
			Method:     strings.TrimLeft(strings.Fields(string(line))[5], "\""),
			Url:        strings.Fields(string(line))[6],
			Protocol:   strings.TrimRight(strings.Fields(string(line))[7], "\""),
			StatusCode: code,
			Length:     length,
		})
	}
	return nil
}

// 返回最大页码
func GetPage(pagecount int) int {
	maxPage := 0
	var log Log
	db.Model(&log).Limit(pagecount).Count(&maxPage)
	if maxPage / pagecount == 0 {
		return maxPage / pagecount
	} else {
		return (maxPage / pagecount) + 1
	}
}

// 通过每页显示书，当前页数，获取log
func GetLog(pageCount int, currentPage int) []Log {
	var logs []Log
	offset := (currentPage - 1) * pageCount
	fmt.Println("pagecount, currentpage, offerset:", pageCount, currentPage, offset)
	db.Select("id, ip, access_time, method, url, protocol, status_code, `length`").Limit(pageCount).Offset(offset).Find(&logs)
	return logs
}

func GetIPCountLog(pageCount int, currentPage int) map[string]int {
	var log Log
	offset := (currentPage - 1) * pageCount
	fmt.Println("pagecount, currentpage, offerset:", pageCount, currentPage, offset)

	ipc := make(map[string]int, 0)
	rows, _ := db.Model(&log).Limit(pageCount).Offset(offset).Select("ip, count(*) as cnt").Order("cnt desc").Group("ip").Rows()
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		ipc[name] = count
	}
	fmt.Println(ipc)
	return ipc
}

func GetStatusCountLog(pageCount int, currentPage int) map[string]int {
	var log Log
	offset := (currentPage - 1) * pageCount
	fmt.Println("pagecount, currentpage, offerset:", pageCount, currentPage, offset)

	ipc := make(map[string]int, 0)
	rows, _ := db.Model(&log).Limit(pageCount).Offset(offset).Select("status_code, count(*) as cnt").Order("cnt desc").Group("status_code").Rows()
	for rows.Next() {
		var statusCode string
		var count int
		rows.Scan(&statusCode, &count)
		ipc[statusCode] = count
	}
	fmt.Println(ipc)
	return ipc
}

func IpUrlCount(pageCount int, currentPage int) map[string]map[string]int{
	var log Log
	offset := (currentPage - 1) * pageCount
	ipc := make(map[string]map[string]int, 0)
	ipcount := make(map[string]int, 0)
	fmt.Println("pagecount, currentpage, offerset:", pageCount, currentPage, offset)
	rows, _ := db.Model(&log).Limit(pageCount).Offset(offset).Select("ip, url , count(*) as cnt").Order("cnt desc").Group("ip, url").Rows()
	for rows.Next() {
		var ip string
		var url string
		var count int
		rows.Scan(&ip, &url, &count)
		ipcount[url] = count
		ipc[ip] = ipcount
	}
	for ip, value := range ipc {
		for url, count := range value {
			fmt.Println("ip, url, count:", ip, url, count)
		}
	}
	return ipc
}