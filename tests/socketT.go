package main

import (
	"fmt"
	"net"
	//"syscall"
)

func main() {
	listener, err:= net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()          // 主协程结束时，关闭listener

	fmt.Println("服务器等待客户端建立连接...")
	// 等待客户端连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	defer conn.Close()             // 使用结束，断开与客户端链接
	fmt.Println("客户端与服务器连接建立成功...")

	// 接收客户端数据
	buf := make([]byte, 1024)        // 创建1024大小的缓冲区，用于read
	n, err := conn.Read(buf)
	fmt.Println(conn.RemoteAddr())
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("服务器读到:", string(buf[:n])) // 读多少，打印多少。
	conn.Write([]byte("WWW"))
	//syscall.Socket()
}
