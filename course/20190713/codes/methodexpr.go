package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

/**
自动生成
func (dog *Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}
*/

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	m1 := (*Dog).Call
	//(*Dog).Call
	// m2 := Dog.SetName
	//Dog.SetName

	fmt.Printf("%T\n", m1)

	dog := Dog{"豆豆"}
	m1(&dog)
	dog.SetName("小白")
	m1(&dog)

	pdog := &Dog{"豆豆"}
	m1(pdog)
	// m2(pdog, "小黑")
	// m1(*pdog)
	// m2(pdog, "小白")
	pdog.SetName("小白")
	m1(pdog)
}
