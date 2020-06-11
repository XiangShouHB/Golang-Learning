//File  : client.go
//Author: duanhaobin
//Date  : 2020/5/29

/*
	模仿 nc 命令向服务器发送数据
	注意：
		1.Read()读取服务器/客户端关，如果返回0，表示对端关闭
		2.Linux\Mac系统回车是一个字符 '\n', 而Windows系统则是两个字符 '\r\n'
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 读取键盘输入(stdin)，将输入数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			// 读取键盘输入
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("键盘输入有误，err:", err)
				continue
			}

			// 将数据发送给服务器
			n, err = conn.Write(str[:n])
			if err != nil {
				fmt.Println("conn.Write，err:", err)
				return
			}

		}

	}()

	// 显示服务器回发的数据
	buf := make([]byte, 4096)
	for {
		// 读取服务器的数据
		n, err := conn.Read(buf)
		// 处理服务器关闭的场景
		if n == 0 {
			fmt.Println("检测到服务器已关闭，断开连接....")
			return
		}
		if err != nil {
			fmt.Println("conn.Read，err:", err)
			return
		}

		// 显示读取的数据
		fmt.Println("客户端读取服务器数据成功，数据：", string(buf[:n]))
	}

}
