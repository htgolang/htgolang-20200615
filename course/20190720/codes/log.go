package main

import (
	"log"
	"os"
	"time"
)

func main() {
	path := "user.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetPrefix("users:")
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Print(time.Now().Unix())
		file.Close()
	}

}
