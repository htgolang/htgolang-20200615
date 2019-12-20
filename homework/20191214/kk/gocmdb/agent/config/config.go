package config

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Config struct {
	UUID     string
	UUIDFile string

	Endpoint string
	Token    string

	LogFile string

	PID     int
	PidFile string

	Heartbeat chan interface{}
	Register chan interface{}
	Log chan interface{}
}

func NewConfig(reader *viper.Viper) (*Config, error) {
	UUIDFile := reader.GetString("uuid_file")
	if UUIDFile == "" {
		UUIDFile = "agentd.uuid"
	}
	PidFile := reader.GetString("pid_file")
	if PidFile == "" {
		PidFile = "agentd.pid"
	}
	LogFile := reader.GetString("log_file")
	if LogFile == "" {
		LogFile = "logs/agent.log"
	}
	Endpoint := reader.GetString("endpoint")
	if Endpoint == "" {
		Endpoint = "http://localhost:8888/v1/api"
	}
	Token := reader.GetString("token")

	UUID := ""
	if cxt, err := ioutil.ReadFile(UUIDFile); err == nil {
		UUID = string(cxt)
	} else if os.IsNotExist(err) {
		//strings.Replace(uuid.New().String(), "-", "", -1)
		UUID = strings.ReplaceAll(uuid.New().String(), "-", "")
		ioutil.WriteFile(UUIDFile, []byte(UUID), os.ModePerm)
	} else {
		return nil, err
	}

	PID := os.Getpid()
	ioutil.WriteFile(PidFile, []byte(strconv.Itoa(PID)), os.ModePerm)

	return &Config{
		Endpoint: Endpoint,
		Token: Token,
		UUID:     UUID,
		UUIDFile: UUIDFile,
		LogFile:  LogFile,
		PID:      PID,
		PidFile:  PidFile,
		Heartbeat: make(chan interface{}, 64),
		Register: make(chan interface{}, 64),
		Log: make(chan interface{}, 10240),
	}, nil
}
