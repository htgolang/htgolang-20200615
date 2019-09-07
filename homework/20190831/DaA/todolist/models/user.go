package models

import (
	"database/sql"
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

var sqlStr string
var user User

func GetUsers(query string) []User {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	var newUser []User

	if query == "" {
		sqlStr = fmt.Sprintf("select id, name, birthday, tel, addr, `desc`, create_time from users")
		rows, err := db.Query(sqlStr)
		if err == nil {
			for rows.Next() {
				_ = rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
				newUser = append(newUser, user)
			}
			return newUser
		}
	} else {
		sqlStr = "select id,name,birthday,tel,addr,`desc`,create_time from users where name like ?  or addr like ? or `desc` like ? "
		row := db.QueryRow(sqlStr, query, query, query)
		_ = row.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
		newUser = append(newUser, user)
		return newUser
	}
	return make([]User, 0)
}

func CreateUser(name, passwd, birthday, sex, tel, addr, desc string) error {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	newbirthday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return err
	}
	newsex, err := strconv.Atoi(sex)
	if err != nil {
		return err
	}

	sqlStr = "insert into users(name, password, birthday, sex, tel, addr, `desc`, create_time) values(?, md5(?), ?, ?, ?, ?, ?, now())"
	if _, err = db.Exec(sqlStr, name, passwd, newbirthday, newsex, tel, addr, desc); err != nil {
		return err
	}
	//错误返回测试
	//return errors.New("验证错误返回")
	return nil
}

func GetUserById(id int) (User, error) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return User{}, err
	}
	if err := db.Ping(); err != nil {
		return User{}, err
	}

	sqlStr := "select id, name, birthday, tel, addr, `desc`, create_time from users where id=?"
	row := db.QueryRow(sqlStr, id)
	err = row.Scan(&user.Id, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
	return user, err
}

func ModifyUser(id, name, birthday, sex, tel, addr, desc string) error {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	newid, err := strconv.Atoi(id)
	if err != nil {
		return err
	} else {
		newsex, err := strconv.Atoi(sex)
		newbirthday, err := time.Parse("2006-01-02", birthday)
		sqlStr = "update users set name=?, birthday=?, sex=?, tel=?, addr=?, `desc`=? where id=?"
		_, err = db.Exec(sqlStr, name, newbirthday, newsex, tel, addr, desc, newid)
		if err != nil {
			return err
		}
	}
	//return errors.New("测试")
	return nil
}

func DeleteUser(id string) error {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	newid, err := strconv.Atoi(id)
	if err != nil {
		return err
	} else {
		sqlStr = "delete from users where id=?"
		_, err = db.Exec(sqlStr, newid)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckUser(username, password string) error {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	sqlStr = "select name, password from users where name=? and password=md5(?)"
	row := db.QueryRow(sqlStr, username, password)
	err = row.Scan(&user.Name, &user.Password)
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}
