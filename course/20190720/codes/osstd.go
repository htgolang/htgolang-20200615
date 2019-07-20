package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//os.Stdin
	//os.Stdout
	fmt.Println("xxx")

	os.Stdout.Write([]byte("xxx"))

	bytes := make([]byte, 5)

	n, err := os.Stdin.Read(bytes)

	fmt.Println(n, err, bytes)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Fprintf(os.Stdout, "%T", 1)
}
