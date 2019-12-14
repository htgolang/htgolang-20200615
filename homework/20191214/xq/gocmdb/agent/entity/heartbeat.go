package entity

import "time"

type Heartbase struct {
	UUID string `json:"uuid"`
	Time time.Time `json:"time"`
}

func NewHeartbeat(uuid string) Heartbase {
	return Heartbase{
		UUID:uuid,
		Time:time.Now(),
	}
}
