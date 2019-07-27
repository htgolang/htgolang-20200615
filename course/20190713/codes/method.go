package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog Dog) SetName(name string) {
	dog.name = name
}

func (dog *Dog) PsetName(name string) {
	dog.name = name
}

func (dog Dog) test(dog2 Dog) {

}

func main() {

	var a int
	dog := Dog{"豆豆"}

	dog.Call()

	// dog.Name = "小黑"
	// dog.Call()
	dog.SetName("小黑")
	dog.Call()

	// (&dog).PsetName("小黑") 取引用
	dog.PsetName("小黑") // 自动取引用 语法糖
	dog.Call()

	pdog := &Dog{"豆豆"}
	// (*pdog).Call()
	pdog.Call()
	pdog.PsetName("小黑")
	// (*pdog).Call() // 解引用
	pdog.Call() // 自动解引用 语法糖

}
