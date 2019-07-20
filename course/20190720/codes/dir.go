package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Mkdir("test01", 0644))
	// os.Rename("test01", "test02")
	// os.Remove("test02")
	fmt.Println(os.MkdirAll("test01/xxx", 0644))
	fmt.Println(os.RemoveAll("test01"))
}
