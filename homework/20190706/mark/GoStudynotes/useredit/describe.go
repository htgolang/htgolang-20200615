package GoStudynotes

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)
func inputString(prompt string)string{
	//var input string
	//fmt.Print(prompt)
	//fmt.Scan(&input)
	//return strings.TrimSpace(input)

	c := make(chan string, 1)
	//go scanIn(c)
	var input string
	go func() {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		c <- input
	}()
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)
	select {
	case <-ctx.Done():
		fmt.Print("\ntime out")
		os.Exit(1)
	case <-c:
		return strings.TrimSpace(input)
	}
	return strings.TrimSpace(input)

}
func getid(users map[int]map[string]string) int{
	var id int
	for k,_ := range users {
		if id < k {
			id = k
		}
	}
	return id +1
}
func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = inputName("请输入名字:")
	user["birthday"] = InputYMD("请按照提示输入出生年月日!")
	user["tel"] = InputTel("请输入联系方式:")
	user["addr"]= inputString("请输入联系地址:")
	user["desc"] = inputString("请输入备注:")
	return user
}
func printUser(pk int,user map[string]string){
	fmt.Println("ID:",pk)
	fmt.Println("名称:",user["name"])
	fmt.Println("年龄:",(GetAge(GetTimeFromStrDate(user["birthday"]))),string('岁'))
	fmt.Println("联系方式:",user["tel"])
	fmt.Println("地址:",user["addr"])
	fmt.Println("出生日期:",user["birthday"])
	fmt.Println("备注信息:",user["desc"])
}

