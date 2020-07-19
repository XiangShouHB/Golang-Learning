//File  : demo07_interface_assert.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		1. Type Assertion
			类型断言，通常有两种方式
			第一种：
				t := i.(T)
			这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，
			如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic。
			第二种：
				t, ok:= i.(T)
			这个表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T;

			如果断言成功，就会返回其类型给 t，并且此时 ok 的值 为 true，表示断言成功。

			通过类型断言可以做到以下几件事情:

			1).检查 i 是否为 nil

			2).检查 i 存储的值是否为某个类型

			如果接口值的类型，并不是我们所断言的 T，就会断言失败，
			但和第一种表达式不同的事，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败，此时t 为 T 的零值

			接口断言，为什么要接口断言？

			因为空接口 interface{}没有定义任何函数，因此 Go 中所有类型都实现了空接口。

			当一个函数的形参是interface{}，那么在函数中，需要对形参进行断言，从而得到它的真实类型。

		2. Type Switch
			如果需要区分多种类型，可以使用 type switch 断言，这个将会比一个一个进行类型断言更简单、直接、高效。
			方式二：switch
				switch instance := 接口对象.(type){
				case 实际类型1:
						....
				case 实际类型2:
						....
				....
				default:
					....
				}
			switch语句判断接口的类型。每一个case会被顺序地考虑。

			当命中一个case 时，就会执行 case 中的语句，因此 case 语句的顺序是很重要的，因为很有可能会有多个 case匹配的情况。
			若果没有命中任何 case ，那么就走 default 分支
	*/
	// 初始化一个圆结构体
	c1 := Circle{radius: 10}
	fmt.Println("=================圆结构体：================")
	fmt.Println("圆的周长为：", c1.perimeter())
	fmt.Println("圆的面积为：", c1.area())

	// 初始化一个三角形结构体
	t1 := Triangle{
		a: 3,
		b: 4,
		c: 5,
	}
	fmt.Println("===============三角形结构体：================")
	fmt.Println("三角形的周长为：", t1.perimeter())
	fmt.Println("三角形的面积为：", t1.area())

	// 初始化一个圆形结构体指针
	var c2 *Circle = &Circle{radius: 5}
	fmt.Println("==============圆形结构体指针：================")
	fmt.Println("圆的周长为：", c2.perimeter())
	fmt.Println("圆的面积为：", c2.area())

	// 1.t, ok:= i.(T) 常用于 if else 结构
	fmt.Println("==============t, ok:= i.(T) 开始接口断言====================")
	getInterfaceType(c1) // 判断该接口是否为 圆形结构体类型
	getInterfaceType(t1) // 判断该接口是否为 圆形结构体类型
	getInterfaceType(c2) // 判断该接口是否为 圆形结构体指针类型

	// 2.t := i.(type) 常用于 switch 结构
	fmt.Println("==============t := i.(type) 开始接口断言====================")
	getInterfaceTypeSwitch(c1) // 判断该接口是否为 圆形结构体类型
	getInterfaceTypeSwitch(t1) // 判断该接口是否为 圆形结构体类型
	getInterfaceTypeSwitch(c2) // 判断该接口是否为 圆形结构体指针类型
}

// 定义接口断言函数
func getInterfaceType(s Shape) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边分别为：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径为；", ins.radius)
	} else if ins, ok := s.(*Circle); ok {
		fmt.Printf("是圆形结构体指针，类型为：%T,存储的地址为：%p，指针自身的地址为：%p\n", ins, &ins, ins)
	} else {
		fmt.Println("无法判断类型...")
	}
}

// 定义接口断言函数,使用 switch
func getInterfaceTypeSwitch(s Shape) {
	switch ins := s.(type) { // 首字母小写的 type
	case Circle:
		fmt.Println("是圆形，半径为；", ins.radius)
	case Triangle:
		fmt.Println("是三角形，三边分别为：", ins.a, ins.b, ins.c)
	case *Circle:
		fmt.Printf("是圆形结构体指针，类型为：%T,存储的地址为：%p，指针自身的地址为：%p\n", ins, &ins, ins)
	default:
		fmt.Println("无法判断类型...")
	}
}

type Shape interface {
	perimeter() float64 // 返回形状的周长
	area() float64      // 返回形状的面积
}

// 定义结构体
type Circle struct {
	radius float64
}

type Triangle struct {
	a, b, c float64
}

// 圆结构体 实现接口方法
func (c Circle) perimeter() float64 {
	return c.radius * math.Pi * 2
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// 三角形结构体 实现接口方法
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
