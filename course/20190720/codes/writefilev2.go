package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("user.txt", []byte("xxxxxxxxxxxxxxxxxxxxx"), os.ModePerm)
	fmt.Println(err)
}
