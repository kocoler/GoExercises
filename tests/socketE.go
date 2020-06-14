package main

import (
	"fmt"
	"net"
)

func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()         // 结束时，关闭连接

	// 发送数据
	_, err = conn.Write([]byte("Are u ready?"))
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
	buf := make([]byte, 1024)        // 创建1024大小的缓冲区，用于read
	_, err = conn.Read(buf)
	fmt.Println(string(buf))
}