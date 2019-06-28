package main

import "fmt"

func main() {
	var maptask map[string]string = map[string]string{"name": "wpfs", "age": "18"}
	mapslice := []string{}
	for _, v := range maptask {
		mapslice = append(mapslice, v)
	}
	fmt.Println(mapslice)
}

/*
 评分: 8
*/
