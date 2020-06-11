//File  : demo05_interface_empty.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import (
	"fmt"
)

func main() {
	/*
		空接口(interface{})：
			不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。

		fmt包下的Print系列函数，其参数大多是空接口类型，也可以说支持任意类型：
			func Print(a ...interface{}) (n int, err error)
			func Println(format string, a ...interface{}) (n int, err error)
			func Println(a ...interface{}) (n int, err error)
	*/
	// 1.定义任意类型的数据
	var a1 Empyt_interface = Animal{"花猫"}
	var a2 Empyt_interface = Person{"Bob", 30}
	var a3 Empyt_interface = "haha"
	var a4 Empyt_interface = 100
	fmt.Println("a1 .............", a1)
	fmt.Println("a2 .............", a2)
	fmt.Println("a3 .............", a3)
	fmt.Println("a4 .............", a4)

	// 2.调用 printInfo 方法
	//a1.printInfo()  报错，无法调用 printInfo 方法。因为，结构体无法实现接口的方法， 是一个空接口，所以，不存在结构体实现接口方法

	// 3.调用 getInfo 函数
	getInfo(a1)
	getInfo(a3)

	getInfo2("10000")
	getInfo2("听歌学习中")

	// 4.定义值为任意的类型的 map
	map1 := make(map[string]Empyt_interface)
	map1["数字"] = 1
	map1["字符串"] = "字符串"
	map1["布尔"] = false
	fmt.Println("map1 ...........", map1)

	// 5.定义任意类型的 切片
	//sli := make([]Empyt_interface{}, 0, 10)  // 编译报错：]Empyt_interface literal is not a type
	sli := make([]interface{}, 0, 10) // 长度为0，容量为10
	sli = append(sli, a1, a2, a3, a4, map1, "切片元素", false)
	fmt.Println("任意类型切片 sli .....", sli)

	// 遍历切片的元素
	printSliceInfo(sli)
}

func printSliceInfo(sli []interface{}) {
	for i, v := range sli {
		fmt.Printf("第%d个元素为：%v\n", i+1, v)
	}
}

// 定义一个空接口
type Empyt_interface interface {
}

// 定义结构体
type Person struct {
	name string
	age  int
}

type Animal struct {
	color string
}

// 定义一个任意类型的方法
func (arg Person) printInfo() {
	fmt.Println("printInfo 方法.....", arg.name)
}

// 定义一个入参为任意类型的函数

func getInfo(arg Empyt_interface) {
	fmt.Println("getInfo 函数.....", arg)
}

// 也可以写成如下形式
func getInfo2(arg interface{}) {
	fmt.Println("getInfo2 函数.....", arg)
}
