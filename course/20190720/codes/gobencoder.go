package main

import (
	"encoding/gob"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
}

func main() {

	// users := map[int]User{
	// 	1: User{1, "烟灰", time.Now(), "123123123", "福建"},
	// 	2: User{2, "xq", time.Now(), "123123124", "上海"},
	// 	3: User{3, "祥哥", time.Now(), "123123123", "杭州"},
	// }

	stud := User{1, "烟灰", time.Now(), "123123123", "福建"}

	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(stud)
	}

}
