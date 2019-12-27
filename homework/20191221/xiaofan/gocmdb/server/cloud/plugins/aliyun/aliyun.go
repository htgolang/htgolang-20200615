package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/dcosapp/gocmdb/server/cloud"
)

type AliCloud struct {
	addr      string
	region    string
	accessKey string
	secretKey string
}

func (a *AliCloud) Type() string {
	return "AliCloud"
}

func (a *AliCloud) Name() string {
	return "阿里云"
}

// 初始化
func (a *AliCloud) Init(addr, region, accessKey, secretKey string) {
	a.addr = addr
	a.region = region
	a.accessKey = accessKey
	a.secretKey = secretKey
}

// 连接测试
func (a *AliCloud) TestConnect() error {
	client, err := ecs.NewClientWithAccessKey(a.region, a.accessKey, a.secretKey)
	if err != nil {
		return err
	}

	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"

	_, err = client.DescribeRegions(request)
	return err
}

// 获取实例
func (a *AliCloud) GetInstance() []*cloud.Instance {
	var (
		offset = 1
		limit  = 100
		total  = 2
		rt     []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = a.getInstanceByOffsetLimit(offset, limit)

		if offset == 0 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
		offset += limit
	}

	return rt
}

// 状态自定义
func (a *AliCloud) transformStatus(status string) string {
	smap := map[string]string{
		"Running":  cloud.StatusRunning,
		"Stopped":  cloud.StatusStopped,
		"Starting": cloud.StatusStarting,
		"Stopping": cloud.StatusStopping,
	}

	if rt, ok := smap[status]; ok {
		return rt
	}
	return cloud.StatusUnknown
}

// 获取实例
func (a *AliCloud) getInstanceByOffsetLimit(offset, limit int) (int, []*cloud.Instance) {
	client, err := ecs.NewClientWithAccessKey(a.region, a.accessKey, a.secretKey)
	if err != nil {
		return 0, nil
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	request.PageNumber = requests.NewInteger(offset)
	request.PageSize = requests.NewInteger(limit)

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
			CPU:          instance.Cpu,
			Memory:       int64(instance.Memory),
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       a.transformStatus(instance.Status),
			CreatedTime:  instance.CreationTime,
			ExpiredTime:  instance.ExpiredTime,
		}
	}

	return total, rt
}

// 启动实例
func (a *AliCloud) StartInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.region, a.accessKey, a.secretKey)
	if err != nil {
		return err
	}

	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = uuid

	_, err = client.StartInstance(request)
	return err
}

// 停止实例
func (a *AliCloud) StopInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.region, a.accessKey, a.secretKey)
	if err != nil {
		return err
	}

	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = uuid

	_, err = client.StopInstance(request)
	return err
}

// 重启实例
func (a *AliCloud) RebootInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.region, a.accessKey, a.secretKey)
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
	cloud.DefaultManager.Register(&AliCloud{})
}
