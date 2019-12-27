package entity

import "time"

type Heartbeat struct {
	UUID string    `json:"uuid"`
	Time time.Time `json:"time"`
}

// 返回uuid和当前时间给ens
func NewHeartbeat(uuid string) Heartbeat {
	return Heartbeat{
		UUID: uuid,
		Time: time.Now(),
	}
}
