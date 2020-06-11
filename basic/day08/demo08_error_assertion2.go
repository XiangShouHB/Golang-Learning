//File  : demo08_error_assertion2.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"fmt"
	"net"
)

func main() {
	/*
		2.断言底层类型，并通过调用struct类型的 方法 获取更多信息。
		举个例子，查看 DNSError 源码：
		// DNSError represents a DNS lookup error.
		type DNSError struct {
			Err         string // description of the error
			Name        string // name looked for
			Server      string // server used
			IsTimeout   bool   // if true, timed out; not all timeouts set this
			IsTemporary bool   // if true, error is temporary; not all errors set this
			IsNotFound  bool   // if true, host could not be found
		}
		// Error() 源码
		func (e *DNSError) Error() string {
			if e == nil {
				return "<nil>"
			}
			s := "lookup " + e.Name
			if e.Server != "" {
				s += " on " + e.Server
			}
			s += ": " + e.Err
			return s
		}
		// DNSError 结构体实现了 Timeout()  Temporary 方法
		func (e *DNSError) Timeout() bool { return e.IsTimeout }

		func (e *DNSError) Temporary() bool { return e.IsTimeout || e.IsTemporary }


		从上面的代码中可以看到，DNSError struct 有两个方法Timeout() bool和Temporary() bool。
		它们返回一个布尔值，表示错误是由于超时还是临时的。

		让我们编写一个断言*DNSError类型的程序，并调用这些方法来确定错误是临时的还是超时的。
	*/

	addrs, err := net.LookupHost("www.bucunzaide.com")

	if err != nil {
		if ins, ok := err.(*net.DNSError); ok {
			if ins.IsTimeout {
				fmt.Println("链接超时......")
			} else if ins.IsTemporary {
				fmt.Println("暂时性错误......")
			} else if ins.IsNotFound {
				fmt.Printf("链接无法找到......,err:%v\n", err)
			} else {
				fmt.Println("未知错误......", err)
			}
		}
		return
	}
	fmt.Println("访问成功，地址为：", addrs)
}
