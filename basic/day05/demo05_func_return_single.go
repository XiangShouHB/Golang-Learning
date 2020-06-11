//File  : demo05_func_return_single.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import "fmt"

func main() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处。执行结果就叫做函数的返回值。

		return语句：
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处。
			函数返回的结果，必须和函数定义的一致：类型，个数，顺序。

			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行

		空白标识符，专门用于舍弃数据：_
	*/
	str := getString()
	fmt.Println("getString() 返回结果为：",str)

	num := getInt()
	fmt.Println("getInt() 返回结果为：",num)  // 返回默认值0



}

// 指定返回值类型，但是未定义返回值名称，return语句要写清楚返回的变量
// 单个返回值: string 也可以写为 (string)
func getString() string {
	fmt.Println("进入 getString()......")
	s := "Hello world"
	return s
}

// 返回值名称 类型 都定义了， return 可以裸写，默认返回定义的返回值
func getInt() (num int64){
	return
}

