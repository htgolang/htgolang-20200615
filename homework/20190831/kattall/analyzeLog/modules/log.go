package modules

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	Id int
	Ip string
	AccessTime time.Time
	Method string
	Url string
	Protocol string
	StatusCode int
	Length int
	Count int
}

func StoreLog(logfile io.Reader) error {
	var log Log
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}


	reader := bufio.NewReader(logfile)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		//fmt.Println(string(line))
		log.Ip = strings.Fields(string(line))[0]
		access_time, err := time.Parse("02/Jan/2006:15:04:05 -0700", strings.Split(strings.Split(string(line), "]")[0], "[")[1])
		if err != nil {
			fmt.Printf("转换 %s, 发生错误: %s\n", line, err)
		}
		log.AccessTime = access_time
		log.Method = strings.TrimLeft(strings.Fields(string(line))[5], "\"")
		log.Url = strings.Fields(string(line))[6]
		log.Protocol = strings.TrimRight(strings.Fields(string(line))[7], "\"")
		code, err := strconv.Atoi(strings.Fields(string(line))[8])
		if err != nil {
			fmt.Printf("转换 %s, 发生错误: %s\n", line, err)
		}
		log.StatusCode = code

		fmt.Println(log.Ip, log.AccessTime, log.Method, log.Url, log.Protocol, log.StatusCode, log.StatusCode)
		_, err = db.Exec("insert into accesslog(ip, accesstime, method, url, protocol, statuscode, `length`) values(?,?,?,?,?,?,?)", log.Ip, log.AccessTime, log.Method, log.Url, log.Protocol, log.StatusCode, log.Length)
		if err != nil {
			fmt.Printf("导入%s, 发生错误: %s\n", line, err)
		}
	}
	db.Close()
	return nil
}

// 返回最大页码
func GetPage(pagecount int) int {
	maxPage := 0
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows := db.QueryRow("select count(*) from accesslog limit ?", pagecount)
	rows.Scan(&maxPage)
	if maxPage / pagecount == 0 {
		return maxPage / pagecount
	} else {
		return (maxPage / pagecount) + 1
	}
}

// 通过每页显示书，当前页数，获取log
func GetLog(pageCount int, currentPage int) []Log {
	var log Log
	var logs []Log
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	offset := (currentPage - 1) * pageCount
	fmt.Println("pagecount, currentpage, offerset:", pageCount, currentPage, offset)

	rows, _ := db.Query("select id, ip, accesstime, method, url, protocol, statuscode, `length` from accesslog limit ? offset ?", pageCount, offset)
	for rows.Next() {
		rows.Scan(&log.Id, &log.Ip, &log.AccessTime, &log.Method, &log.Url, &log.Protocol, &log.StatusCode, &log.Length)
		fmt.Println(log.Ip, log.AccessTime, log.Method, log.Url, log.Protocol, log.StatusCode, log.Length)
		logs = append(logs, log)
	}
	return logs
}

func GetStatusCountLog() []Log{
	var log Log
	var logs []Log
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select statuscode , count(*) from accesslog order by statuscode ")
	for rows.Next() {
		rows.Scan(&log.StatusCode, &log.Count)
		fmt.Println(log.StatusCode, log.Count)
		logs = append(logs, log)
	}
	return logs
}

func GetIPCountLog() []Log {
	var log Log
	var logs []Log
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select ip, count(*) from accesslog order by ip ")
	for rows.Next() {
		rows.Scan(&log.Ip, &log.Count)
		fmt.Println(log.Ip, log.Count)
		logs = append(logs, log)
	}
	return logs
}

func IpUrlCount() []Log {
	var log Log
	var logs []Log
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select ip, url , count(*) from accesslog order by ip, url ")
	for rows.Next() {
		rows.Scan(&log.Ip, &log.Url, &log.Count)
		fmt.Println(log.Ip, log.Url, log.Count)
		logs = append(logs, log)
	}
	return logs
}