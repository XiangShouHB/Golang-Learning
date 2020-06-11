//File  : demo07_struct_anonymous.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段：

			匿名结构体：没有名字的结构体，
				在创建匿名结构体时，同时创建对象
				变量名 := struct{
					定义字段Field
				}{
					字段进行赋值
				}

			匿名字段：一个结构体的字段没有字段名

	*/
	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)

	// 创建一个匿名函数
	func() {
		fmt.Println("hello world...")
	}()

	// 创建一个匿名结构体，在创建时就会初始化，一般作用不大。了解这个语法即可
	s2 := struct {
		name string
		age  int
	}{}
	fmt.Println(s2.name, s2.age)

	bob := Worker{"Bob", 22}
	fmt.Println(bob)
	fmt.Println(bob.string)
	fmt.Println(bob.int)
}

type Worker struct {
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字，那么匿名字段的类型就不能重复，否则会冲突
	//string
}

type Student struct {
	name string
	age  int
}
