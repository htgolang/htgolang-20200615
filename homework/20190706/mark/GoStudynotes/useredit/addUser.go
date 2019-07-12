package GoStudynotes

import "fmt"

func AddUser(users map[int]map[string]string){
	id := getid(users)
	user := inputUser()
	users[id] = user
	fmt.Printf("[ok]添加%v成功\n",id)
}
