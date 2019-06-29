package main

import "fmt"

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func calc2(a, b int) (sum, diff, product, merchant int) {
	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b
	return
}

func main() {
	a, b, _, _ := calc(9, 3)
	fmt.Println(a, b)

	fmt.Println(calc2(5, 2))
}
