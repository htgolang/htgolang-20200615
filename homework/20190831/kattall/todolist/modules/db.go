package modules



import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


const (
	dbUser     string = "root"
	dbPassword string = "123123"
	dbHost     string = "localhost"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
