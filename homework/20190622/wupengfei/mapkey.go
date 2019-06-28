package main

import "fmt"

func main() {
	var maptask map[string]string = map[string]string{"name": "wpfs", "age": "18"}
	mapslice := []string{}
	for k, _ := range maptask {
		mapslice = append(mapslice, k)
	}
	fmt.Println(mapslice)
}

/*
 评分: 7
*/
