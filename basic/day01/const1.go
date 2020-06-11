package main

import "fmt"

type numberConst int64

func main() {
	/*
		基本语法--常量constant
		常量是一个简单值的标识符，在程序运行时，不会被修改的量。
		1.数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
		2.满足多重赋值
		3.常量只定义不使用编译不会报错
		4.常量可以作为枚举，常量组
		5.常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
		6.显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值
	*/
	// 1.常量定义
	const monday int = 1 // 显式类型定义
	const tuesday = "2"  // 隐式类型定义，不使用 编译不会报错
	fmt.Printf("monday 的数据类型为：%T，值为：%d\n", monday, monday)
	fmt.Printf("tuesday 的数据类型为：%T，值为：%s\n", tuesday, tuesday)
	// 2.定义一组常量
	const MONDAY, TUESDAY, WEDNESDAY = 1, 2, 3
	//const (
	//	one   = 1
	//	two   = 2
	//	three = 3
	//)
	//fmt.Println("ONE的值为：", one)

	// 3.一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		FOUR = 4
		FIVE = 5
		FIVE_REPEAT
	)

	fmt.Printf("FIVE_REPEAT的类型为：%T,值为：%d\n", FIVE_REPEAT, FIVE_REPEAT)

	// 4.枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SUCESS    = 0
		ERROR     = 1
		EXCEPTION = 2
	)
	// 5.定义枚举
	const (
		one numberConst = 1 + iota
		two
		three
		four
		five
	)
	fmt.Printf("four 的数据类型为：%T, 值为：%d\n", four, four)
	fmt.Printf("five 的数据类型为：%T, 值为：%d\n", five, five)

}
