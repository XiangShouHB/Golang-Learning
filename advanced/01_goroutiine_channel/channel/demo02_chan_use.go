//File  : demo02_chan_use.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*
		通道的使用语法:
		1.发送和接收:
			data := <- a // read from channel a
			a <- data // write to channel a
		在通道上箭头的方向指定数据是发送还是接收。

		 2.发送和接收默认是阻塞的:
			一个通道发送和接收数据，默认是阻塞的。
			当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个Goroutine从该通道读取数据。
			相对地，当从通道读取数据时，读取被阻塞，直到一个Goroutine将数据写入该通道。
			这些通道的特性是帮助Goroutines有效地进行通信，而无需像使用其他编程语言中非常常见的显式锁或条件变量。

		3.一次发送对应一个获取
	*/
	// 创建一个通道
	chStr := make(chan string)
	// 创建一个 goroutine
	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(3 * time.Second)

			chStr <- ("数据发送到通道了" + strconv.Itoa(i)) // 在循环内将数据发送到通道,输出内容如以下注释：
			fmt.Printf("开启一个goroutine, i---->%d\n", i)
			/*
				开启一个goroutine, i---->1
				从通道接收的数据为： 数据发送到通道了
				main....over.....

				解释：一次发送对应一个获取
			*/

		}
		// 数据发送到通道
		//chStr <- "数据发送到通道了" // 在循环外将数据发送到通道，是正常循环完，按执行顺序输出 正常 的内容
		fmt.Println("子 goroutine over.....")
	}()

	time.Sleep(5 * time.Second)

	// 从通道接收数据

	resStr := <-chStr
	fmt.Println("从通道接收的数据为：", resStr)
	fmt.Println("main....over.....")

}
