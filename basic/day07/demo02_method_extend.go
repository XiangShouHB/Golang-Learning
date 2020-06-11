//File  : demo02_method_extend.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import "fmt"

func main() {
	/*
		OOP中的继承性：
			如果两个类(class)存在继承关系，其中一个是子类，另一个作为父类，那么：

			1.子类可以直接访问父类的属性和方法
			2.子类可以新增自己的属性和方法
			3.子类可以重写父类的方法(orverride，就是将父类已有的方法，重新实现)


		Go语言的结构体嵌套：
			1.模拟继承性：is - a
				type A struct{
					field
				}
				type B struct{
					A //匿名字段
				}

			2.模拟聚合关系：has - a
				type C struct{
					field
				}
				type D struct{
					c C //聚合关系
				}

	*/

	// 1.访问父类的属性
	// 先初始化父类
	bob := Person{name: "Bob", age: 10, sex: "男"}
	fmt.Printf("bob 的姓名为：%s，年龄：%d，性别，%s\n", bob.name, bob.age, bob.sex)
	bob.printName() // 执行父类方法

	// 自定义 student 类
	jack := Student{
		Person: bob, // 直接同步父类的属性
		school: "实验小学",
	}
	// 子类对象，可以直接访问父类的字段，(其实就是提升字段)
	jack.name = "Jack" // 也可写为：jack.Person.name
	jack.age = 9
	jack.sex = "男"
	fmt.Printf("jack 的姓名为：%s，年龄：%d，性别，%s\n", jack.name, jack.age, jack.sex)

	jerry := Student{
		Person: Person{name: "Jerry", age: 11, sex: "男"}, // 使用父类的字段，自定义子类内容
		school: "实验二小",
	}
	fmt.Printf("jerry 的姓名为：%s，年龄：%d，性别，%s\n", jerry.name, jerry.age, jerry.sex)

	// 2.子类调用父类的方法
	jack.printName()         // 如果存在方法的重写，子类对象访问重写的方法
	jerry.printStudentInfo() // 子类对象访问自己新增的方法，存在继承关系时，按照就近原则，进行调用

	// 父类不能调用子类
	//bob.printStudentInfo()  报错，提示 printStudentInfo()不存在

}

type Person struct {
	name string
	age  int
	sex  string
}

func (person Person) printName() {
	fmt.Println("父类的方法，打印姓名：", person.name)
}

func (person Student) printName() {
	fmt.Println("父类的方法 printName2 ，打印姓名：", person.name)
}

type Student struct {
	Person // 匿名结构体，可以理解为 Person 就是 Student 的父类
	school string
}

func (stu Student) printStudentInfo() {
	fmt.Printf("子类的方法，学生的信息为：%s，年龄：%d，性别：%s\n", stu.name, stu.age, stu.sex)

}
