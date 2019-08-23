package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// 用户（包含任务列表）
type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"username"`
	Tasks    []Task    `json:"tasks"`
	Birthday time.Time `json:"birthday"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
	Password string    `json:"password"`
}

// 任务
type Task struct {
	Id       int    `"json"": "id"`
	Name     string `"json"": "name"`
	Progress int    `"json"": "progress"`
	Desc     string `"json"": "desc"`
	Status   string `"json"": "status"`
}

func loadUsers() ([]User, error) {
	bytes, err := ioutil.ReadFile("datas/usertask.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, err
		}
		return nil, err
	} else {
		var users []User
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}

func storeUser(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/usertask.json", bytes, os.ModePerm)
}

func getUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for _, user := range users {
		if id < user.Id {
			id = user.Id
		}
	}
	return id + 1, nil
}

func ListUser() []User {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	return users
}

func CreateUser(username, password string, birthday time.Time, addr, desc string) error {
	users, err := loadUsers()
	fmt.Println("createUser LoadUser:", users, err)
	if err != nil {
		return err
	}

	var id int
	if len(users) == 0 {
		id = 1
	} else {
		// 判断用户名是否存在
		for _, user := range users {
			if user.UserName == username {
				return errors.New("用户名已经存在, 请重新添加用户.")
			}
		}
		id, err = getUserId()
		if err != nil {
			return err
		}
	}

	fmt.Println("create user. getid: ", id)

	users = append(users, User{
		Id:       id,
		UserName: username,
		Tasks:    []Task{},
		Birthday: birthday,
		Addr:     addr,
		Desc:     desc,
		Password: password,
	})

	return storeUser(users)
}

func CheckUserPwd(username, password string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if user.UserName == username && user.Password == password {
			return user, nil
		}
	}
	return User{}, errors.New("用户不存在")
}

func getTaskId(user User) int {
	var id int
	if len(user.Tasks) == 0 {
		id = 0
	} else {
		for _, task := range user.Tasks {
			if id < task.Id {
				id = task.Id
			}
		}
	}
	return id + 1
}

func CreateTask(username, name, desc string) error {
	users, err := loadUsers()
	if err != nil {
		return err
	}
	newUsers := make([]User, 0)
	for _, user := range users {
		if user.UserName == username {
			// 获取tasks列表，获取最新id
			id := getTaskId(user)
			fmt.Println("crate task id:", id)
			user.Tasks = append(user.Tasks, Task{
				Id:       id,
				Name:     name,
				Progress: 0,
				Desc:     desc,
				Status:   "new",
			})
			newUsers = append(newUsers, user)
		} else {
			newUsers = append(newUsers, user)
		}
	}
	fmt.Println("create task: ", newUsers)
	return storeUser(newUsers)
}

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return User{}, errors.New("Not Found")
}

func GetUserByName(username string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if user.UserName == username {
			return user, nil
		}
	}
	return User{}, errors.New("Not Found")
}

func ModifyUser(id int, name string, birthday time.Time, addr, desc string) error {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, 0)
	for _, user := range users {
		if user.Id == id {
			user.UserName = name
			user.Birthday = birthday
			user.Addr = addr
			user.Desc = desc
			newUsers = append(newUsers, user)
		} else {
			newUsers = append(newUsers, user)
		}
	}
	return storeUser(newUsers)
}

func DeleteUser(id int) error {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, 0)
	for _, user := range users {
		if user.Id != id {
			newUsers = append(newUsers, user)
		}
	}
	return storeUser(newUsers)
}

func ModifyUserPassword(username, password string) error {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, 0)
	for _, user := range users {
		if user.UserName == username {
			user.Password = password
			newUsers = append(newUsers, user)
		} else {
			newUsers = append(newUsers, user)
		}
	}
	fmt.Println("modify userpasswrd:", newUsers)
	return storeUser(newUsers)
}