package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量，表示票
var ticket = 10 //100张票

var mutex sync.Mutex //创建锁头

var wg sync.WaitGroup //同步等待组对象
func main() {
	/*
		在主流的编程语言中为了保证多线程之间共享数据安全性和一致性，都会提供一套基本的同步工具集，如锁，条件变量，原子操作等等。
		Go语言标准库也毫不意外的提供了这些同步机制，使用方式也和其他语言也差不多。

		sync是synchronization同步这个词的缩写，所以也会叫做同步包。
		这里提供了基本同步的操作，比如互斥锁等等。这里除了Once和WaitGroup类型之外，大多数类型都是供低级库例程使用的。

		更高级别的同步最好通过channel通道和communication通信来完成

		Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。
		Mutex 是最简单的一种锁类型，互斥锁，同时也比较暴力，
		当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex。

		每个资源都对应于一个可称为 “互斥锁” 的标记，这个标记用来保证在任意时刻，只能有一个协程（线程）访问该资源。其它的协程只能等待。

		互斥锁是传统并发编程对共享资源进行访问控制的主要手段，它由标准库sync中的Mutex结构体类型表示。
		sync.Mutex类型只有两个公开的指针方法，Lock和Unlock。Lock锁定当前的共享资源，Unlock进行解锁。

		在使用互斥锁时，一定要注意：对资源操作完成后，一定要解锁，否则会出现流程执行异常，死锁等问题。
		通常借助defer。锁定后，立即使用defer语句保证互斥锁及时解锁。

		Lock()方法：锁定m。如果该锁已在使用中，则调用goroutine将阻塞，直到互斥体可用。

		Unlock()方法，解锁解锁m。如果m未在要解锁的条目上锁定，则为运行时错误。

		锁定的互斥体不与特定的goroutine关联。允许一个goroutine锁定互斥体，然后安排另一个goroutine解锁互斥体。

		4个goroutine，模拟4个售票口，


	*/

	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait() //main要等待
	fmt.Println("程序结束了。。。")

	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		//上锁
		mutex.Lock()    //g2
		if ticket > 0 { //ticket 1 g1
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket) // 1
			ticket--                         // 0
		} else {
			mutex.Unlock() //条件不满足，也要解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		mutex.Unlock() //解锁
	}
}
