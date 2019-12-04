package alibaba

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/xxdu521/cmdbgo/server/cloud"
)

type AliCloud struct {
	addr		string
	region		string
	accessKey	string
	secrectKey	string
	ecsClient	*ecs.Client	//创建一个凭证对象，这里好像不对，这个对象不单单存储凭证。
	//request		*ecs.DescribeInstancesRequest	//创建一个请求对象
}

func (c *AliCloud) Type() string{
	return "alibaba"
}

func (c *AliCloud) Name() string {
	return "阿里云"
}

func (c *AliCloud) Init(addr, region, accessKey, secrectKey string) {
	c.addr = addr
	c.region = region
	c.accessKey = accessKey
	c.secrectKey = secrectKey

	//做一个credential对象。
	ecsClient, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		panic(err)
	}
	c.ecsClient = ecsClient

}

func (c *AliCloud) TestConnect() error {
	//创建一个请求对象，设置返回的数量
	request := ecs.CreateDescribeInstancesRequest()
	request.PageSize = requests.NewInteger(10)

	response,err := c.ecsClient.DescribeInstances(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}

func (c *AliCloud) GetInstance() []*cloud.Instance {
	return nil
}

func (c *AliCloud) StartInstance(uuid string) error {
	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = uuid

	response, err := c.ecsClient.StartInstance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}

func (c *AliCloud) StopInstance(uuid string) error{
	//"i-wz962mggaelnnz3kupfw"
	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = uuid

	response, err := c.ecsClient.StopInstance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}

func (c *AliCloud) RestartInstance(uuid string) error{
	request := ecs.CreateRebootInstanceRequest()
	request.Scheme = "https"
	request.InstanceId = uuid

	response, err := c.ecsClient.RebootInstance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}

func init(){
	cloud.DefaultManager.Register(&AliCloud{})
}

