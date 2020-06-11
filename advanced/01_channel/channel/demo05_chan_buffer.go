//File  : demo05_chan_buffer.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import "fmt"

func main() {
	/*
		非缓冲通道：make(chan T)
			一次发送，一次接收，都是阻塞的
		缓冲通道：make(chan T , capacity)
			发送：缓冲区的数据满了，才会阻塞
			接收：缓冲区的数据空了，才会阻塞
	*/

	// 定义一个缓冲个数为5的通道
	ch_buffer := make(chan int, 5)
	fmt.Printf("ch_buffer 长度为%d，容量为：%d\n", len(ch_buffer), cap(ch_buffer))

	// 如果缓冲区满了，再往进写数据会 panic,死锁
	ch_buffer <- 1
	ch_buffer <- 2
	ch_buffer <- 3
	ch_buffer <- 4
	ch_buffer <- 5
	fmt.Printf("ch_buffer 长度为%d，容量为：%d\n", len(ch_buffer), cap(ch_buffer))
	//ch_buffer <- 6  //all goroutines are asleep - deadlock!

	// 开启goroutine
	go sendData(ch_buffer)
	// 读取数据
	for v := range ch_buffer {
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main...over...")

}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据 i:", i)
	}

	close(ch)
}
