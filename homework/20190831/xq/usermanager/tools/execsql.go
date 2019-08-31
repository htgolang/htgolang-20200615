package tools

import (
	"database/sql"
	"fmt"
	"os"
)

func main(){

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser,dbPass,dbHost,dbPort,dbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	if err := db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}


	// 预处理
	//sql := fmt.Sprintf("select name, password from users where name=? and password=md5(?)")

	//sql := "insert into users (name, password, sex, brithday, tel,addr, descs, create_time) " +
	//	//	"values (?, md5(?), ?, ?, ?, ?, ?, ?)"

	//sql := "delete from users where id=?"

	sql := "update users set name=?, password=md5(?) where id=?"

	// //插入
	//result, err := db.Exec(sql, "test06", "test06", 2, "1999-01-01", "a", "b", "c", time.Now())

	// //删除
	//result, err := db.Exec(sql, 17)

	// //更新
	result, err := db.Exec(sql, "test006", "test006", 18)

	if err != nil {
		panic(err)
	}

	result.LastInsertId()
	result.RowsAffected()



	db.Close()


}
