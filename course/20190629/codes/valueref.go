package main

import "fmt"

func main() {
	// 值类型  引用类型
	// 将变量赋值给新的一个变量，并修改新变量的值，如果对旧变量有影响 引用类型， 无影响值类型
	array := [3]string{"A", "B", "C"}
	slice := []string{"A", "B", "C"}

	arrayA := array
	sliceA := slice

	arrayA[0] = "Z"
	sliceA[0] = "Z"
	fmt.Println(arrayA, array)
	fmt.Println(sliceA, slice)

	fmt.Printf("%p %p\n", &arrayA, &array)
	fmt.Printf("%p %p\n", &sliceA, &slice)

	//int bool float array slice map 指针

	// 值类型: int bool float array 指针
	// 引用类型: slice map

	m := map[string]string{}

	mA := m
	mA["kk"] = "西安"
	fmt.Println(mA, m)
	fmt.Printf("%p %p\n", &mA, &m)
	age := 30

	ageA := age
	ageA = 31

	fmt.Println(ageA, age)
	pointer := &age

	*pointer = 31

	fmt.Println(age, *pointer)
}
