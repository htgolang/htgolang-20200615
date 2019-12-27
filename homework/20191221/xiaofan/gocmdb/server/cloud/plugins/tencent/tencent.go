package tencent

import (
	"github.com/dcosapp/gocmdb/server/cloud"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TencentCloud struct {
	addr       string
	region     string
	accessKey  string
	secretKey  string
	credential *common.Credential
	profile    *profile.ClientProfile
}

func (t *TencentCloud) Type() string {
	return "TencentCloud"
}

func (t *TencentCloud) Name() string {
	return "腾讯云"
}

// 初始化
func (t *TencentCloud) Init(addr, region, accessKey, secretKey string) {
	t.addr = addr
	t.region = region
	t.accessKey = accessKey
	t.secretKey = secretKey
	t.credential = common.NewCredential(
		accessKey,
		secretKey,
	)

	t.profile = profile.NewClientProfile()
	t.profile.HttpProfile.Endpoint = t.addr
}

// 连接测试
func (t *TencentCloud) TestConnect() error {
	client, err := cvm.NewClient(t.credential, t.region, t.profile)
	if err != nil {
		return err
	}
	request := cvm.NewDescribeRegionsRequest()
	_, err = client.DescribeRegions(request)
	return err
}

// 获取实例
func (t *TencentCloud) GetInstance() []*cloud.Instance {
	var (
		offset int64 = 0
		limit  int64 = 100
		total  int64 = 1
		rt     []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = t.getInstanceByOffsetLimit(offset, limit)

		if offset == 0 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
		offset += limit
	}

	return rt
}

// 自定义状态
func (t *TencentCloud) transformStatus(status string) string {
	smap := map[string]string{
		"PENDING":       cloud.StatusPending,
		"LAUNCH_FAILED": cloud.StatusLaunchFailed,
		"RUNNING":       cloud.StatusRunning,
		"STOPPED":       cloud.StatusStopped,
		"STARTING":      cloud.StatusStarting,
		"TOPPING":       cloud.StatusStopping,
		"REBOOTING":     cloud.StatusRebooting,
		"TERMINATING":   cloud.StatusTerminating,
		"SHUTDOWN":      cloud.StatusShutdown,
	}

	if rt, ok := smap[status]; ok {
		return rt
	}
	return cloud.StatusUnknown
}

// 获取实例
func (t *TencentCloud) getInstanceByOffsetLimit(offset, limit int64) (int64, []*cloud.Instance) {
	client, err := cvm.NewClient(t.credential, t.region, t.profile)
	if err != nil {
		return 0, nil
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Limit = &limit
	request.Offset = &offset

	response, err := client.DescribeInstances(request)
	if err != nil {
		return 0, nil
	}

	total := *response.Response.TotalCount
	instances := response.Response.InstanceSet

	rt := make([]*cloud.Instance, len(instances))

	for index, instance := range instances {
		publicAddrs := make([]string, len(instance.PublicIpAddresses))
		privateAddrs := make([]string, len(instance.PrivateIpAddresses))
		for i, addr := range instance.PublicIpAddresses {
			publicAddrs[i] = *addr
		}
		for i, addr := range instance.PrivateIpAddresses {
			privateAddrs[i] = *addr
		}
		rt[index] = &cloud.Instance{
			UUID:         *instance.InstanceId,
			Name:         *instance.InstanceName,
			OS:           *instance.OsName,
			CPU:          int(*instance.CPU),
			Memory:       *instance.Memory * 1024,
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       t.transformStatus(*instance.InstanceState),
			CreatedTime:  *instance.CreatedTime,
			ExpiredTime:  *instance.ExpiredTime,
		}
	}
	return total, rt
}

// 启动实例
func (t *TencentCloud) StartInstance(uuid string) error {
	client, _ := cvm.NewClient(t.credential, t.region, t.profile)

	request := cvm.NewStartInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err := client.StartInstances(request)

	return err
}

// 停止实例
func (t *TencentCloud) StopInstance(uuid string) error {
	client, _ := cvm.NewClient(t.credential, t.region, t.profile)

	request := cvm.NewStopInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err := client.StopInstances(request)

	return err
}

// 重启实例
func (t *TencentCloud) RebootInstance(uuid string) error {
	client, _ := cvm.NewClient(t.credential, t.region, t.profile)

	request := cvm.NewRebootInstancesRequest()
	request.InstanceIds = []*string{&uuid}

	_, err := client.RebootInstances(request)

	return err
}

func init() {
	cloud.DefaultManager.Register(&TencentCloud{})
}
