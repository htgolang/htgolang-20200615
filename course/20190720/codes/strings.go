package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("abcdef12\n34567")

	// bytes := make([]byte, 3)
	// for {
	// 	n, err := reader.Read(bytes)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println(err)
	// 		}
	// 		break
	// 	} else {
	// 		fmt.Println(n, bytes[:n])
	// 	}
	// }
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	var builder strings.Builder

	builder.Write([]byte("abc"))
	builder.WriteString("abcdef!@#")

	fmt.Println(builder.String())
}
