package main

import (
	"fmt"

	"github.com/aliyun/aliyun-sts-go-sdk/sts"
	"os"
)

const (
	accessKeyID     = ""
	accessKeySecret = ""
	roleArn         = ``
	sessionName     = ""
)

func handleError(err error) {
	fmt.Println(err)
	os.Exit(-1)
}

func main() {
	stsClient := sts.NewClient(accessKeyID, accessKeySecret, roleArn, sessionName)

	resp, err := stsClient.AssumeRole(3600)
	if err != nil {
		handleError(err)
	}

	fmt.Printf("Credentials:\n")
	fmt.Printf("    AccessKeyID:%s\n", resp.Credentials.AccessKeyId)
	fmt.Printf("    AccessKeySecret:%s\n", resp.Credentials.AccessKeySecret)
	fmt.Printf("    SecurityToken:%s\n", resp.Credentials.SecurityToken)
	fmt.Printf("    Expiration:%s\n", resp.Credentials.Expiration)
}