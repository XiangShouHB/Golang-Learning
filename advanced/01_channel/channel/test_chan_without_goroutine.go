//File  : test_chan_without_goroutine.go
//Author: duanhaobin
//Date  : 2020/5/6

package main

import "fmt"

func main() {
	ch := make(chan int)
	//go readFunc(ch)  创建个Go程，从通道中读取出数据就不会死锁了

	ch <- 1
	/*
			报错信息；
			fatal error: all goroutines are asleep - deadlock!

			goroutine 1 [chan send]:
			main.main()
		        F:/WorkSpaces/go/src/Golang-Learning/advanced/day02/channel/test_chan_wi
			thout_goroutine.go:11 +0x5c

			在 ch <- 1这行代码处程序发生死锁，因为代码执行后，向通道ch发送了信息，但是没有gotoutine从通道ch接收信息

			因此一直堵塞在 chan send 处，程序发生死锁

	*/
	go readFunc(ch)
}
func readFunc(ch chan int) {
	result := <-ch
	fmt.Println(result)
}
