package tenant

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"github.com/xxdu521/cmdbgo/server/cloud"
)

type TenantCloud struct {
	addr		string
	region		string
	accessKey	string
	secrectKey	string
	credential	*common.Credential
	profile		*profile.ClientProfile
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

	c.credential = common.NewCredential(c.accessKey,c.secrectKey)
	c.profile = profile.NewClientProfile()
	c.profile.HttpProfile.Endpoint = c.addr

}
func (c *TenantCloud) TestConnect() error {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)
	if err != nil {
		return err
	}
	request := cvm.NewDescribeRegionsRequest()
	_, err = client.DescribeRegions(request)
	return err
}
func (c *TenantCloud) GetInstance() []*cloud.Instance {
	return nil
}
func (c *TenantCloud) StartInstance(uuid string) error {
	return nil
}
func (c *TenantCloud) StopInstance(uuid string) error {
	return nil
}
func (c *TenantCloud) RestartInstance(uuid string) error {
	return nil
}

func init(){
	cloud.DefaultManager.Register(&TenantCloud{})
}