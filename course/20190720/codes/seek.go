package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("user.txt")

	// 偏移量，相对位置
	// 文件开始 0 os.SEEK_SET
	// 当前位置 1 os.SEEK_CUR
	// 文件末尾 2 os.SEEK_END
	fmt.Println(file.Seek(5, 0))

	bytes := make([]byte, 100)
	n, err := file.Read(bytes)

	fmt.Println(n, err, string(bytes[:n]))

	n, err = file.Read(bytes)

	fmt.Println(n, err, bytes[:n])

	file.Close()
}
