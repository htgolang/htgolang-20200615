package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

type LogInfo struct {
	ID          int
	IP          string
	Logtime     time.Time
	Method      string
	Url         string
	Status_code int
	Bytes       int
	Count       string
}

var upload_file string
var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

func CreateSqlTable() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	sql := `create table accesslog (
    id int primary key auto_increment,
    ip VARCHAR(64) not NULL DEFAULT "0.0.0.0",
    logtime  datetime not null,
    method VARCHAR(8) not null DEFAULT "GET",
    url VARCHAR(1024) not null DEFAULT "",
    status_code int not NULL DEFAULT 200,
    bytes int Not null DEFAULT 0
) engine=innodb default charset utf8mb4;`

	//fmt.Println("\n" + sql + "\n")
	smt, _ := db.Prepare(sql)

	smt.Exec()
	smt.Close()

}

func InsertSQL(ip, logtime, method, url string, status_code, bytes int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()
	_, err = db.Exec("insert into accesslog(ip, logtime, method, url, status_code, bytes) values(?,str_to_date(?, '%d/%M/%Y:%H:%i:%S'),?,?,?,?)", ip, logtime, method, url, status_code, bytes)

	if err != nil {
		panic(err)
	}
}

func Count_loginfo() []LogInfo {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select ip, method, url, status_code, count(*) from accesslog group by ip, method, url, status_code")
	if err != nil {
		panic(err)
	}

	loginfos := make([]LogInfo, 0)
	for rows.Next() {
		var loginfo LogInfo
		if err := rows.Scan(&loginfo.IP, &loginfo.Method, &loginfo.Url, &loginfo.Status_code, &loginfo.Count); err == nil {
			loginfos = append(loginfos, loginfo)
		} else {
			fmt.Println(err)
		}
	}
	return loginfos
}

func Querylog(q string) []LogInfo {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, ip, logtime, method, url, status_code, bytes from accesslog")
	if err != nil {
		panic(err)
	}

	loginfos := make([]LogInfo, 0)
	for rows.Next() {
		var loginfo LogInfo
		if err := rows.Scan(&loginfo.ID, &loginfo.IP, &loginfo.Logtime, &loginfo.Method, &loginfo.Url, &loginfo.Status_code, &loginfo.Bytes); err == nil {
			if q == "" || strings.Contains(loginfo.IP, q) ||
				strings.Contains(loginfo.Logtime.Format("2006-01-02 15:04:05"), q) || strings.Contains(loginfo.Method, q) ||
				strings.Contains(loginfo.Url, q) {
				loginfos = append(loginfos, loginfo)
			}
		} else {
			fmt.Println(err)
		}
	}
	return loginfos
}

func Readfile(upload_file string) {
	file, err := os.Open(upload_file)
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				str := strings.Split(string(line), " ")
				code, _ := strconv.Atoi(str[8])
				bytes, _ := strconv.Atoi(str[9])

				//fmt.Println(str[0], strings.Trim(str[3], "["), strings.Trim(str[5], "\""), str[6], code, bytes)
				InsertSQL(str[0], strings.Trim(str[3], "["), strings.Trim(str[5], "\""), str[6], code, bytes)
			}
		}
	}

}

func UploadAction(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("upload.html").ParseFiles("views/upload.html"))
		tpl.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseMultipartForm(1024)
		if file, header, err := request.FormFile("log"); err == nil {
			upload_file = filepath.Join("data", header.Filename)
			newFile, err := os.Create(upload_file)
			if err == nil {
				defer newFile.Close()
				io.Copy(newFile, file)
				Readfile(upload_file)
				http.Redirect(response, request, "/log/", http.StatusFound)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func CountLogAction(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("countlog.html").ParseFiles("views/countlog.html"))
		tpl.Execute(response, Count_loginfo())

	}
}

func QueryAction(response http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("query"))
	tpl := template.Must(template.New("log.html").ParseFiles("views/log.html"))
	tpl.Execute(response, struct {
		Query string
		Logs  []LogInfo
	}{q, Querylog(q)})
}

func Login(respnse http.ResponseWriter, request *http.Request) {
	http.Redirect(respnse, request, "/upload/", http.StatusFound)
}

func main() {
	CreateSqlTable()
	addr := "0.0.0.0:9999"
	http.HandleFunc("/", Login)
	http.HandleFunc("/upload/", UploadAction)
	http.HandleFunc("/countlog/", CountLogAction)
	http.HandleFunc("/log/", QueryAction)
	http.ListenAndServe(addr, nil)

}
