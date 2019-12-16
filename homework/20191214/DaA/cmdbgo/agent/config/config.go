package config

import (
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	UUID string
	UUIDFile string

	Endpoint string
	Token string

	LogFile string

	PID int
	PidFile string

	Heartbeat 	chan interface{}
	Register 	chan interface{}
	Log 		chan interface{}
}

func NewConfig() (*Config, error) {
	UUIDFile := "run/agentd.uuid"
	PidFile := "run/agentd.pid"
	LogFile := "logs/agent.log"

	UUID := ""
	if cxt, err := ioutil.ReadFile(UUIDFile);err == nil {
		UUID = string(cxt)
	} else if os.IsNotExist(err) {
		UUID = strings.ReplaceAll(uuid.New().String(),"-","")
		ioutil.WriteFile(UUIDFile, []byte(UUID), os.ModePerm)
	} else {
		return nil,err
	}

	PID := os.Getpid()
	ioutil.WriteFile(PidFile,[]byte(strconv.Itoa(PID)),os.ModePerm)

	return &Config{
		Endpoint: "http://localhost:8080/v1/api",
		UUID: UUID,
		UUIDFile: UUIDFile,
		LogFile: LogFile,
		PID: PID,
		PidFile: PidFile,
		Heartbeat: make(chan interface{}, 64),
		Register: make(chan interface{}, 64),
		Log: make(chan interface{}, 102400),
	},nil

}
