package birthday

import (
	"time"
)

// 将输入的字符串数据转换为时间类型，格式为2006-01-02，返回时间类型的值和错误信息
func Birthday_time(s string) (time.Time, error) {

	T, Err := time.Parse("2006-01-02", s)

	return T, Err

}
