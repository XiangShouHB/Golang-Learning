//File  : socket.go
//Author: duanhaobin
//Date  : 2020/5/31

/*
	IP类型定义：
		type IP []byte
	ParseIP(s string) IP函数：
		把一个IPv4 或者IPv6的地址转化成IP类型
	os.Args：Args保留以程序名称开头的命令行参数。
*/
package main

import (
	"fmt"
	"net"
	"os"
)

// 执行之后，输入一个IP地址就会给出相应的IP格式
// ?? 执行后直接报错 TODO
func main() {
	args := os.Args
	fmt.Println("args:", args)
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s ip-addr\n", args[0])
		// 程序正常退出
		os.Exit(1)
	}
	name := args[1]
	// 把一个IPv4 或者IPv6的地址转化成IP类型
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String()) // 返回IP地址ip的字符串形式
	}
	// 程序正常退出
	os.Exit(0)
}
