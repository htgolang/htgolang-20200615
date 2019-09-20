package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

type Log struct {
	gorm.Model
	IP          string    `gorm:"type:varchar(64); not null; default:'0.0.0.0' "`
	Logtime     time.Time `gorm:"type:datetime; not null"`
	Method      string    `gorm:"type:varchar(8); not null; default:'GET' "`
	Url         string    `gorm:"type:varchar(1024); not null; default:'' "`
	Status_code int       `gorm:"type:int; not null; default:200 "`
	Bytes       int       `gorm:"type:int; not null; default:0 "`
	Count       string    `gorm:"type:varchar(8); not null; default:'' "`
}

var upload_file string
var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
var db *gorm.DB

func InsertSQL(ip, logtime, method, url string, status_code, bytes int) {
	_ = db.Exec("insert into logs(created_at, updated_at, ip, logtime, method, url, status_code, bytes) values(now(), now(), ?,str_to_date(?, '%d/%M/%Y:%H:%i:%S'),?,?,?,?)", ip, logtime, method, url, status_code, bytes)
}

func Count_loginfo() []Log {
	rows, err := db.Model(&Log{}).Select("ip, method, url, status_code, count(*)").Group("ip, method, url, status_code").Rows()
	if err != nil {
		panic(err)
	}

	logs := make([]Log, 0)
	for rows.Next() {
		var log Log
		if err := rows.Scan(&log.IP, &log.Method, &log.Url, &log.Status_code, &log.Count); err == nil {
			logs = append(logs, log)
		} else {
			fmt.Println(err)
		}
	}
	return logs
}

func Querylog(q string) []Log {
	var logs []Log
	if q == "" {
		db.Find(&logs)
	} else {
		db.Where("ip like ?", "%"+q+"%").Or("method like ?", "%"+q+"%").Or("logtime like ?", "%"+q+"%").Or("url like ?", "%"+q+"%").Or("Status_code like ?", "%"+q+"%").Find(&logs)
	}

	return logs
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
		Logs  []Log
	}{q, Querylog(q)})
}

func Login(respnse http.ResponseWriter, request *http.Request) {
	http.Redirect(respnse, request, "/upload/", http.StatusFound)
}

func main() {
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil || db.DB().Ping() != nil {
		panic("Error Connect DB")
	}

	if !db.HasTable(&Log{}) {
		db.CreateTable(&Log{})
	}

	addr := "0.0.0.0:9999"
	http.HandleFunc("/", Login)
	http.HandleFunc("/upload/", UploadAction)
	http.HandleFunc("/countlog/", CountLogAction)
	http.HandleFunc("/log/", QueryAction)
	http.ListenAndServe(addr, nil)

}
