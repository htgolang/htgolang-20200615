package main

import "fmt"

func main() {
	key_slice := []string{}

	string_map := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"}

	for k, _ := range string_map {
		key_slice = append(key_slice, k)
	}

	fmt.Println(key_slice)
}

/*
	评分: 8
*/
