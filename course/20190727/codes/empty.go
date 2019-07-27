package main

import "fmt"

type EStruct struct{}

type Empty interface {
}

func fargs(args ...interface{}) {
	fmt.Println("------------------------")
	for _, arg := range args {
		fmt.Println(arg)
		switch v := arg.(type) {
		case int:
			fmt.Printf("Int: %T %v\n", v, v)
		case string:
			fmt.Printf("String: %T %v\n", v, v)
		default:
			fmt.Printf("Other: %T %v\n", v, v)
		}
	}
}

func main() {
	es := EStruct{}

	var e interface{} = 1

	fmt.Println(es, e)

	e = "tet"
	fmt.Println(e)

	e = true
	fmt.Println(e)

	e = es

	fmt.Println(e)
	fargs(1, "test", true, es)
}
