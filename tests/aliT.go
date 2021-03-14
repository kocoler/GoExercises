package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

)

func main() {
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "<accessKeyId>", "<accessSecret>")

	request := ecs.CreateDescribeInstancesFullStatusRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstancesFullStatus(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
