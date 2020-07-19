//File  : demo04_interface_define.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import "fmt"

func main() {
	/*
		接口：interface
			在Go中，接口是一组方法签名。

			它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

			接口定义了一组方法，如果某个结构体实现了某个接口的所有方法，则此结构体就实现了该接口。

			Go语言中，接口和类型的实现关系，是非嵌入式的

			其他语言中，要显示的定义:
			class class_name implements interface_name{}

			1.当需要接口类型的对象时，可以使用任意实现类对象代替
			2.接口对象不能访问实现结构体中的属性

			interface可以被任意的对象实现
			一个对象可以实现任意多个interface
			任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface


		多态：一个事物的多种形态
			go语言通过接口模拟多态

			就一个接口的实现
				1.看成实现本身的类型，能够访问实现结构体中的属性和方法
				2.看成是对应的接口类型，那就只能够访问接口中的方法

		接口的用法：
			1.一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数。
			2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象
	*/
	// 1.创建结构体
	mate30 := Huawei{
		name:  "Mate 30",
		price: 6999,
	}
	mate30.call()
	mate30.seenMessage()

	xiaomi9 := Xiaomi{
		name:  "Xiao 9",
		price: 4999,
	}

	xiaomi9.seenMessage()

	// 2.将结构体赋值给接口
	var phone Phone
	phone = mate30
	phone.call()

	//phone.name  // 不能访问结构体的属性。 可以理解为：创建了 Phone 这个接口，里面有打电话和发短信功能。
	//谁打的电话？名字是什么？ Phone 不在乎，它只负责提供这些功能。
	//然后具体怎么实现打电话信号更好，发短信更方便，Phone 是不关心的，只有实现这个接口的结构体会关心

	//  3.判断结构体是否实现了接口 Golang中使用 new() 可以判断
	var _ Phone = new(Huawei)
	var _ Phone = (*Huawei)(nil) // new(Huawei)写是编译的时候检查，这样写是运行的时候检查

	//var _ Phone = new(Xiaomi) // 这行代码编译会报错，因 为Xiaomi 没有实现 Phone 接口的所有方法，报错内容如下
	// Cannot use 'new(Xiaomi)' (type *Xiaomi) as type Phone Type does not implement 'Phone' as some methods are missing: call()

}

// 1.声明一个接口，包含两个方法
type Phone interface {
	call()
	seenMessage()
}

// 2.声明结构体，用来实现接口

// Huawei 实现了 Phone的 所有方法，所以可以任务 Huawei 类型实现了 Phone 接口
type Huawei struct {
	name  string
	price float64
}

type Xiaomi struct {
	name  string
	price float64
}

func (huawei Huawei) call() {
	fmt.Printf("%s 有打电话功能.....\n", huawei.name)
}

func (huawei Huawei) seenMessage() {
	fmt.Printf("%s 有发短信功能.....\n", huawei.name)
}

// Xiaomi 只实现了 Phone的 发短信方法，所以认为 Xiaomi 类型未实现了 Phone 接口
// 知识点是这样讲概念的，但是在代码中，是怎么判断一个结构体是否实现了某个接口呢？
func (xiaomi Xiaomi) seenMessage() {
	fmt.Printf("%s 只有发短信功能.....\n", xiaomi.name)
}
