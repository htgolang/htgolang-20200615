package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串 => 其他类型
	// => bool

	if v, err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	// int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	if v, err := strconv.ParseInt("64", 16, 64); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	// float
	if v, err := strconv.ParseFloat("1.1", 64); err == nil {
		fmt.Println(v)
	}

	sd := fmt.Sprintf("%d", 12)
	fmt.Println(sd)

	sf := fmt.Sprintf("%.3f", 12.01)
	fmt.Println(sf)

	fmt.Printf("%q\n", strconv.FormatBool(false))
	fmt.Printf("%q\n", strconv.Itoa(12))
	fmt.Printf("%q\n", strconv.FormatInt(12, 16))
	fmt.Printf("%q\n", strconv.FormatFloat(10.1, 'f', -1, 64))
}
