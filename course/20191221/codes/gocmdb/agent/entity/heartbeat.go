package entity

import "time"

type Heartbeat struct {
	UUID string    `json:"uuid"`
	Time time.Time `json:"time"`
}

func NewHeartbeat(uuid string) Heartbeat {
	return Heartbeat{
		UUID: uuid,
		Time: time.Now(),
	}
}
