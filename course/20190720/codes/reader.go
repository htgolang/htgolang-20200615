package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)

		bytes := make([]byte, 5)
		for {
			n, err := reader.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(n, err, string(bytes))
			}
		}
	}
}
