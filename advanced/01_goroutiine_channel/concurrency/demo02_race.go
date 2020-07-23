package main

import (
	"fmt"
	"time"
)

var a = 1

func main() {
	/*
		临界资源：
		指并发环境中多个进程/线程/协程共享的资源。

		但是在并发编程中对临界资源的处理不当， 往往会导致数据不一致的问题。
	*/
	go func() {
		a = 2
		fmt.Println("goroutine中。。", a)
	}()

	// 加锁也不管用
	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine...", a)

}
