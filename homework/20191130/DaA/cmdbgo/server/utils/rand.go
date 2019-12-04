package utils

import (
	"math/rand"
	"time"
)

func RandString(length int) string{
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWSYZ0123456789-_:@"
	chars := make([]byte,length)
	for i := 0; i< length; i++ {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}


func init(){
	rand.Seed(time.Now().UnixNano())
}