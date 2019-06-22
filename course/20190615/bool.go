package main

import "fmt"

func main() {
	// 布尔类型 表示真假
	// 标识符: bool
	// 字面量: true/false
	// 零值: false
	var zero bool
	isBoy := true
	isGirl := false

	fmt.Println(zero, isBoy, isGirl)
	// 操作
	// 逻辑运算(与 &&, 或 ||, 非 !)
	// aBool bBool
	// aBool && bBool 当两个都为True的时候结果为True
	//        bBool    true    false
	// aBool
	// true            true    false
	// false           false   false

	fmt.Println("&&")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	// aBool || bBool 只要有一个为True的时候结果为True
	//        bBool    true    false
	// aBool
	// true            true    true
	// false           true    false
	fmt.Println("||")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	// !aBool 取反

	fmt.Println("!")
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(!isBoy)
	fmt.Println(!isGirl)

	// 关系运算(==, !=)
	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != zero)

	fmt.Printf("%T %t %t\n", isBoy, isBoy, isGirl)
}
