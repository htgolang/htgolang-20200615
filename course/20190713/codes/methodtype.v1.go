package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	dog := Dog{"豆豆"}

	m1 := dog.Call //dog 拷贝
	fmt.Printf("%T\n", m1)
	m1()

	dog.SetName("小黑")
	dog.Call()
	m1()

	pdog := &Dog{"豆豆"}
	m2 := pdog.Call // pdog会自动解引用，拷贝
	fmt.Printf("%T\n", m2)
	m2()

	pdog.SetName("小黑")
	pdog.Call()
	m2()
}
