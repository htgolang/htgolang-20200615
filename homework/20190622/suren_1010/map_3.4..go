package main

import "fmt"

func main() {
	var mapData map[string]int
	var keySlice []string
	var valueSlice []int

	mapData = map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "F": 5}
	for k, v := range mapData {
		keySlice = append(keySlice, k)
		valueSlice = append(valueSlice, v)
	}
	fmt.Println(keySlice)
	fmt.Println(valueSlice)
}
