package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("user.txt")
	if err == nil {
		fmt.Println(string(bytes))
	}
}
