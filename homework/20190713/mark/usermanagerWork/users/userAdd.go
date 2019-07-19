package users

import "fmt"

func AddUser(){
	id := getid()
	setname := inputUser()
	setname.id = id
	users[id] = setname
	fmt.Printf("[ok]添加%v成功\n",id)
}