package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup //创建同步等待组的对象
func main() {
	/*
		要想解决临界资源安全的问题，很多编程语言的解决方案都是同步。

		通过上锁的方式，某一时间段，只能允许一个goroutine来访问这个共享数据，当前goroutine访问完毕，
		解锁后，其他的goroutine才能来访问。

		我们可以借助于sync包下的锁操作。
		WaitGroup，同步等待组。

		在类型上，它是一个结构体。一个WaitGroup的用途是等待一个goroutine的集合执行完成。
		主goroutine调用了Add()方法来设置要等待的goroutine的数量。

		然后，每个goroutine都会执行并且执行完成后调用Done()这个方法。
		与此同时，可以使用Wait()方法来阻塞，直到所有的goroutine都执行完成。

		WaitGroup：同步等待组
			Add()，用来设置到WaitGroup的计数器的值。
			我们可以理解为每个waitgroup中都有一个计数器 用来表示这个同步等待组中要执行的goroutin的数量。

			如果计数器的数值变为0，那么就表示等待时被阻塞的goroutine都被释放，
			如果计数器的数值为负数，那么就会引发恐慌，程序就报错了

			Wait()，表示让当前的goroutine等待，进入阻塞状态。
			一直到WaitGroup的计数器为零。才能解除阻塞， 这个goroutine才能继续执行。

			Done()，同Add(-1)
			就是当WaitGroup同步等待组中的某个goroutine执行完毕后，设置这个WaitGroup的counter数值减1。

	*/
	wg.Add(2) //counter  3 2 1

	go fun1()
	go fun2()

	fmt.Println("main 进入阻塞状态。。等待wg中的子goroutien结束。。")
	wg.Wait() // 表示main goroutine进入等待，意味着阻塞
	fmt.Println("main..解除阻塞。。")

}

func fun1() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	for i := 1; i < 100; i++ {
		fmt.Println("fun1()函数中打印。。A ", i)
	}

	wg.Done() //给wg等待组中的counter数值减1。同	wg.Add(-1)
}

func fun2() {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	defer wg.Done()
	for j := 1; j < 100; j++ {
		fmt.Println("\tfun2()函数打印。。", j)
	}

}
