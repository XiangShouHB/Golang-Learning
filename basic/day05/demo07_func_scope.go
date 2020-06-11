//File  : demo07_func_scope.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import "fmt"

var global_str = "global 变量"
const GLOBAL_NUM = 100
func main() {
	/*
		作用域：变量可以使用的范围。
			局部变量：函数内部定义的变量，就叫做局部变量。
						变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。
						变量名可以同名
			全局变量：函数外部定义的变量，就叫做全局变量。
						所有的函数都可以使用，而且共享这一份数据，如果不希望修改全局变量，可以定义为常量：
						如: var const
	*/
	num1 := 5
	fmt.Println("函数 main 中，局部变量为：",num1)
	f1()
	f2()
	fmt.Println("全局变量为：",global_str) // hello
	fmt.Println("全局常量为：",GLOBAL_NUM)

}

// 局部变量
func f1()  {
	num1 := 1
	global_str = "hello"  // 更改全局变量
	fmt.Println("函数 f1 中，局部变量为：",num1)
	fmt.Println("函数 f1 中，全局变量为：",global_str)  // hello
}

func f2()  {
	num1 := 3  // 可以同名，作用域不同
	fmt.Println("函数 f12 中，局部变量为：",num1)
	fmt.Println("函数 f12 中，全局变量为：",global_str)  // hello
}