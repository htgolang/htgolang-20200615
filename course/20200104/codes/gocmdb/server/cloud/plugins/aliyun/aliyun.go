package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/imsilence/gocmdb/server/cloud"
)

type Aliyun struct {
	addr       string
	region     string
	accessKey  string
	secrectKey string
}

func (c *Aliyun) Type() string {
	return "aliyun"
}

func (c *Aliyun) Name() string {
	return "阿里云"
}

func (c *Aliyun) Init(addr, region, accessKey, secrectKey string) {
	c.addr = addr
	c.region = region
	c.accessKey = accessKey
	c.secrectKey = secrectKey
}

func (c *Aliyun) TestConnect() error {
	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		return err
	}

	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"

	_, err = client.DescribeRegions(request)
	return err
}

func (c *Aliyun) GetInstance() []*cloud.Instance {
	var (
		offset int = 1
		limit  int = 100
		total  int = 2
		rt     []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = c.getInstanceByOffsetLimit(offset, limit)
		if offset == 1 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
	}
	return rt
}

func (c *Aliyun) transformStatus(status string) string {
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

func (c *Aliyun) getInstanceByOffsetLimit(offset, limit int) (int, []*cloud.Instance) {
	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		return 0, nil
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	request.PageNumber = requests.NewInteger(offset)
	request.PageSize = requests.NewInteger(100)

	response, err := client.DescribeInstances(request)
	if err != nil {
		return 0, nil
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
			OS:           instance.OSName,
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

func (c *Aliyun) StartInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		return err
	}

	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = uuid

	_, err = client.StartInstance(request)
	return err
}

func (c *Aliyun) StopInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		return err
	}

	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = uuid

	_, err = client.StopInstance(request)
	return err
}

func (c *Aliyun) RebootInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(c.region, c.accessKey, c.secrectKey)
	if err != nil {
		return err
	}

	request := ecs.CreateRebootInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = uuid

	_, err = client.RebootInstance(request)
	return err
}

func init() {
	cloud.DefaultManager.Register(new(Aliyun))
}
