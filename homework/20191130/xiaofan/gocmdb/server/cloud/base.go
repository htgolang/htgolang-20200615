package cloud

type Instance struct {
	Key          string
	UUID         string
	Name         string
	OS           string
	CPU          int
	Memory       int
	PublicAddrs  []string
	PrivateAddrs []string
	Status       string
	CreatedTime  string
	ExpiredTime  string
}

type ICloud interface {
	Type() string
	Name() string
	Init(string, string, string, string)
	TestConnect() error
	GetInstance() []*Instance
	StartInstance(string) error
	StopInstance(string) error
	RebootInstance(string) error
}
