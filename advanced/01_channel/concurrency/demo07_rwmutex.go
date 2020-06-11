package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	/*
		RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能。

		RWMutex是读/写互斥锁。锁可以由任意数量的读取器或单个编写器持有。RWMutex的零值是未锁定的mutex。

		如果一个goroutine持有一个rRWMutex进行读取，而另一个goroutine可能调用lock，
		那么在释放初始读取锁之前，任何goroutine都不应该期望能够获取读取锁。

		特别是，这禁止递归读取锁定。这是为了确保锁最终可用；被阻止的锁调用会将新的读卡器排除在获取锁之外。

		我们怎么理解读写锁呢？当有一个 goroutine 获得'写'锁定，其它无论是读锁定还是写锁定都将阻塞直到'写'解锁；
		当有一个 goroutine 获得'读'锁定，其它'读'锁定仍然可以继续；
		当有一个或任意多个'读'锁定，'写'锁定将等待所有'读'锁定解锁之后才能够进行'写'锁定。
		所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，
		你给我站一边去，等它们读（读解锁）完你再来写（写锁定）。

		我们可以将其总结为如下三条：
			1.同时只能有一个 goroutine 能够获得写锁定。
			2.同时可以有任意多个 gorouinte 获得读锁定。
			3.同时只能存在写锁定或读锁定（读和写互斥）。

		所以，RWMutex这个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景。

		读写锁的写锁只能锁定一次，解锁前不能多次锁定，读锁可以多次，但读解锁次数最多只能比读锁次数多一次，一般情况下我们不建议读解锁次数多余读锁次数。

		基本遵循两大原则：
			1.可以随便读，多个goroutine同时读。
			2.写的时候，啥也不能干。不能读也不能写。

		读写锁即是针对于读写操作的互斥锁。它与普通的互斥锁最大的不同就是，它可以分别针对读操作和写操作进行锁定和解锁操作。

		读写锁遵循的访问控制规则与互斥锁有所不同。
		在读写锁管辖的范围内，它允许任意个读操作的同时进行。但是在同一时刻，它只允许有一个写操作在进行。

		并且在某一个写操作被进行的过程中，读操作的进行也是不被允许的。
		也就是说读写锁控制下的多个写操作之间都是互斥的，并且写操作与读操作之间也都是互斥的。

		但是，多个读操作之间却不存在互斥关系。

		最后概括：

		1.读锁不能阻塞读锁
		2.读锁需要阻塞写锁，直到所有读锁都释放
		3.写锁需要阻塞读锁，直到所有写锁都释放
		4.写锁需要阻塞写锁

	*/
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	//wg.Add(2)
	//
	////多个同时读取
	//go readData(1)
	//go readData(2)

	wg.Add(4)
	go writeData(1)
	go readData(2)
	go readData(3)
	go writeData(4)

	wg.Wait()
	fmt.Println("main..over...")
}

func writeData(i int) {
	defer wg.Done()

	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "开始写：write start。。")
	fmt.Println(i, "正在写：writing。。。。")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()

	fmt.Println(i, "写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()

	rwMutex.RLock() //读操作上锁
	fmt.Println(i, "开始读：read start。。")
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() //读操作解锁

	fmt.Println(i, "读结束：read over。。。")
}
