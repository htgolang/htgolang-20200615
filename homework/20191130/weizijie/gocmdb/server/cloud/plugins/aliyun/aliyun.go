package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/imsilence/gocmdb/server/cloud"
)

type AliyunCloud struct {
	addr       string
	region     string
	accessKey  string
	secrectKey string
	ecsClient  *ecs.Client
	request    *ecs.DescribeInstancesRequest
}

func (c *AliyunCloud) Type() string {
	return "aliyun"
}

func (c *AliyunCloud) Name() string {
	return "阿里云"
}

func (c *AliyunCloud) Init(addr, region, accessKey, secrectKey string) {

	c.addr = addr
	c.region = region
	c.accessKey = accessKey
	c.secrectKey = secrectKey

	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		panic(err)
	}
	c.ecsClient = client

	c.request = ecs.CreateDescribeInstancesRequest()
	c.request.PageSize = requests.NewInteger(10)
}

func (c *AliyunCloud) TestConnect() error {
	_, err := c.ecsClient.DescribeInstances(c.request)
	if err != nil {
		return err
	}
	return nil
}

func (c *AliyunCloud) GetInstance() []*cloud.Instance {
	return nil
}

func (c *AliyunCloud) StartInstance(uuid string) error {
	return nil
}

func (c *AliyunCloud) StopInstance(uuid string) error {
	return nil
}

func (c *AliyunCloud) RebootInstance(uuid string) error {
	return nil

}

func init() {
	cloud.DefaultManager.Register(new(AliyunCloud))
}
