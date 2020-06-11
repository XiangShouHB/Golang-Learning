//File  : demo08_func_defer.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import "fmt"

func main() {
	/*
		defer的词义："延迟","推迟"
		在go语言中，使用defer关键字来延迟一个函数或者方法的执行。

		1.defer函数或方法：一个函数或方法的执行被延迟了。

		2.defer的用法：
			A：对象.close(),临时文件的删除。。。
					文件.open()
					defer close()
					读或写

			B：go语言中关于异常的处理，使用panic()和recover()
				panic函数用于引发恐慌，导致程序中断执行
				recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行。


		3.如果多个defer函数:先延迟的后执行，后延迟的先执行。栈的数据结构特点

		4.defer函数传递参数的时候：defer函数调用时，就已经传递了参数数据了，只是暂时不执行函数中的代码而已。

		5.defer函数注意点：
			defer函数：
			当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
			当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
			当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
	*/
	// 1.defer函数 延迟执行
	//a := 1
	//b := 2
	//defer fmt.Println(b)
	//fmt.Println(a)

	// 2. 延迟参数，延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行
	fmt.Println("....................2.....................")
	str := "普通字符串"
	fmt.Println(str)
	defer func1(str)  // func1函数被延迟执行，在main函数执行完后，才执行func1函数，参数在执行 defer 语句时就已经执行了
	str = "改变后的字符串"
	fmt.Println(str)


	// 3.当一个函数有多个 defer 调用时，它们被添加到一个堆栈中，并在Last In First Out（LIFO）后进先出的顺序中执行
	name := "Alice"
	fmt.Println("正常的名字：",name)
	fmt.Println("反转的名字：")
	for _,v := range name{
		defer fmt.Printf("%c",v)
	}
}

func func1(str string)  {
	fmt.Println("func1函数执行了，入参为:",str)
}

func func2(str string)  {
	fmt.Println("func2函数执行了，入参为:",str)
}