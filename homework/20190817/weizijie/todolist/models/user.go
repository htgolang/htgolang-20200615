package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	ID       int    `json:id`
	Name     string `json:name`
	Birthday string `json:birthday`
	Addr     string `json:addr`
	Tel      string `json:tel`
	Desc     string `json:desc`
	//Accountuser string `json:accountuser`
}

type Account struct {
	AccountName string
	Passwd      string
}

var AccountMap = make(map[string][]User)

func Init() {
	logfile := "user.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
	}
}

func loadAccount() ([]Account, error) {
	if bytes, err := ioutil.ReadFile("datas/accounts.json"); err != nil {
		if os.IsNotExist(err) {
			log.Println("accounts.json 文件不存在")
			return []Account{}, nil
		}
		return nil, err
	} else {
		var accounts []Account
		if string(bytes) == "" {
			return []Account{}, nil
		}
		if err := json.Unmarshal(bytes, &accounts); err == nil {
			log.Println("Account Load Success")
			return accounts, nil
		} else {
			log.Println("Account Load Failed")
			return nil, err
		}
	}
}

func Storepasswd(accounts []Account) error {
	bytes, err := json.Marshal(accounts)
	if err != nil {
		return err
	}
	log.Println("Password Store Success")
	return ioutil.WriteFile("datas/accounts.json", bytes, 0X066)
}

func GetAccountByname(name string) (Account, error) {
	accounts, err := loadAccount()
	if err != nil {
		panic(err)
	}
	for _, account := range accounts {
		if name == account.AccountName {
			log.Printf("Name为 %s 的Account Get Success", name)
			return account, nil
		}
	}
	log.Printf("Name为 %s 的Account Get Failed", name)
	return Account{}, errors.New("Account Not Found")
}

func ModifyPasswd(name, passwd string) {
	accounts, err := loadAccount()
	if err != nil {
		panic(err)
	}

	newAccounts := make([]Account, len(accounts))
	for i, account := range accounts {
		if name == account.AccountName {
			account.AccountName = name
			account.Passwd = passwd
		}
		newAccounts[i] = account
	}
	log.Printf("Name为 %s 的Passwd Modidy Success", name)
	Storepasswd(newAccounts)
}

func AccountCreate(name, passwd string) error {
	accounts, err := loadAccount()
	if err != nil {
		fmt.Println(err)
	}
	account := Account{
		AccountName: name,
		Passwd:      passwd,
	}

	for _, account := range accounts {
		if account.AccountName == name {
			log.Printf("%s Account Exist", name)
			return errors.New("Account Exist")
		}
	}

	accounts = append(accounts, account)
	Storepasswd(accounts)
	log.Printf("%s Account Create Success", name)
	return nil

}

// 重新赋值AccountMap
func RenewAccountMap(accountMaps map[string][]User) map[string][]User {
	new_accountMaps := accountMaps
	for i, _ := range accountMaps {
		new_accountMaps[i] = GetUsers()
	}

	return new_accountMaps
}

func GetAccount() []Account {
	accounts, err := loadAccount()
	//fmt.Println(accounts, err)
	if err == nil {
		//fmt.Printf("%T\n", accounts)
		return accounts
	}
	panic(err)
}

func AccountVerify(name, passwd string) (bool, error) {
	accounts := GetAccount()

	for _, account := range accounts {
		if account.AccountName == name {
			if account.Passwd == passwd {
				log.Printf("%s Account Verify Success", name)
				return true, nil
			} else {
				log.Printf("%s Account Verify Failed", name)
				return false, errors.New("Verity Failed")
			}
		}
	}
	log.Printf("%s Account Not Exist", name)
	return false, errors.New("Account not Exist")
}

func loadUsers() ([]User, error) {
	if bytes, err := ioutil.ReadFile("datas/user.json"); err != nil {
		if os.IsNotExist(err) {
			log.Println("datas/user.json 文件不存在")
			return []User{}, nil
		}
		return nil, err
	} else {
		var users []User
		if err := json.Unmarshal(bytes, &users); err == nil {
			log.Println("Users Load Success")
			return users, nil
		} else {
			log.Println("Users Load Failed")
			return nil, err
		}
	}
}

func storeUsers(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Println("Users Store Failed")
		return err
	}
	log.Println("Users Store Success")
	return ioutil.WriteFile("datas/user.json", bytes, 0X066)
}

func GetUsers() []User {
	users, err := loadUsers()
	//fmt.Println(users, err)
	if err == nil {
		return users
	}
	panic(err)
}

func GetUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for _, user := range users {
		if id < user.ID {
			id = user.ID
		}
	}
	return id + 1, nil
}

func CreateUser(name, birthday, addr, tel, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	user := User{
		ID:       id,
		Name:     name,
		Birthday: birthday,
		Desc:     desc,
		Addr:     addr,
		Tel:      tel,
	}
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users = append(users, user)
	storeUsers(users)
	log.Printf("%s user Create Success", name)
	// 重新更新AccountMap的值
	AccountMap = RenewAccountMap(AccountMap)
}

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if id == user.ID {
			return user, nil
		}
	}
	log.Printf("ID == %d: Users Get Success", id)
	return User{}, errors.New("Not Found")
}

func GetUserByName(name string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if name == user.Name {
			return user, nil
		}
	}
	log.Printf("Name == %s: Users Get Success", name)
	return User{}, errors.New("Not Found")
}

func ModifyUser(id int, name, birthday, addr, tel, desc string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, len(users))
	for i, user := range users {
		if id == user.ID {
			user.Name = name
			user.Desc = desc
			user.Birthday = birthday
			user.Addr = addr
			user.Tel = tel
		}
		newUsers[i] = user
	}
	storeUsers(newUsers)
	log.Printf("%s user Create Success", name)

	AccountMap = RenewAccountMap(AccountMap)
}

func DeleteUser(id int) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, 0)
	for _, user := range users {
		if id != user.ID {
			newUsers = append(newUsers, user)
		} else {
			log.Printf("%s user will Delete", user)
		}
	}
	log.Printf("ID=%d user Delete Success", id)
	storeUsers(newUsers)
	AccountMap = RenewAccountMap(AccountMap)
}
