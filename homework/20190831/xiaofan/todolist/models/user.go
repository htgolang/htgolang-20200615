package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Birthday   time.Time `json:"birthday"`
	Tel        string    `json:"tel"`
	Addr       string    `json:"addr"`
	Desc       string    `json:"desc"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
}

func GetUsers(query string) []User {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	var user User
	var newUser []User
	var sql string

	if query == "" {
		sql = "select id, name, birthday, tel, addr, `desc`, create_time from todolist_user"
		rows, err := db.Query(sql)
		if err == nil {
			for rows.Next() {
				_ = rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
				newUser = append(newUser, user)
			}
			return newUser
		}
	} else {
		sql = fmt.Sprintf("select id, name, birthday, tel, addr, `desc`, create_time from todolist_user "+
			"where name like '%%%s%%' or addr like '%%%s%%' or `desc` like '%%%s%%'", query, query, query)
		fmt.Println(sql)
		rows, err := db.Query(sql)
		if err == nil {
			for rows.Next() {
				_ = rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
				newUser = append(newUser, user)
			}
			return newUser
		}
		fmt.Println(err)
	}

	return make([]User, 0)
}

func CreateUser(name, pwd, sex, bir, tel, addr, desc string) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}

	sexI, err := strconv.Atoi(sex)
	if err != nil {
		fmt.Println(err)
	}
	sql := "insert into todolist_user(name, password, sex, birthday, tel, addr, `desc`, create_time) " +
		"values(?, md5(?), ?, ?, ?, ?, ?, now())"

	_, err = db.Exec(sql, name, pwd, sexI, birParse, tel, addr, desc)
	if err != nil {
		panic(err)
	}
}

func GetUserById(id int) (User, error) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "select id, name, birthday, tel, addr, `desc`, create_time from todolist_user where id=?"
	row := db.QueryRow(sql, id)

	var user User
	err = row.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
	if err != nil {
		fmt.Println(err)
		return User{}, errors.New("not found")
	}

	return user, nil
}

func ModifyUser(id int, name, bir, tel, addr, desc string) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "update todolist_user set name=? ,birthday=?, tel=?, addr=?, `desc`=? where id=?"
	_, err = db.Exec(sql, name, bir, tel, addr, desc, id)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(id int) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "delete from todolist_user where id=?"
	_, err = db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
}

func CheckUser(username, password string) bool {
	var user User
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return false
	}

	if err := db.Ping(); err != nil {
		return false
	}
	sql := "select name, password from todolist_user where name=? and password=md5(?)"

	row := db.QueryRow(sql, username, password)

	err = row.Scan(&user.Name, &user.Password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ModifyPassword(name, password string) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "update todolist_user set password=md5(?) where name=?"
	_, err = db.Exec(sql, password, name)
	if err != nil {
		panic(err)
	}
}

func GetUserByName(name string) (User, error) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "select id, name, birthday, tel, addr, `desc`, create_time from todolist_user where name=?"
	row := db.QueryRow(sql, name)

	var user User
	err = row.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
	if err != nil {
		fmt.Println(err)
		return User{}, errors.New("not found")
	}

	return user, nil
}

func ValidateCreateUser(name, password, bir, tel string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "用户名长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "用户名重复"
	}

	if len(password) > 30 || len(password) < 6 {
		fmt.Println(len(password))
		errors["password"] = "密码长度必须在6~30之间"
	}

	birParse, err := time.Parse("2006-01-02", bir)
	if err != nil {
		fmt.Println(err)
	}
	if birParse.Year() > 2019 || birParse.Year() < 1960 {
		errors["bir"] = "出生年必须在1960~2019之间"
	}

	if len(tel) != 11 {
		errors["tel"] = "手机号码必须为11位"
	} else if _, err := strconv.Atoi(tel); err != nil {
		errors["tel"] = "手机号码必须为数字"
	}

	return errors
}
