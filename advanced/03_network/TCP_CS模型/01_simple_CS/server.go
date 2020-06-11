//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/28

package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器通信协议及ip地址 端口号，并不是监听客户端。创建一个用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待建立客户端连接 2....")
	// 阻塞监听客户端连接请求，真正监听客户端。成功建立连接，返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器于客户端成功建立连接....")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器已经读取到客户端信息：", string(buf[:n]))

	// 回发数据，将读到数据原原本本回发
	conn.Write(buf[:n])
}
