//File  : demo04_type_alias_struct.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import "fmt"

// 定义一个Person结构体
type Person struct {
	name string
}

// 定义接受者为 Person 的方法
func (p Person) ShowPerson() {
	fmt.Println("Person 的名字：", p.name)
}

// 给 Person 起别名
type People = Person

// 定义接受者为 Person 的方法
func (p People) ShowPeople() {
	fmt.Println("People 的名字：", p.name)
}

// 定义一个 Student 结构体，嵌套了两个结构体
type Student struct {
	Person // 匿名结构体
	People // 匿名结构体
}

func main() {
	/*
		当类型别名作为结构体嵌入的成员时会发生什么情况？

		会引起歧义，调用者不知道该访问 哪个结构体的属性或方法，因为别名结构体 和 原始的结构体本质上是一个类型
	*/
	// Person 结构体实例化
	bob := Person{name: "Bob"}
	fmt.Println("Bob -> ", bob.name)
	bob.ShowPerson()
	bob.ShowPeople() // 可以调用别名 People 的方法
	fmt.Printf("bob 的类型：%T，地址：%p\n", bob, &bob)

	// People 结构体实例化
	jack := People{name: "Jack"}
	fmt.Println("Jack -> ", jack.name)
	jack.ShowPerson()
	jack.ShowPeople()                              // 可以调用别名 People 的方法
	fmt.Printf("jack 的类型：%T，地址：%p\n", jack, &jack) // 其类型其实也是Peoson

	// Student 结构体实例化
	student := Student{
		Person: bob,
		People: jack,
	}
	// 以下两行代码会运行时报错，
	//fmt.Println(student.name) // ambiguous selector student.name
	//student.ShowPeople()      // ambiguous selector student.ShowPeople  含糊的选择者student.ShowPeople

	// 既然会引起歧义，那么消除歧义，指定访问具体的结构体，是不是就能解决上述问题？尝试一下
	fmt.Println("Student 结构体访问 Person 的属性，结果：", student.Person.name)
	fmt.Println("Student 结构体访问 People 的属性，结果：", student.People.name)

	// 同理，方法也可以调用
	student.Person.ShowPerson()
	student.People.ShowPeople()
	fmt.Println("=========反转调用的方法=======")
	student.Person.ShowPeople()
	student.People.ShowPerson()

	//
	fmt.Printf("student.Person 的类型：%T，地址：%p\n", student.Person, &student.Person)
}
