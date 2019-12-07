package cloud

type Instance struct {}

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