package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/imsilence/testmod/gopkg"
)

func main() {
	fmt.Println(gopkg.Version)
	beego.Run()
}
