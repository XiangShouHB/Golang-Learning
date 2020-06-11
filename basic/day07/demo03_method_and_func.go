//File  : demo03_method_and_func.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import "fmt"

func main() {
	/*
		既然我们已经有了函数，为什么还要使用方法？
		查看如下代码：

	*/
	bob := Student{
		name: "Bob",
		sex:  "男",
		age:  20,
	}
	// 调用函数，也输出了学生信息，和 demo02_method_extend.go 中，Student结构体的方法功能一致
	printStudentInfo(bob)

	/*
		那为什么我们可以用函数来写相同的程序呢?有以下几个原因

		1.Go不是一种纯粹面向对象的编程语言，它不支持类。因此，类型的方法是一种实现类似于类的行为的方法。

		2.相同名称的方法可以在不同的类型上定义，而具有相同名称的函数是不允许的。假设我们有一个正方形和圆形的结构。
		可以在正方形和圆形上定义一个名为Area的方法。这是在下面的程序中完成的。
	*/

}

type Student struct {
	name, sex string
	age       int
}

func printStudentInfo(stu Student) {
	fmt.Printf("学生的信息为：%s，年龄：%d，性别：%s", stu.name, stu.age, stu.sex)
}
