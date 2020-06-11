//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/28

/*
	UDP 服务端实现
	总结：
		1.读写操作都会与 udp 地址结构挂钩，因为UDP是一种无连接的协议。为了保证数据在网络中传输，每次读写操作都得传入udp地址结构体
		2. UDP 并没有阻塞监听的机制，但是会阻塞读socket

*/
package main

import (
	"fmt"
	"net"
)

func main() {

	// 创建一个 udp 地址结构。  指定服务器ip地址 端口号
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("服务器地址结构创建完成:", srvAddr)
	// 创建通信socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("服务器通信socket创建完成!")

	// 读取客户端发送的数据  此处阻塞，和TCP过程不太一样。
	buf := make([]byte, 4096)
	// 返回三个值:读取到字节数， 客户端 udp 地址结构， 错误
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("udpConn.ReadFromUDP err:", err)
		return
	}
	fmt.Printf("客户端地址：%v，客户端的数据：%s", cltAddr, string(buf[:n])) // 读多少显示多少

	// 向客户端回发数据，
	_, err = udpConn.WriteToUDP([]byte("Hello Client...."), cltAddr)
	if err != nil {
		fmt.Println("udpConn.WriteToUDP err:", err)
		return
	}
}
