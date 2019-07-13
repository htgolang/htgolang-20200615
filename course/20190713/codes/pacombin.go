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
	*User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%#v\n", me)

	me02 := Employee{
		User: &User{
			ID:   1,
			Name: "KK",
			Addr: Address{"西安市", "锦业路", "001"},
		},
		Salary: 10000,
		Name:   "小K",
	}

	fmt.Printf("%#v\n", me02)
	fmt.Println(me02.Name)
	fmt.Println(me02.Addr)

	fmt.Println(me02.User.Name)
	fmt.Println(me02.User.Addr)
}
