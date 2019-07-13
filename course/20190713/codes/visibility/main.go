package main

import (
	"fmt"
	"visibility/users"
)

func main() {
	var u users.User
	// var a users.address

	fmt.Printf("%#v\n", u)
	// fmt.Printf("%#v\n", a)

	fmt.Println(u.ID)
	fmt.Println(u.Name)
	// fmt.Println(u.birthday)
	// fmt.Println(u.addr)
}
