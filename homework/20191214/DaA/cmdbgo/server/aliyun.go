package main

import (
	"fmt"
	//"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

)

//import "github.com/aliyun/alibaba-cloud-sdk-go/tree/master/services/ecs"
//LTAI4Ftu6kGNth8dAqpPkVv1 rWOpA52tJZaATuyw3TMQJkjvnchdOK ecs.aliyuncs.com

func main() {
	// 创建ecsClient实例
	ecsClient, err := ecs.NewClientWithAccessKey(
		"cn-shenzhen",             // 地域ID
		"LTAI4FkfkEVNFGV7S3294foA",         // 您的Access Key ID
		"fzotL4uCygsstuie6WzUs0tIRd1Lfy")        // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	/*
	// 创建API请求并设置参数
	request := ecs.CreateDescribeInstancesRequest()
	// 等价于 request.PageSize = "10"
	request.PageSize = requests.NewInteger(10)
	// 发起请求并处理异常
	response, err := ecsClient.DescribeInstances(request)
	if err != nil {
		// 异常处理
		panic(err)
	}
	fmt.Printf("类型:%#\n值:%v\n",response.Instances,response.Instances)
	 */

	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = "i-wz962mggaelnnz3kupfw"
	response, err := ecsClient.DescribeInstanceAttribute(request)
	fmt.Printf("值:%v \n", response.ImageId)
}


