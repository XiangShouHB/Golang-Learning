package main

import (
	"fmt"
	"time"
)

/*
	select 是 Go 中的一个控制结构。

	select 语句类似于 switch 语句，但是 select 会随机执行一个可运行的 case。
	如果没有 case 可运行，它将阻塞，直到有 case 可运行。
	select 语句的语法结构和 switch 语句很相似，也有 case 语句和 default 语句

	注意：
		nil channel在select语句中是可以正确运行的，不会报错，但是不会被case到的，也就是说永辉被阻塞到case处

	本文件用到的知识汇总：
		1. select的使用
		2. 三种定时器的使用
		3. 在select中使用 nil channel，小技巧
*/

// doWork 执行worker
func doWork(id int, ch chan int) {
	for i := range ch {
		// 每隔1s 从通道中读取一个数
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, i)
	}
}

// createWorker 并发创建worker
func createWorker(id int) chan int {
	out := make(chan int)
	go doWork(id, out)
	return out
}

// genChannel 生成int类型的channel
func genChannel() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 每隔500ms 发送一个数据
			time.Sleep(time.Millisecond * 500)
			out <- i
			i++
		}
	}()
	return out
}

// selectDemo01 select初步使用
func selectDemo01() {
	var c1, c2 = genChannel(), genChannel()
	// 使用select读取channel中的数据
	for {
		select {
		case <-c1:
			fmt.Printf("Received %d from c1\n", <-c1)
		case <-c2:
			fmt.Printf("Received %d from c2\n", <-c2)
		default:
			fmt.Println("No value received......")
		}
	}
}

// selectDemo02 select进阶使用
func selectDemo02() {
	var (
		c1, c2 = genChannel(), genChannel()
		worker = createWorker(0)
		// 用来保存数据，保证从c1 c2读取的值是连续的，不会断开(并发程序过快，会导致前面已经读取的数据无法打印出来)
		values []int
	)
	// 定时器使用1：10s之后结束select。从该代码执行开始之后的10s
	tm := time.After(time.Second * 10)
	// 定时器使用2：查看5s内数据积压量
	// 因为 genChannel() 和 doWork() 的时间不同步，大致为发送数据要比接收数据块1倍，因为会有数据积压到values中
	tick := time.Tick(2 * time.Second)
	for {
		var (
			activeWorker chan int
			activeValue  int
		)
		// values 有值才初始化 activeWorker，并且activeValue取 values 的第一个元素
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)

		case n := <-c2:
			values = append(values, n)

		// 利用select不用case channel 为 nil的特性
		case activeWorker <- activeValue:
			// 因为activeValue取了 values 第一个元素的值，所以当 activeWorker 接收到值后，要去除已经接收了的值
			values = values[1:]

		// 定时器的使用3：select中400ms内没有case成功执行(即接收数据)，则输出超时内容
		case <-time.After(400 * time.Millisecond):
			fmt.Println("time out.......")

		// 查看 values 中积压了多少数据，因为 genChannel 中sleep了1000ms
		case <-tick:
			fmt.Println("lengh of values : ", len(values))

		// 10s之后本函数退出
		case <-tm:
			fmt.Println("10 seconds have come...... bye")
			return
		}
	}

}

func main() {
	//selectDemo01()
	selectDemo02()
}
