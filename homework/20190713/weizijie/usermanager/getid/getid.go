package getid

import (
	"github.com/JevonWei/usermanager/userstruct"
)

// 获取用户的最大ID，且返回ID+1
func GetId() int {
	var Id int
	for k := range userstruct.User {
		if Id < k {
			Id = k
		}
	}
	return Id + 1
}
