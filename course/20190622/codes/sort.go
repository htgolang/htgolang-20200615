package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 5, 7, 8, 6}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"test", "kk", "123", "zzz", "xxx"}
	sort.Strings(names)
	fmt.Println(names)

	heights := []float64{1.1, -1.1, 3.3, 2.2}
	sort.Float64s(heights)
	fmt.Println(heights)
}
