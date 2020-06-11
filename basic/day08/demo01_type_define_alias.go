//File  : demo01_type_define_alias.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import "fmt"

// 自定义类型
type myInt int // 一种新的类型

type myString string

// 类型别名
type intAlias = int // 不是重新定义新的数据类型，只是给int起别名，和int可以通用

func main() {
	/*
		type：用于类型定义和类型别名
			1.类型定义：type 类型名 Type
			2.类型别名：type  类型名 = Type
		实践：
			1.使用type 可以定义结构体类型，之前学过，可查看之前笔记
			2.使用type 可以定义接口类型，之前学过，可查看之前笔记
			3.定义其他的新类型
				type 类型名 Type
			4.类型别名
				type 类型别名 = Type
	*/

	// 定义内置的数据类型
	var num = 100
	var str = "Hello"
	fmt.Printf("num 的数据类型为：%T，值为；%d\n", num, num)
	fmt.Printf("str 的数据类型为：%T，值为；%s\n", str, str)

	// 定义自定义数据类型
	var myNum myInt = 200
	var myStr myString = "World"
	fmt.Printf("myNum 的数据类型为：%T，值为；%d\n", myNum, myNum)
	fmt.Printf("myStr 的数据类型为：%T，值为；%s\n", myStr, myStr)

	// 起别名
	var myNumAlias intAlias = 300
	fmt.Printf("myNum32 的数据类型为：%T，值为；%d\n", myNumAlias, myNumAlias)

	// 不同类型之间的类型不能相互赋值
	//myNum = num // cannot use num (type int) as type myInt in assignment

	// 别名之间赋值是不影响的
	num = myNumAlias
	fmt.Println("num 重新赋值后为：", num)

}
