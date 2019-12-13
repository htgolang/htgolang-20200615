package main

import (
	"fmt"
	"os/exec"
)

func main(){
	cmd := exec.Command("ping","-c","2","www.baidu.com")
	ctx,_ := cmd.Output()
	fmt.Println(string(ctx))
}
