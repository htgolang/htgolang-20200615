package modules

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"todolist/utils"
)

type User struct {
	Id          int
	Name        string
	Password    string
	Sex         string
	Birthday    time.Time
	Tel         string
	Addr        string
	Desc        string
	Create_time time.Time
}

func (u User) VaildatePassword(password string) bool {
	return u.Password == utils.Md5(password)
}

func GetUsers(q string) []User {
	var user User
	var users []User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	if q != "" {
		sql := "select id, name, birthday, tel, addr, `desc` from user where name like '%?%' || addr like '%?%' || `desc` like '%?%' || tel like '%?%'"
		rows, _ := db.Query(sql, )
		for rows.Next() {
			rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc)
			fmt.Println(user.Id, user.Name, user.Birthday, user.Tel, user.Addr, user.Desc)
			users = append(users, user)
		}
	} else {
		rows, _ := db.Query("select id, name, birthday, tel, addr, `desc` from user")
		for rows.Next() {
			rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc)
			fmt.Println(user.Id, user.Name, user.Birthday, user.Tel, user.Addr, user.Desc)
			users = append(users, user)
		}
	}
	return users
}

func GetUserByName(name string) (User, bool) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	row := db.QueryRow("select id, name, password from user where name=?", name)
	err = row.Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return user, false
	}
	fmt.Println("user:", user)
	return user, true
}

func ValidateCreateUser(name, password, password2, birthday, tel, desc, addr string) map[string]string {
	fmt.Println("validateCreateUser:", name, password, password2, birthday, tel, desc, addr)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	if password != password2 {
		errors["password"] = "两次密码输入不一样"
	} else if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6-30之间"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}

	return errors
}

func CreateUser(name, password, birthday, tel, desc, addr string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into user(name, password, birthday, tel, `desc`, addr, Create_time) values (?,md5(?),?,?,?,?,now())", name, password, birthday, tel, desc, addr)
	if err != nil {
		panic(err)
	}
}

func ValidateModifyUser(name string, birthday string) map[string]string {
	fmt.Println("ValidateModifyUser:", name, birthday)
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4-12之间"
	} else if _, err := GetUserByName(name); err {
		errors["name"] = "名称重复"
	}

	bir, _ := time.Parse("2006-01-02", birthday)
	start_time, _ := time.Parse("2006-01-02", "1960-01-01")
	if bir.Before(start_time) || bir.After(time.Now()) {
		errors["birthday"] = "出生日期必须在1960-01-01至今日之间"
	}
	return errors
}

func ModifyUser(id int, name, birthday, tel, desc, addr string) {
	fmt.Println("ModifyUser:", id)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	_, err = db.Exec("update user set name=?, birthday=?, tel=?, `desc`=?, addr=? where id = ?", name, birthday, tel, desc, addr, id)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(id int) {
	fmt.Println("delete user:", id)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("delete from user where id = ?", id)
	if err != nil {
		panic(err)
	}
}

func GetUserById(id int) (User, error){
	var user User
	fmt.Println("get user by id:", id)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	row := db.QueryRow("select id, name, password, sex, birthday, tel, addr, `desc` from user where id=?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Sex, &user.Birthday, &user.Tel, &user.Addr, &user.Desc)
	if err != nil {
		return user, err
	}
	fmt.Println("user:", user)
	return user, nil
}