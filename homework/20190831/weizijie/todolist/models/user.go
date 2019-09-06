package models

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	ID         int       `json:id`
	Name       string    `json:name`
	Birthday   time.Time `json:birthday`
	Sex        bool
	Addr       string `json:addr`
	Tel        string `json:tel`
	Desc       string `json:desc`
	Password   string `json:password`
	CreateTime time.Time
}

type User_str struct {
	ID         int    `json:id`
	Name       string `json:name`
	Birthday   string `json:birthday`
	Sex        bool
	Addr       string `json:addr`
	Tel        string `json:tel`
	Desc       string `json:desc`
	Password   string `json:password`
	CreateTime time.Time
}

var Register bool
var Login_User string

func (u User) ValidatePassword(passwd string) bool {
	log.Printf("%s Account Verify Success", u.Name)
	return passwd == u.Password
}

func Init() {
	logfile := "user.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
	}

	//CreateTable()
}

// func CreateTable() {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	sql := `create table users1 (
// 	id int primary key auto_increment,
// 	name varchar(64) unique not null default "",
// 	password varchar(1024) not null default "",
// 	sex boolean not null default false,
// 	birthday date,
// 	tel varchar(32) not null default "",
// 	addr varchar(128) not null default "",
// 	desc text,
// 	create_time datetime not null
// ) engine=innodb default charset utf8mb4;`

// 	fmt.Println("\n" + sql + "\n")
// 	smt, _ := db.Prepare(sql)

// 	smt.Exec()
// 	smt.Close()
// }

func GetUsers(q string) []User_str {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name, sex, birthday, addr, tel, `desc` from users")
	if err != nil {
		panic(err)
	}

	users := make([]User_str, 0)
	for rows.Next() {
		var user User
		var user_str User_str
		if err := rows.Scan(&user.ID, &user.Name, &user.Sex, &user.Birthday, &user.Addr, &user.Tel, &user.Desc); err == nil {
			if q == "" || strings.Contains(user.Name, q) ||
				strings.Contains(user.Tel, q) || strings.Contains(user.Addr, q) ||
				strings.Contains(user.Desc, q) {
				user_str.ID = user.ID
				user_str.Name = user.Name
				user_str.Sex = user.Sex
				user_str.Birthday = user.Birthday.Format("2006-01-02")
				user_str.Addr = user.Addr
				user_str.Tel = user.Tel
				user_str.Desc = user.Desc
				users = append(users, user_str)
			}
		}
	}
	return users
}

func GetUserByName(name string) (User, error) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user, err
	}

	if err := db.Ping(); err != nil {
		return user, err
	}

	defer db.Close()
	sql := "select id,name,password,sex,birthday,tel,addr,create_time from users where name=?"
	row := db.QueryRow(sql, name)
	err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Sex, &user.Birthday, &user.Tel, &user.Addr, &user.CreateTime)
	return user, err

}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "Name长度需在4-12位之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "Name已存在"
	}

	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "password长度需在6-30位之间"
	}

	return errors
}

func CreateUser(name, password, birthday, tel, addr, desc string, sex bool) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into users(name, password, sex, birthday, tel, addr, `desc`, create_time) values(?,?,?,?,?,?,?,?)", name, password, sex, birthday, tel, addr, desc, time.Now())

	if err != nil {
		panic(err)
	}
}

func GetUserById(id int) (User_str, error) {
	var user User
	var user_str User_str

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user_str, nil
	}

	if err := db.Ping(); err != nil {
		return user_str, nil
	}

	defer db.Close()

	row := db.QueryRow("select id, name, sex, birthday, tel, addr, `desc` from users where id=?", id)
	err = row.Scan(&user.ID, &user.Name, &user.Sex, &user.Birthday, &user.Tel, &user.Addr, &user.Desc)

	user_str.ID = user.ID
	user_str.Name = user.Name
	user_str.Sex = user.Sex
	user_str.Birthday = user.Birthday.Format("2006-01-02")
	user_str.Addr = user.Addr
	user_str.Tel = user.Tel
	user_str.Desc = user.Desc

	return user_str, err

}

func ModifyUser(id int, name, birthday, addr, tel, desc string, sex bool) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("update users set name=?, birthday=?, addr=?, tel=?, `desc`=?, sex=? where id = ?", name, birthday, addr, tel, desc, sex, id)

	if err != nil {
		panic(err)
	}

	log.Printf("%s user Create Success", name)

}

func DeleteUser(id int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("delete from users where id = ?", id)

	if err != nil {
		panic(err)
	}

	log.Printf("ID=%d user Delete Success", id)

}

func ModifyPasswd(name, passwd string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("update users set password=? where name = ?", passwd, name)

	if err != nil {
		panic(err)
	}

	log.Printf("Name为 %s 的Passwd Modidy Success", name)

}
