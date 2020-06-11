//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/29
/*
	简单版的C/S程序，进行一次通信后就会断开连接，程序也随之退出。显然这种场景并不符合人们正常的通信，通常都是在不断发生和接收消息

	并发服务器：
		1.每个 go 协程负责一个socket通信
		2.随时监听客户端
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

// 与客户端交互
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	// 获取客户端地址信息
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接!")
	// 开始交互,循环读取客户端发送的请求
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		// 加 '\n' 的原因是，nc 编写完命令后，会回车来执行命令
		// 加 '\r\n' 的原因是，windows 编写完命令后，回车执行命令是两个字符，和Linux,Mac系统不同

		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Printf("服务器检测到客户端执行关闭命令[%s]，断开连接....\n", string(buf[:n]))
			return
		}
		// 处理客户端已关闭的场景,用于主动关闭，如ctrl+c,会点右上角关闭按钮
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭，断开连接....")
			return
		}
		// 处理 err
		if err != nil {
			fmt.Println("服务端读取数据异常，err:", err)
			return
		}
		// 打印读取到的数据
		fmt.Println("服务端读取到的数据：", string(buf[:n]))

		// 简单处理数据，回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	// 创建监听 socket
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		go HandleConnect(conn)
	}
}
