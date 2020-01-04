package utils

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type Sms struct {
	endpoint  string
	secretId  string
	secretKey string
	appid     string
	sign      string
}

func NewSms(endpoint, secretId, secretKey, appid, sign string) *Sms {
	return &Sms{
		endpoint:  endpoint,
		secretId:  secretId,
		secretKey: secretKey,
		appid:     appid,
		sign:      sign,
	}
}

func (s *Sms) Send(templateId string, phones []string, params []string) error {
	credential := common.NewCredential(s.secretId, s.secretKey)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = s.endpoint
	client, err := sms.NewClient(credential, "", cpf)
	if err != nil {
		return err
	}

	phoneSet := []*string{}
	for _, phone := range phones {
		temp := phone
		phoneSet = append(phoneSet, &temp)
	}
	paramSet := []*string{}
	for _, param := range params {
		temp := string([]rune(param)[:12])
		paramSet = append(paramSet, &temp)
	}

	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = phoneSet
	request.TemplateID = &templateId
	request.SmsSdkAppid = &s.appid
	request.Sign = &s.sign
	request.TemplateParamSet = paramSet

	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Printf("%s", response.ToJsonString())
	return err
}