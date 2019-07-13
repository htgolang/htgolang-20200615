package main

import "fmt"

type User struct {
	id   int
	name string
}

func (user User) GetID() int {
	return user.id
}
func (user User) GetName() string {
	return user.name
}

func (user *User) SetID(id int) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

type Employee struct {
	User
	Salary float64
	name   string
}

func main() {
	var me Employee = Employee{
		User:   User{1, "KK"},
		Salary: 1000,
	}

	fmt.Println(me.User.GetName()) // kk , ""
	fmt.Println(me.GetName())      // kk, ""
	me.SetName("silence")          // employee name, employee.user.name
	fmt.Println(me.GetName())      // "silence", ""
}
