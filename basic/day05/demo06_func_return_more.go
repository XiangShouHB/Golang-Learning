//File  : demo06_func_return_more.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import "fmt"

func main() {
	/*
		多个返回值：
			1.多个返回值类型都相同，可以在最后一个返回值后指定数据类型。
				如：func getInt() (num1,num2,num3 int){}
			2.多个返回值类型只有一部分相同，那么相同类型的返回值可以指定数据类型，不同的数据类型要单独指定
				如：func getContent() (num1,num2, int, str string)

		return语句：词义"返回"
			A：一个函数有返回值，那么使用return将返回值返回给调用处
			B：同时意味着结束了函数的执行

		注意点：
			1.一个函数定义了返回值，必须使用return语句将结果返回给调用处。return后的数据必须和函数定义的一致：个数，类型，顺序。
			2.可以使用_,来舍弃多余的返回值
			3.如果一个函数定义了有返回值，那么函数中有分支，循环，那么要保证，无论执行了哪个分支，都要有return语句被执行到
			4.如果一个函数没有定义返回值，那么函数中也可以使用return，专门用于结束函数的执行。。
	*/
	a, b, c := getInt()
	fmt.Println("getInt() 函数调用，返回值：", a, b, c)

	num1, _, str := getContent() // _ 舍弃其中一个返回值
	fmt.Println("getContent() 函数调用，返回值：", num1, str)

}

// 多个返回值类型都相同,return 不显示写返回值列表，则默认全部返回
func getInt() (num1, num2, num3 int) {
	//return 1, 2,3
	return
}

// 多个返回值类型只有一部分相同
func getContent() (num1, num2 int, str string) {
	return 1, 2, "Hello"
}
