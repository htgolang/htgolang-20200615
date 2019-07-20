package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "user.txt"
	file, err := os.Open(path)
	// fmt.Println(err)
	// fmt.Println(file)
	// fmt.Printf("%T\n", file)
	if err != nil {
		fmt.Println(err)
	} else {
		bytes := make([]byte, 20)

		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Print(string(bytes[:n]))
			}
		}
		file.Close()
	}
}
