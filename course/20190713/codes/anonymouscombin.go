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

type Employee struct {
	User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%T, %#v\n", me, me)

	me02 := Employee{
		User: User{
			ID:   1,
			Name: "KK",
			Addr: Address{"西安", "锦业路", "001"},
		},
		Salary: 10000,
	}

	fmt.Printf("%T, %#v\n", me02, me02)

	fmt.Println(me02.User.Name)
	fmt.Println(me02.User.Addr.Street)
	me02.User.Addr.Street = "未央路"
	fmt.Printf("%T, %#v\n", me02, me02)

	fmt.Println(me02.Name)
	fmt.Println(me02.Addr.Street)
	me02.Addr.No = "0003"
	me02.Name = "小k"
	fmt.Printf("%T, %#v\n", me02, me02)

}
