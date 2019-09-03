package models

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"strings"
	"time"
)

type Log struct {
	Ip     string
	Method string
	URL    string
	Code   int
	Count  int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func LogStore(log io.Reader) {
	db, err := sql.Open("mysql", dsn)
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	cmd := "insert into accesslog(ip, logtime, method, url, status_code, bytes) values(?,?,?,?,?,?)"

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

		_, err = db.Exec(cmd, n[0], logtime, strings.TrimLeft(n[5], `"`), n[6], n[8], n[9])
		if err != nil {
			panic(err)
		}
	}
}

func LogQuery() []Log {
	var log Log
	var logs []Log
	db, err := sql.Open("mysql", dsn)
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	cmd := "select ip, method, url, status_code, COUNT(status_code) from accesslog GROUP BY url"
	rows, err := db.Query(cmd)
	check(err)

	for rows.Next() {
		err := rows.Scan(&log.Ip, &log.Method, &log.URL, &log.Code, &log.Count)
		if err != nil {
			fmt.Println(err)
			continue
		}

		b := logCheck(&log)
		if b {
			continue
		}

		logs = append(logs, log)
	}
	return logs
}

func logCheck(log *Log) bool {
	if log.Ip == "" {
		return true
	}
	if len(log.Method) > 7 {
		return true
	}
	return false
}
