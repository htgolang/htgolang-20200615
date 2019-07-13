package main

import "fmt"

type Counter int

type User map[string]string

type Callback func(...string)

func main() {
	var counter Counter = 20

	counter += 10
	fmt.Println(counter)

	me := make(User)
	me["name"] = "kk"
	me["addr"] = "西安市"

	fmt.Println(me)
	fmt.Printf("%T, %T\n", counter, me)

	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}

	list("a", "b", "c")

	var counter2 int = 10
	fmt.Println(int(counter) > counter2)
	fmt.Println(counter > Counter(counter2))
}
