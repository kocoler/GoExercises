package main

import (
	"fmt"
	"net/url"
)

func main() {
	//v := url.Values{}
	//v.Add("avatar", "http://thirdqq.qlogo.cn/g?b=oidb&k=m7HO2TOUZldUicHIIcPN9Ng&s=40&t=1563984011")
	//body := v.Encode()
	//fmt.Print(body)
	//m, err := url.ParseQuery("avatar=http%3a%2f%2fthirdqq.qlogo.cn%2fg%3fb%3doidb%26k%3dm7HO2TOUZldUicHIIcPN9Ng%26s%3d40%26t%3d1563984011")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(m["avatar"])
	//str, _ := url.QueryUnescape("http%3a%2f%2fthirdqq.qlogo.cn%2fg%3fb%3doidb%26k%3dm7HO2TOUZldUicHIIcPN9Ng%26s%3d40%26t%3d1563984011")
	//fmt.Println(str)
	str := url.QueryEscape("http://thirdqq.qlogo.cn/g?b=oidb&k=m7HO2TOUZldUicHIIcPN9Ng&s=40&t=1563984011")
	fmt.Println(str)
}
