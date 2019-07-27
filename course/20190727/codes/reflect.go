package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i int = 1

	fmt.Printf("%T\n", i)

	var typ reflect.Type = reflect.TypeOf(i)

	fmt.Println(typ.Name())
	fmt.Println(typ.Kind() == reflect.Int)
	fmt.Println(typ)
}
