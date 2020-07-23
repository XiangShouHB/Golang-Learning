package main

import (
	"fmt"
	"sync"
)

/*
	使用channel来等待goroutine结束
	知识点：
		1.channel 的收发及关闭操作
		2.waitGroup的使用
*/

// Worker 结构体
type Worker struct {
	in   chan int // int类型通道
	done func()   // waitGroup的done操作
}

// doWork 执行worker
func doWork(id int, w Worker) {
	for i := range w.in {
		fmt.Printf("Worker %d received %c\n", id, i)
		w.done()
	}
}

// createWorker 并发创建worker
func createWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

// handleChan 向通道发送消息
func handleChan() {
	var wg sync.WaitGroup
	var workers [10]Worker
	// 创建10个worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	// 添加20个任务
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//wg.Add(1) // 当然也可以在for中，循环一次添加一个任务
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//wg.Add(1)
	}
	wg.Wait()
}
func main() {
	handleChan()
}
