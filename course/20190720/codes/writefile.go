package main

import (
	"fmt"
	"os"
)

func main() {
	path := "user2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		// 写文件
		file.Write([]byte("123"))
		file.WriteString("xyz")
		// 关闭文件
		file.Close()
	}
}
