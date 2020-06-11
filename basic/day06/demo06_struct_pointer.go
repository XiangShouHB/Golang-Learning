//File  : demo06_struct_pointer.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}
type Student struct {
}

func main() {
	/*
			数据类型：
				值类型：int，float，bool，string，array，struct  注意，结构体是值类型

				引用类型：slice，map，function，pointer

			通过指针：
				new()，不是nil，空指针
					指向了新分配的类型的内存空间，里面存储的零值。
					new()创建的内容都为指针
			make、new操作:

				make用于内建类型（map、slice 和channel）的内存分配。

		        new用于各种类型的内存分配 内建函数new本质上说跟其它语言中的同名函数功能一样：
				new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
				用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：new返回指针

				内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，
				并且返回一个有初始值(非零)的T类型，而不是*T。

				本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
				例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；
				在这些项目被初始化之前，slice为nil。
				对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

				make返回初始化后的（非零）值。
	*/
	// 1.结构体是值类型
	bob := Person{
		name: "Bob",
		age:  19,
		sex:  "男",
	}
	fmt.Println("bob 结构体：", bob)
	fmt.Printf("bob 结构体地址为：%p,类型为：%T\n", &bob, bob)

	bob_bro := bob
	fmt.Println("bob_bro 结构体：", bob_bro)
	fmt.Printf("bob_bro 结构体地址为：%p,类型为：%T\n", &bob_bro, bob_bro)

	// 如果修改了 bob_bro 的name, bob的name会不会改变
	bob_bro.name = "Jack"
	fmt.Println("bob_bro 结构体修改name后：", bob_bro)
	fmt.Println("bob 结构体未发生变化，name还是 bob：", bob) // name还是bob，因为 struct 是值类型

	// 2.结构体指针  jerryPtr 是一个指针，指向结构体
	var jerryPtr *Person
	jerryPtr = &bob
	fmt.Println("jerryPtr 结构体指针为：", jerryPtr)
	fmt.Printf("jerryPtr 结构体指针地址为：%p,类型为：%T\n", &jerryPtr, jerryPtr)

	// 如果修改了 jerryPtr 的name, bob的name会不会改变
	jerryPtr.name = "Jerry"
	fmt.Println("jerryPtr 结构体修改name后：", jerryPtr)
	fmt.Println("bob 结构体发生变化，name变为 Jerry：", bob) // name改为Jerry,因为 structPtrJerry为结构体指针，引用传递

	// 3.内置函数 new()
	alan := new(Person)
	fmt.Println("alan 为：", alan) //不是nill

	fmt.Println("alan == nill :", alan == nil)

	fmt.Printf("alan 的地址为：%p,数据类型为；%T\n", alan, alan)
	alan.name = "Alan"
	alan.age = 19
	alan.sex = "男"
	fmt.Println("alan 结构体的内容为：", alan)

	// 4.new() 函数   func new(Type) *Type 返回的是指针
	int_ptr := new(int)
	fmt.Printf("int_ptr 的数据类型为：%T\n", int_ptr) // *int 一个指针，指向int
	fmt.Println(*int_ptr)                      // 取 int_ptr 指针的值，默认为0

	// 5.定义一个空结构体 Student 没有任何成员变量
	student := new(Student) // 即使没有成员变量，通过 new() 创建结构体也不是nil
	fmt.Printf("student 的数据类型为:%T,值为：%v\n", student, student)
	fmt.Println("student == nill :", student == nil)

	// 6. nil 是Golang中的一种类型
	var s []int64
	fmt.Println("s :", s)
	fmt.Println("s == nil：", s == nil)

}
