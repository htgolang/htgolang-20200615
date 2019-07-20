package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.TempDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.UserHomeDir())

	fmt.Println(os.Getwd())

}
