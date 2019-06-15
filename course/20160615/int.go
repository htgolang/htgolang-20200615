package main

import "fmt"

func main() {
	// 整数类型
	// 标识符: int/int*/uint/uint*/uintptr/byte/rune
	// 字面量: 十进制, 八进制0777 = 7* 8^2 + 7 * 8^1 + 7 * 8^0, 十六进制0X0-9A-F
	var age int = 31
	fmt.Printf("%T %d\n", age, age)

	fmt.Println(0777, 0X10)

	//操作
	// 算术运算(+, -, *, /, %, ++, --)
	fmt.Println(1 + 2)
	fmt.Println(3 - 10)
	fmt.Println(3 * 9)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)
	// fmt.Println(9 / 0)

	age++
	fmt.Println(age)

	age--
	fmt.Println(age)

	// 关系运算(== != > >= < <=)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 > 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 < 3)
	fmt.Println(2 <= 3)

	// 位运算 二进制的运算 10 => 2
	// & | ^ << >> &^
	// 十进制 => 2
	// 7/2 => 1 3/2 =>1 1/2 => 1 0 0111
	// 2 => 0010
	// 7 & 2 => 0111 & 0010 => 0010 => 2
	// 7 | 2 => 0111 | 0010 => 0111 => 7
	// 7 ^ 2 => 0111 ^ 0010 => 0101 => 5
	// 2 << 1 => 0010 << 1 => 0100 => 4
	// 2 >> 1 => 0010 >> 1 => 0001 => 1
	// 7 &^ 2 => 0111 &^ 0010 => 0101 => 5
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(7 ^ 2)
	fmt.Println(2 << 1)
	fmt.Println(2 >> 1)
	fmt.Println(7 &^ 2)

	// 赋值(=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=, &^=)
	// a+=b ==> a = a+b
	age = 1
	age += 3 // age = age+3
	fmt.Println(age)

	// int/uint/byte/rune/int*
	var intA int = 10
	var uintB uint = 3
	fmt.Println(intA + int(uintB))
	fmt.Println(uint(intA) + uintB)

	// 大 -> 小转换可能出现溢出
	var intC int = 0XFFFF
	fmt.Println(intC, uint8(intC), int8(intC))

	//fmt.Printf
	// int/unit/int*/uint*
	// byte, rune
	var a byte = 'A'
	var w rune = '中'
	fmt.Println(a, w)

	age = 21
	fmt.Printf("%T %d %b %o %x\n", age, age, age, age, age)
	fmt.Printf("%T %c \n", a, a)
	fmt.Printf("%T %q %U\n", w, w, w)
	fmt.Printf("%-5d\n", age)
	fmt.Printf("%05d\n", age)
}
