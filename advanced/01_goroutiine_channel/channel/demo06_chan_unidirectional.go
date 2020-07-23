//File  : demo06_chan_unidirectional.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	/*
		双向：
			之前学习的通道都是双向通道，可读可写
			chan T
				chan <- data，发送数据，写出
				data <- chan，获取数据，读取

		单向：定向，只支持读写其中的一个操作
			chan <- T，只支持写
			<- chan T，只读
		注意:
			通常，单向通道都是定义在入参的时候，这样就能保护自己的内部数据
		调函数处理的时候，最好是定义双向通道来处理逻辑。

	*/
	// 定义一个双向通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 调用只写通道函数，写入数据
	go writeChan(ch1)
	data := <-ch1

	fmt.Println("writeChan获取的数据为：", data)

	//ch2 <- "只读" // 如果这行先写，那么会发生死锁，因为channel是和 goroutine相辅相成的
	//直接ch2 <- "只读"，写通道会陷入阻塞，没有对应 goroutine 去接收数据，所以一直堵塞，程序死锁

	// 调用只读通道函数，读取数据
	go readChan(ch2)
	// 给通道数据，保证能够读取
	ch2 <- "只读"
	fmt.Println("main...over...")
	log.Fatal(http.ListenAndServe(":6060", nil)) //http://localhost:6060/debug/pprof/  可以查看goroutine数量
}

// func(ch   chan<- int){}  参数为只发通道，就是只能向通道发送数据，不能从通道接收数据
// 定义writeChan函数，入参为 只写单向通道
func writeChan(wCh chan<- string) {
	wCh <- "只写"
	//resu := <-wCh  // Invalid operation: <-wCh (receive from send-only type chan<- string)
	fmt.Println("writeChan只写通道函数执行完毕......")
}

// 定义readChan函数，入参为 只读单向通道
func readChan(rCh <-chan string) {
	result := <-rCh
	fmt.Println("writeChan 只读通道函数，数据为：", result)
}
