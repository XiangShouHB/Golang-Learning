//File  : demo01_chan_define.go
//Author: duanhaobin
//Date  : 2020/5/5

package main

import "fmt"

func main() {
	/*
		通道是什么：
			通道就是goroutine之间的通道。它可以让goroutine之间相互通信。

			每个通道都有与其相关的类型。该类型是通道允许传输的数据类型。
			通道的零值为nil。nil通道没有任何用处，因此通道必须使用类似于map和切片的方法来定义。
		通道的声明：
			声明一个通道和定义一个变量的语法一样：
			1.声明通道
				var 通道名 chan 数据类型

			2.创建通道：如果通道为nil(就是不存在)，就需要先创建通道
				通道名 = make(chan 数据类型)
		通道的数据类型：
			channel是引用类型的数据，在作为参数传递的时候，传递的是内存地址。
	*/
	var chInt chan int // 定义一个空通道
	if chInt == nil {
		fmt.Println("channel 是 nil 的, 不能使用，需要先创建通道。。")
		chInt = make(chan int)
		fmt.Printf("chInt 的数据类型为：%T，值为：%v\n", chInt, chInt)
	}
	// 简短声明，并初始化
	chStr := make(chan string)
	fmt.Println(chStr)

	// 给通道发送数据
	getChan(chStr) // 数值打印的是内存地址，说明通道是引用类型

}

func getChan(ch chan string) {
	fmt.Printf("ch 参数为的数据类型为：%T，数值为：%v\n", ch, ch)
}

func chanDemo() {
	var chanals [10]chan int
	for i := 0; i < 10; i++ {
		chanals[i] <- 'a' + i // 'a' 对应的ASCII值为97
	}
}
