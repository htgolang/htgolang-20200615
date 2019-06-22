package main

import "fmt"

func main() {
	/*
	 1 * 1 = 1
	 1 * 2 = 2    2 * 2 = 4
	 1 * 3 = 3    2 * 3 = 6    3 * 3 = 9
	 ...
	 1 * 9 = 9      ...                          9 * 9 = 81
	*/
	for line := 1; line <= 9; line++ {
		for i := 1; i <= line; i++ {
			// fmt.Print(i, "*", line, "=", i*line, "\t")
			fmt.Printf("%d * %d = %-2d    ", i, line, i*line)
		}
		fmt.Println()
	}
}
