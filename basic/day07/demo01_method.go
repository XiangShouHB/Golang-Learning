//File  : demo01_method.go
//Author: duanhaobin
//Date  : 2020/4/30

package main

import "fmt"

func main() {
	/*
		方法：method
			一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
			方法只是一个函数，它带有一个特殊的接收器类型，它是在func关键字和方法名之间编写的。
			接收器可以是struct类型或非struct类型。接收方可以在方法内部访问。
			不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

		语法：
			func (接受者) 方法名(参数列表)(返回值列表){

			}
		注意事项：
		方法可以同名：
			1.虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
			2.method里面可以访问接收者的字段
			3.调用method通过.访问，就像struct里面访问字段一样

		总结：method，同函数类似，区别需要有接受者。(也就是调用者)

		对比函数：
			A：意义
				方法：某个类别的行为功能，需要指定的接受者调用
				函数：一段独立功能的代码，可以直接调用

			B：语法
				方法：方法名可以相同，只要接受者不同
				函数：命名不能冲突
	*/
	tom := Worker{
		name: "Tom",
		age:  20,
		sex:  "男",
	}
	// 1.调用方法
	tom.work()

	jack := Teacher{
		name: "Jack",
		age:  30,
		sex:  "男",
	}
	// 2.调用同名方法
	jack.work()

	// 3.不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型
	bob := Student{
		name: "Bob",
		age:  15,
		sex:  "男",
	}
	// study方法的接受者为指针类型，但是定义的 bob 是值类型结构体，依然可以调用 study方法
	bob.study()
}

// 1.定义结构体
type Worker struct {
	name string
	age  int
	sex  string
}

type Teacher struct {
	name string
	age  int
	sex  string
}

type Student struct {
	name string
	age  int
	sex  string
}

// 2. 定义方法
func (worker Worker) work() {
	fmt.Println(worker.name, "在工作.....")
}

// 可以定义相同的方法名，接受者不一致就表示是不同类型的方法
func (teacher Teacher) work() {
	fmt.Println(teacher.name, "在工作.....")

}

// 定义接受者为指针类型的方法
func (student *Student) study() {
	fmt.Printf("%s在学习....\n", (*student).name) // 可以省略 *
	fmt.Printf("%s在学习....\n", student.name)
}
