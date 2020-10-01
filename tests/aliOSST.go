package main

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	delete()
}

func delete() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "oss-cn-beijing.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := ""
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	objectName := "15006464_p0.jpg"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func upload() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "oss-cn-beijing.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := ""
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 创建存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	response, err := http.Get("https://www.gravatar.com/avatar/6384e2b2184bcbf58eccf10ca7a6563c?d=retro")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	//log.Println(string(contents))

	// 上传文件。
	err = bucket.PutObject("qq.png", bytes.NewReader(contents))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
