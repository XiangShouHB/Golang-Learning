//File  : array_demo1.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：整数，浮点，布尔，字符串
			复合类型：array，slice，map，struct，pointer，function，channel。。。

		数组：
			1.概念：存储一组相同数据类型的数据结构。Python中，数据类型可不相同，任意类型均可
					理解为容器，存储一组数据
			2.语法：
					var 数组名 [长度] 数据类型
					var 数组名 = [长度] 数据类型{元素1，元素2.。。}
					数组名 := [长度]数据类型{元素。。。}
					可以忽略声明中数组的长度并将其替换为…让编译器为你找到长度
					如：数组名 := [...]数据类型{元素。。。}
			3.通过下标访问
				下标，也叫索引：index，
				默认从0开始的整数，直到长度减1
				数组名[index]
					赋值
					取值

				不能越界：[0,长度-1]

			4.长度和容量：go语言的内置函数
				len(array/map/slice/string)，长度
				cap()，容量
			5.内存分析
				arr := [3] int
	            执行上述代码:
				1).Go 语言分配一个内存空间，该内存空间被等分成三份，第一份内存的地址也叫首地址，同时也是arr指向的内存地址
				2).Go 语言中，数组是连续的内存，知道了首项地址，就可以推断出后面的值
	 */

	// 1.定义
	// 创建方式1
	var arrInt1 [5] int  // 未初始化，默认元素都为0
	fmt.Println("arrInt1 :",arrInt1)


	// 创建方式2
	var arrInt2 = [3]int{1,3,4}
	fmt.Println("arrInt2 :",arrInt2)

	// 创建方式3 定义一个容量为4，长度为4的数组
	arrInt3 := [4]int{1,5,9}  // 元素数量不够，会使用默认值来填充
	fmt.Println("arrInt3 :",arrInt3)
	fmt.Println("数组的长度：",len(arrInt3)) //容器中实际存储的数据量
	fmt.Println("数组的容量：",cap(arrInt3)) //容器中能够存储的最大的数量

	// 创建方式4  给定数组大小，且初始化指定元素的值
	var arrInt4 = [5]int{1:1,3:2}  // 下标为1的元素值为1，下标为3的元素值为2
	fmt.Println("arrInt4 :",arrInt4)

	// 创建方式5 不指定大小，由编译器推断
	arrInt5 := [...]int{1,2,3,4,5}
	fmt.Println("arrInt5 :",arrInt5)
	fmt.Println(len(arrInt5))

	arrInt6:=[...]int{1:3,6:5}
	fmt.Println("arrInt6 :",arrInt6)
	fmt.Println(len(arrInt6))

	// 2.数组的访问
	fmt.Println("数组的第一个元素：", arrInt3[0])


}

