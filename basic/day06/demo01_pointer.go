//File  : demo01_pointer.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

func main() {
	/*
	指针：
		定义：
			指针是存储另一个变量的内存地址的变量
		获取变量的地址：
			Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
		声明指针：
			声明指针，*T是指针变量的类型，它指向T类型的值。
			语法： var a *int  声明了一个名为a的指针，a指向int类型
			指针变量的赋值操作只能接受 内存地址
		空指针：
			1.Go 空指针 当一个指针被定义后没有分配到任何变量时，它的值为 nil。 nil 指针也称为空指针。
			2.nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
			3.一个指针变量通常缩写为 ptr。
		获取指针的值：
			获取一个指针意味着访问指针指向的变量的值。
			语法：*ptr
		操作指针改变变量的数值:
			*ptr = new_value, 改变值后，ptr存储的内存地址值 指向的变量引用的值 也会发生改变

		指针的指针：
			如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
			语法：**pptr
	*/
	// 1.获取变量的内存地址
	a := 100
	fmt.Println("a 的内存地址：",&a)

	// 2.声明指针
	var ptr *int
	fmt.Println("p 为指针，其值为：",ptr)  // <nil> 空指针

	// 指针变量 p 也有自己的内存地址
	fmt.Println("p 指针变量自己的内存地址：",&ptr)  // 0xc0000ca020

	// 指针变量的赋值操作只能接受 内存地址
	ptr = &a  // 将 a 的内存地址赋值给 ptr 指针
	fmt.Println("p 指针变量的存储地址为：",ptr)  // 0xc0000a0068 和a的内存地址一样

	// 3.获取指针的值
	fmt.Println("获取指针 ptr 的值：",*ptr)  // 100，和 a 的值一致，因为 ptr 指针存储的是 a 的内存地址

	// 4.操作指针改变变量的数值
	*ptr = 200
	fmt.Println("a 变量的值也改变了：",a)   // 200
	fmt.Println("获取指针 ptr 的值",*ptr)  // 200

	// 5.指针的指针
	var pptr **int
	fmt.Println("pptr 指针的值为：",pptr)  // <nil>

	pptr = &ptr
	fmt.Println("pptr 指针变量的存储地址为：", pptr)  // 存储的是 ptr指针 的内存地址

	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)



}

