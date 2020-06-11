//File  : client.go
//Author: duanhaobin
//Date  : 2020/5/28

/*
	UDP	客户端实现
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	var (
		n int
	)
	// 指定服务器 ip+port，创建通信socket
	conn, err := net.Dial("udp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 向服务端端发出通信内容
	n, err = conn.Write([]byte("Hello 我是客户端信息"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	// 接收服务端回发的内容
	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("接收服务端信息异常,err:", err)
		return
	}

	// 处理服务端的内容,读了多少，则从buf中取多少
	fmt.Println("服务端回发信息成功，数据为:", string(buf[:n]))
}
