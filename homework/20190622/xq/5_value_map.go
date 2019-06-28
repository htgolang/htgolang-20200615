package main

import (
	"fmt"
)

func main() {

	value_slice := []string{}

	string_map := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"}

	for _, v := range string_map {
		value_slice = append(value_slice, v)
	}

	fmt.Println(value_slice)

}

/*
	评分: 8
*/
