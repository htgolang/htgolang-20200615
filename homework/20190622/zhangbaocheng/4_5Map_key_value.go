package main

import "fmt"

func main() {
	slice_key := []string{}
	slice_value := []int{}
	msg_map_ := map[string]int{"老王": 18, "老李": 17, "老三": 16}

	for key, value := range msg_map_ {
		slice_key = append(slice_key, key)
		slice_value = append(slice_value, value)
	}
	fmt.Printf("map slice key: %s \n", slice_key)
	fmt.Printf("map slice value: %d \n", slice_value)

}

/*
	评分: 8
*/
