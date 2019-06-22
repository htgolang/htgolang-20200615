package main

import "fmt"

func main() {
	var marrays [3][2]int

	// 长度为2的int类型数组 => [2]int

	fmt.Println(marrays)
	fmt.Println(marrays[0])
	fmt.Println(marrays[0][0])

	marrays[0] = [2]int{1, 3}
	fmt.Println(marrays)

	marrays[1][1] = 1000
	fmt.Println(marrays)

	marrays = [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(marrays)

	for _, v := range marrays {
		for _, vv := range v {
			fmt.Print(vv, "\t")
		}
		fmt.Println()
	}
}
