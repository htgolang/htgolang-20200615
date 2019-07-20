package main

import "os"

func main() {
	os.Rename("user.log", "user.v2.log")
	os.Remove("user.txt")
}
