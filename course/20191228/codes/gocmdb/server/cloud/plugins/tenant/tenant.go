package tenant

import (
	"fmt"
	"github.com/imsilence/gocmdb/server/cloud"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TenantCloud struct {
	addr string
	region string
	accessKey string
	secrectKey string
	credential *common.Credential
	profile *profile.ClientProfile
}

func (c *TenantCloud) Type() string {
	return "tenant"
}


func (c *TenantCloud) Name() string {
	return "腾讯云"
}

func (c *TenantCloud) Init(addr, region, accessKey, secrectKey string) {
	c.addr = addr
	c.region = region
	c.accessKey = accessKey
	c.secrectKey = secrectKey

	c.credential = common.NewCredential(c.accessKey, c.secrectKey)
	c.profile = profile.NewClientProfile()
    c.profile.HttpProfile.Endpoint = c.addr
}

func (c *TenantCloud) TestConnect() error {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request := cvm.NewDescribeRegionsRequest()
	_, err = client.DescribeRegions(request)
	fmt.Println(err)
	return err
}

func (c *TenantCloud) GetInstance() []*cloud.Instance {
	var (
		offset int64 = 0
		limit int64 = 100
		total int64 = 1
		rt []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = c.getInstanceByOffsetLimit(offset, limit)
		// 判断第一次请求初始化rt
		if offset == 0 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
		offset += limit
	}

	return rt
}


func (c *TenantCloud) transformStatus(status string)string {
	smap := map[string]string{
		"PENDING" : cloud.StatusPending,
		"LAUNCH_FAILED" : cloud.StatusLaunchFailed,
		"RUNNING" : cloud.StatusRunning,
		"STOPPED" : cloud.StatusStopped,
		"STARTING" : cloud.StatusStarting,
		"STOPPING" : cloud.StatusStopping,
		"REBOOTING" : cloud.StatusRebooting,
		"SHUTDOWN" : cloud.StatusShutdown,
		"TERMINATING" : cloud.StatusTerminating,
	}

	if rt, ok := smap[status]; ok {
		return rt
	}
	return cloud.StatusUnknow
}

func (c *TenantCloud) getInstanceByOffsetLimit(offset, limit int64) (int64, []*cloud.Instance) {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		// 通过log记录
		return 0, nil
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Limit = &limit
	request.Offset = &offset

	response, err := client.DescribeInstances(request)
	if err != nil {
		// 通过log记录
		return 0, nil
	}

	//总数
	total := *response.Response.TotalCount
	instances := response.Response.InstanceSet

	rt := make([]*cloud.Instance, len(instances))

	for index, instance := range instances {
		publicAddrs := make([]string, len(instance.PublicIpAddresses))
		privateAddrs := make([]string, len(instance.PrivateIpAddresses))
		for index, addr := range instance.PublicIpAddresses {
			publicAddrs[index] = *addr
		}

		for index, addr := range instance.PrivateIpAddresses {
			privateAddrs[index] = *addr
		}

		rt[index] = &cloud.Instance{
			UUID: *instance.InstanceId,
			Name: *instance.InstanceName,
			OS: *instance.OsName,
			CPU: int(*instance.CPU),
			Mem: *instance.Memory * 1024,
			PublicAddrs: publicAddrs,
			PrivateAddrs: privateAddrs,
			Status: c.transformStatus(*instance.InstanceState),
			CreatedTime: *instance.CreatedTime,
			ExpiredTime: *instance.ExpiredTime,
		}
	}

	return total, rt
}

func (c *TenantCloud) StartInstance(uuid string) error {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		return err
	}

	request := cvm.NewStartInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err = client.StartInstances(request)
	return err
}

func (c *TenantCloud) StopInstance(uuid string) error {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		return err
	}

	request := cvm.NewStopInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err = client.StopInstances(request)
	return err

}

func (c *TenantCloud) RebootInstance(uuid string) error {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		return err
	}

	request := cvm.NewRebootInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err = client.RebootInstances(request)
	return err
}


func init() {
	cloud.DefaultManager.Register(new(TenantCloud))
}