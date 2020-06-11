//File  : demo05_struct_define.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

// 编程规范，结构体都是独立于 main 函数定义的
type Student struct {
	name  string
	age   int
	class string
}

func main() {
	/*
		结构体初识:
			概念：
				Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
				结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
				和Java Python中的class一样
			定义和初始化：
				语法定义：
					type struct_variable_type struct {
					   member definition;
					   member definition;
					   ...
					   member definition;
					}
				初始化：
					variable_name := structure_variable_type {value1, value2...valuen}

	*/

	// 1.定义,在main函数中定义结构体是支持的
	type Person struct {
		name  string
		age   int
		hight float64
	}

	// 2.初始化
	// 2.1 方法一：简短声明 初始化
	bob := Person{"Bob", 19, 1.85}
	fmt.Printf("Bob 数据类型：%T,值为：%v\n", bob, bob)

	// 2.2 方法二：var 定义
	var alan Person
	fmt.Println("alan 结构体", alan) // {  0 0}  注意，第一个值是空字符串，控制台输出不明显，看不出来
	alan.name = "Alan"
	alan.age = 20
	alan.hight = 1.78
	fmt.Println("Alan 结构体内容：", alan)

	// 2.3 方法三：Person{}
	tom := Person{}
	tom.name = "Tom"
	tom.age = 21
	tom.hight = 1.73
	fmt.Println("Tom 结构体内容：", tom)

	// 2.4 方法四：Person{}变体
	jack := Person{
		name:  "Jack",
		age:   19,
		hight: 1.69,
	}
	fmt.Println("Jack 结构体内容：", jack)

	// 3. 空结构体并不是 nil
	//var jerry Person
	//fmt.Printf(jerry == nil) //报错：invalid operation: alan == nil (mismatched types Person and nil)
	//var nil Type // Type must be a pointer, channel, func, interface, map, or slice type

}
