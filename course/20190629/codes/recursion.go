package main

import "fmt"

func addN(n int) int {
	if n == 1 {
		return 1
	}
	return n + addN(n-1)
}

func main() {
	fmt.Println(addN(5))
}

// addN(5) => 5 + addN(4)
// addN(4) => 4 + addN(3)
// addN(3) => 3 + addN(2)
// addN(2) => 2 + addN(1)
// addN(1) => 1

// addN(1) = 1 + addN(0)
// addN(0) = 0 + addN(-1)
//....
