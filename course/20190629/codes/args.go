package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addN(a, b int, args ...int) int {
	total := a + b
	for _, v := range args {
		total += v
	}
	return total
}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...)
	}
	return -1
}

func main() {
	fmt.Println(add(1, 5))
	fmt.Println(addN(1, 4))
	fmt.Println(addN(1, 4, 5))
	fmt.Println(addN(1, 4, 5, 8))

	fmt.Println(calc("add", 1, 2))
	fmt.Println(calc("add", 1, 2, 5))
	fmt.Println(calc("add", 1, 2, 5, 8))

	args := []int{1, 2, 5, 6, 10}
	fmt.Println(addN(1, 2, args...))
	fmt.Println(addN(1, 2, 1, 2, 5, 6, 10))
	fmt.Println(calc("add", 1, 3, args...))

	nums := []int{1, 3, 5, 8}
	fmt.Println(nums[:1])
	// nums[:1] + nums[2:]
	fmt.Println(nums[2:])
	// nums = append(nums[:1], 5, 8)
	nums = append(nums[:1], nums[2:]...)

	fmt.Println(nums)
}
