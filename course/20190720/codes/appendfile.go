package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	path := "user.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		file.WriteString("\n")
		file.Close()
	}
}
