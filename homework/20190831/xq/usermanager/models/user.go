package models

import (
	"crypto/md5"
	"encoding/hex"

	"errors"
	"fmt"

	"strconv"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID int
	Name string
	Password string
	Sex int
	Brithday string
	Tel string
	Addr string
	Descs string
	Create_time time.Time

}


const (
	dbUser string = "root"
	dbPass string = "123456"
	dbHost string = "127.0.0.1"
	dbPort int = 3306
	dbName string = "usermanager"

)

func OpenDB() (success bool, db *sql.DB){

	var isOpen bool

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser,dbPass,dbHost,dbPort,dbName)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		isOpen = false
	}else {
		isOpen = true
	}


	return isOpen, db

}
func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func AutoUserFromDB(username, password string) (User, error) {

	var user User

	opend, db := OpenDB()
	defer db.Close()

	if opend {

		sql := "select id,name,password, sex, brithday, tel, addr, descs, create_time from users where name=? and password=md5(?)"

		rows := db.QueryRow(sql, username, password)

		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Sex, &user.Brithday, &user.Tel,
			&user.Addr, &user.Descs, &user.Create_time)

		if err != nil {

			return user, err
		}

		return user,nil

	}else{
		fmt.Println("open db fails")
		return user, nil
	}

}

func GetUserListFromDB(q string) []User {
	opend, db := OpenDB()
	defer db.Close()

	var user User
	newUsers := make([]User, 0)

	if opend {

		sql := "select id, name, password, sex, brithday, tel, addr, descs, create_time from users"

		rows, err:= db.Query(sql)

		CheckErr(err)

		for rows.Next(){
			err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Sex, &user.Brithday, &user.Tel,
				&user.Addr, &user.Descs, &user.Create_time)
			if err == nil {
				if q == "" || strings.Contains(user.Name, q) ||
					strings.Contains(user.Tel, q) || strings.Contains(user.Addr, q) ||
					strings.Contains(user.Descs, q) {
					newUsers = append(newUsers, user)
				}
			}

		}
		return newUsers

	}else {
		return nil
	}


}

func GetUserIdFromDB(name string) (error) {
	var user User

	opend, db := OpenDB()

	if opend {
		sql := "select id from users where name=?"
		rows := db.QueryRow(sql, name)
		err := rows.Scan(&user.ID)
		if err == nil {
			return errors.New("用户存在")
		}else {
			return nil
		}

	}
	return nil
}

func GetUserFromDB(id int) (User, error) {

	var user User


	opend, db := OpenDB()

	if opend {

		sql := "select id, name, password, sex, brithday, tel, addr, descs from users where id=?"

		rows := db.QueryRow(sql, id)

		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Sex, &user.Brithday, &user.Tel,
			&user.Addr, &user.Descs)

		if err == nil {

			return user, err
		}

		return user,nil

	}else{
		fmt.Println("open db fails")
		return user, nil
	}

}

func InsertUserToDB(name, password, sex, brithday, tel, addr, descs string){
	opend, db := OpenDB()
	defer db.Close()

	if opend {

		sql := "insert into users (name, password, brithday, tel, addr, descs, sex, create_time) values(?,md5(?),?,?,?,?,?,?)"
		sexs,_:= strconv.Atoi(sex)

		result, err := db.Exec(sql, name, password, brithday, tel, addr, descs, sexs, time.Now())
		if err != nil {
			panic(err)
		}
		result.LastInsertId()
		result.RowsAffected()
	}

}

func ModifyUserFromDB(id int, name, sex, brithday, tel, addr, descs string){
	opend, db := OpenDB()
	if opend{

		sql := "update users set name=?,brithday=?, tel=?, addr=?, descs=?, sex=? where id=?"
		sexs,_:= strconv.Atoi(sex)
		result, err := db.Exec(sql, name, brithday, tel, addr, descs, sexs, id)
		if err != nil {
			panic(err)
		}
		result.LastInsertId()
		result.RowsAffected()

	}

}

func ModifyPassFromDB(id int, password string){

	fmt.Println(id, password)

	opend, db := OpenDB()
	defer db.Close()
	if opend {
		sql := "update users set password=md5(?) where id=?"
		result, err := db.Exec(sql, password, id)
		if err != nil {
			panic(err)
		}
		result.LastInsertId()
		result.RowsAffected()
	}

}
func DeleteUserFromDB(id int){
	opend, db := OpenDB()
	defer db.Close()

	if opend {
		sql := "delete from users where id=?"

		result, err := db.Exec(sql, id)
		if err != nil {
			panic(err)
		}

		result.LastInsertId()
		result.RowsAffected()


	}

}


func md5Pass(pass string) string{
	ctx := md5.New()
	ctx.Write([]byte(pass))
	return hex.EncodeToString(ctx.Sum(nil))
}




