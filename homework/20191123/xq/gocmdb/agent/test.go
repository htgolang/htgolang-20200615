package main


import (
	"fmt"
	"time"
)

func main(){
	a, err := time.Parse("1/2/2006", "1/3/2019")

	fmt.Println(a, err)


}

