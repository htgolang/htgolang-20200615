package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"user/utils"

	_ "github.com/go-sql-driver/mysql"
)

type UserDescribe struct {
	Id         int
	Name       string
	Addr       string
	Desc       string
	Sex        string
	Tel        string
	Birthday   time.Time
	Password   string
	Createtime time.Time
	ChangeTime time.Time
}

func (u UserDescribe) ValidatePassword(password string) bool {
	fmt.Println("utils.Md5New(password):", utils.Md5New(password))
	fmt.Println("u.Passwd:", u.Password)
	return utils.Md5New(password) == u.Password
}

func GetUsers(q string) []UserDescribe {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select id,name,birthday,sex,tel,addr,`desc`,create_time,change_time from todolist_user1")
	if err != nil {
		panic(err)
	}
	users := make([]UserDescribe, 0)
	for rows.Next() {
		var user UserDescribe
		if err := rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.Desc,
			&user.Createtime, &user.ChangeTime); err == nil {
			if q == "" ||
				strings.Contains(user.Name, q) ||
				strings.Contains(user.Tel, q) ||
				strings.Contains(user.Addr, q) ||
				strings.Contains(user.Desc, q) {
				fmt.Println("GetUsers:", user.Name, user.Tel, user.Addr, user.Desc)
				users = append(users, user)
			}
		}
	}
	return users
}
func GetUserByName(name string) (UserDescribe, error) {
	var user UserDescribe
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow("select id,name,password,birthday,sex,tel,addr,create_time from todolist_user1 where name=?", name)
	err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.Createtime)
	fmt.Println(user, err)
	return user, err
}

func ValidateCreateUser(name, password string, birthday, sex, tel, desc, addr string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6~30位之间"
	}
	_, err := strconv.ParseInt(tel, 11, 0)
	if err != nil {
		errors["tel"] = "请确定你输入的是11位手机号码或者7位数的电话号码"
	} else if len(tel) != 7 && len(tel) != 11 {
		fmt.Println(len(tel))
		errors["tel"] = "电话号码必须是7位或者11位"
	}
	return errors
}

func CreateUser(name, password string, birthday, sex, tel, addr, desc string) {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqll, err := db.Exec("insert into todolist_user1(name,password,birthday,sex,tel,addr,`desc`,create_time,change_time) values(?,md5(?),?,?,?,?,?,?,?)",
		name, password, birthday, sex, tel, addr, desc, time.Now(), time.Now()) // time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(err, sqll)
	if err != nil {
		panic(err)
	}
}

// (id, name, birthday, tel, desc, addr, password)
// name, password, birthday, tel, desc, addr
func ModifyUser(id int, name, password, birthday, sex, tel, addr, desc string) {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	// update todolist_user set name='?',password='md5(password)',birthday='?',sex='?',tel='?',addr='上海',`desc`='淞虹路' where id = ?;
	sql1, err := db.Exec("update todolist_user1 set name=?,password=md5(?),birthday=?,sex=?,tel=?,addr=?,`desc`=?,change_time=? where id = ?;",
		name, password, birthday, sex, tel, addr, desc, time.Now(), id)
	fmt.Println(sql1)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
}
func GetUserByID(id int) (UserDescribe, error) {
	var user UserDescribe
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// "select id,name,password,birthday,sex,tel,addr,create_time from todolist_user1 where name=?"
	row := db.QueryRow("select id,name,password,birthday,sex,tel,addr,`desc`,create_time from todolist_user1 where id=?;", id)
	err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.Desc, &user.Createtime)
	fmt.Println(user, err)
	return user, err //errors.New("not found")
}

func DeleteUser(id int) {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("delete from todolist_user1 where id=?", id) // delete from todolist_user where id=14
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
}
