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

	_,err := c.ecsClient.DescribeInstances(request)
	if err != nil {
		return err
	}
	//fmt.Printf("response is %#v\n", response)
	return nil
}
func (c *AliCloud) GetInstance() []*cloud.Instance {
	var (
		offset int = 1
		limit int = 100
		total int = 2
		rt []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = c.getInstanceByOffsetLimit(offset, limit)
		if offset == 1 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
		offset += limit
	}
	return rt
}
func (c *AliCloud) getInstanceByOffsetLimit(offset , limit int) (int, []*cloud.Instance) {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	request.PageNumber = requests.NewInteger(offset)
	request.PageSize = requests.NewInteger(100)

	response, err := c.ecsClient.DescribeInstances(request)
	if err != nil {
		return 0,nil
	}

	total := response.TotalCount
	instances := response.Instances.Instance

	rt := make([]*cloud.Instance, len(instances))

	for index, instance := range instances {
		publicAddrs := make([]string, 0)
		privateAddrs := make([]string, 0)

		if "" != instance.EipAddress.IpAddress {
			publicAddrs = append(publicAddrs, instance.EipAddress.IpAddress)
		}
		publicAddrs = append(publicAddrs, instance.PublicIpAddress.IpAddress...)

		privateAddrs = append(privateAddrs, instance.InnerIpAddress.IpAddress...)
		privateAddrs = append(instance.VpcAttributes.PrivateIpAddress.IpAddress)

		rt[index] = &cloud.Instance{
			UUID:         instance.InstanceId,
			Name:         instance.InstanceName,
			Os:           instance.OSName,
			Mem:          int64(instance.Memory),
			CPU:          instance.Cpu,
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       c.transformStatus(instance.Status),
			CreatedTime:  instance.CreationTime,
			ExpiredTime:  instance.ExpiredTime,
		}
	}

	return total, rt

}
func (c *AliCloud) transformStatus(status string) string {
	smap := map[string]string{
		"Running":  cloud.StatusRunning,
		"Stopped":  cloud.StatusStopped,
		"Starting": cloud.StatusStarting,
		"Stopping": cloud.StatusStopping,
	}

	if rt, ok := smap[status]; ok {
		return rt
	}
	return cloud.StatusUnknow
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

