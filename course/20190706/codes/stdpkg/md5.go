package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	bytes := md5.Sum([]byte("i'amkk"))

	x := fmt.Sprintf("%X", bytes)

	fmt.Println(hex.EncodeToString(bytes[:]))
	fmt.Println(x)

	m := md5.New()
	m.Write([]byte("i'am"))
	m.Write([]byte("kk"))

	fmt.Printf("%x\n", m.Sum(nil))
}
