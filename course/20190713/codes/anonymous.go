package main

import "fmt"

func main() {
	var me struct {
		ID   int
		Name string
	}

	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)
	fmt.Println(me.ID)
	me.Name = "kk"
	fmt.Printf("%#v\n", me)

	me2 := struct {
		ID   int
		Name string
	}{1, "KK"}

	fmt.Printf("%#v\n", me2)
}
