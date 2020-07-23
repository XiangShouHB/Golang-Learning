//File  : demo04_chan_range.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import "fmt"

func main() {
	/*
		通道上的范围循环:
			遍历通道，除了用裸for获取通道数据外，Golang提供了一个更优雅，更遍历的循环:

			for range,该形式可用于从通道接收值，直到它关闭为止
	*/
	ch := make(chan int)

	go sendData(ch)

	for v := range ch {
		fmt.Printf("读取的数据为:%d\n", v)
	}
	fmt.Println("main...over...")
}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

}
