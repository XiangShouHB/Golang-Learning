//File  : demo06_interface_nesting.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import (
	"fmt"
)

func main() {
	/*
		接口的嵌套：
	*/
	var person Person = Person{}
	// 实现 C 接口的所有方法
	person.test1()
	person.test2()
	person.test3()

	fmt.Println("-----实现 A 接口的方法:---------")

	// 实现 A 接口的方法
	var a1 A = Person{}
	a1.test1()

	fmt.Println("-------实现 B 接口的方法:-------")
	// 实现 B 接口的方法
	var b1 B = Person{}
	b1.test2()

	fmt.Println("-----实现 C 接口的方法:---------")
	// 实现 C 接口的方法
	var c1 C = Person{}
	c1.test1()
	c1.test2()
	c1.test3()

	fmt.Println("-----测试:---------")
	//var c2 C = a1 // 编译报错,如下：
	/*
		Cannot use 'a1' (type A) as type C Type does not implement 'C' as some methods are missing: test2() test3()
	*/
	var a2 A = c1
	a2.test1()
	/*
		可以理解为，大的可以容下小的，但是小的装不下大的。符合常识
		c2 为 C 接口类型,不仅嵌套了A B 接口，还有自己的方法，而 a1 只是 A 接口，所以var c2 C = a1会报错
	*/
}

// 定义3个接口
type A interface {
	test1()
}

type B interface {
	test2()
}

// 定义嵌套接口
type C interface {
	A
	B
	test3()
}

type Person struct {
	//如果想实现接口C，那不止要实现接口C的方法，还要实现接口A，B中的方法
}

func (p Person) test1() {
	fmt.Println("test1 方法................")
}

func (p Person) test2() {
	fmt.Println("test2 方法................")
}

func (p Person) test3() {
	fmt.Println("test3 方法................")
}
