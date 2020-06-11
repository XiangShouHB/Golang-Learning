//File  : demo03_chan_close.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import "fmt"

func main() {
	/*
		关闭通道：close(ch)
			子goroutine：写出10个数据
				每写一个，阻塞一次，主goroutine读取一次，解除阻塞

			主goroutine，读取数据
				每次读取数据，阻塞一次，子goroutine，写出一个，解除阻塞
	*/
	// 创建通道
	ch := make(chan int)
	// 启动子 goroutine
	go sendData(ch)

	// 因为发送了多次，所以要接收多次。一般的开发过程，可能不知道具体发送了多少次
	// 裸for 来表示 while true
	for {
		v, ok := <-ch // ok表示通道中有数据就是 true
		if !ok {
			fmt.Printf("通道数据接收完毕，ok:%v,试着查看关闭通道后的零值：%d\n", ok, v)
			break
		}
		fmt.Println("读取的数据：", v)
	}
	fmt.Println("main...over...")

}

// 定义发送数据函数
func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 发送完数据要关掉通道
}
