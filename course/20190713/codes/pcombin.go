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
	Addr *Address
}

func main() {
	var me01 User
	fmt.Printf("%#v\n", me01)

	me02 := User{
		ID:   1,
		Name: "KK",
		Addr: &Address{"西安市", "锦业路", "01"},
	}
	fmt.Printf("%#v\n", me02)

	fmt.Println(me02.Addr.Street)
	me02.Addr.Region = "北京市"
	fmt.Printf("%#v\n", me02.Addr)
}
