package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

type Company struct {
	ID     int
	Name   string
	Addr   Address
	Salary float64
}

type Employee struct {
	User
	Company
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%T, %#v\n", me, me)

	me.Company.Name = "BB"
	me.User.Name = "KK"
	fmt.Println(me.Company.Name)
	fmt.Println(me.User.Name)

	fmt.Printf("%#v\n", me)
}
