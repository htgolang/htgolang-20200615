package utils

import (
	"math/rand"
	"time"
)


func RandString(length int) string {

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#"

	//count := len(letters)

	chars := make([]byte, length)

	for i := 0;i<length;i++ {
		chars[i] = letters[rand.Intn(len(letters))]
	}

	return string(chars)

}

func init()  {
	rand.Seed(time.Now().UnixNano())
}
