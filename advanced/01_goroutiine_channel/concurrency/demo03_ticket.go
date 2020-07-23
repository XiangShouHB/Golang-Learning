package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 全局变量，表示票
var ticket = 10 // 10张票
func main() {
	/*
		 临界资源安全问题：
			并发本身并不复杂，但是因为有了资源竞争的问题，就使得我们开发出好的并发程序变得复杂起来，因为会引起很多莫名其妙的问题。

			如果多个goroutine在访问同一个数据资源的时候，其中一个线程修改了数据，那么这个数值就被修改了，对于其他的goroutine来讲，这个数值可能是不对的。
			4个goroutine，模拟4个售票口，
	*/
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(5 * time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			// 我们为了更好的观察临界资源问题，每个goroutine先睡眠一个随机数，然后再售票，
			// 我们发现程序的运行结果，还可以卖出编号为负数的票。
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
	}
}
