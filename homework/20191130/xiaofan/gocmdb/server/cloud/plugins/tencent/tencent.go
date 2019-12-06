package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"gocmdb/cloud"
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
	return "tencent"
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
	var limit int64 = 100
	client, err := cvm.NewClient(t.credential, t.region, t.profile)
	if err != nil {
		return nil
	}
	request := cvm.NewDescribeInstancesRequest()
	request.Limit = &limit

	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil
	}

	instances := make([]*cloud.Instance, *response.Response.TotalCount)
	for index, instance := range response.Response.InstanceSet {
		publicAddrs := make([]string, len(instance.PublicIpAddresses))
		privateAddrs := make([]string, len(instance.PrivateIpAddresses))
		for i, addr := range instance.PublicIpAddresses {
			publicAddrs[i] = *addr
		}
		for i, addr := range instance.PrivateIpAddresses {
			privateAddrs[i] = *addr
		}
		instances[index] = &cloud.Instance{
			Key:          *instance.InstanceId,
			UUID:         *instance.Uuid,
			Name:         *instance.InstanceName,
			OS:           *instance.OsName,
			CPU:          int(*instance.CPU),
			Memory:       int(*instance.Memory),
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       *instance.InstanceState,
			CreatedTime:  *instance.CreatedTime,
			ExpiredTime:  *instance.ExpiredTime,
		}
	}
	return instances
}

// 启动实例
func (t *TencentCloud) StartInstance(string) error {
	return nil
}

// 停止实例
func (t *TencentCloud) StopInstance(string) error {
	return nil
}
func (t *TencentCloud) RebootInstance(string) error {
	return nil
}

func init() {
	cloud.DefaultManager.Register(&TencentCloud{})
}
