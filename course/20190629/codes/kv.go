package main

import "fmt"

func main() {
	users := map[string]int{"祥哥": 9, "萌新": 8, "烟灰": 9}

	keySlice := make([]string, len(users))
	valueSlice := []int{}

	i := 0
	for k, v := range users {
		keySlice[i] = k
		i++

		valueSlice = append(valueSlice, v)
	}
	fmt.Println(keySlice, valueSlice)

	for k, _ := range users {
		fmt.Println(k)
	}

	for k := range users {
		fmt.Println(k)
	}

	for _, v := range users {
		fmt.Println(v)
	}
}
