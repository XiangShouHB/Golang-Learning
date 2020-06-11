//File  : array_type.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：int，float，string，bool。。
			复合类型：array，slice，map，function，pointer，channel。。

		数组的数据类型：
			[size]type
			数组的大小是类型的一部分。因此[5]int和[25]int是不同的类型

		值类型：理解为存储的数值本身
			当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
				如：int,float,string,bool,array

		引用类型：理解为存储的数据的内存地址
				如：slice，map。。
	*/
	// 值类型
	s1 := "Hello" // s1开辟一个内存空间，有自己的地址
	s2 := s1  // s2开辟一个内存空间，有自己的地址，但是将s1的副本(Hello)分配给s2，s2改变，s1不会发生变化
	fmt.Printf("s1地址：%p,s2地址：%p\n",&s1,&s2)
	s2 = "World"
	fmt.Println("s1：",s1)
	fmt.Println("s2：",s2)

	arr1 := [3] int {1, 2, 3}
	arr2 := arr1
	arr2[1] = 100 // arr2改变，arr1不会改变
	fmt.Println("arr1：",arr1)
	fmt.Println("arr2：",arr2)

	// 数组的大小是类型的一部分
	fmt.Println(arr1 == arr2)  //比较数组的对应下标位置的数值是否相等
	//arr3 := [2] int{}
	//fmt.Println(arr1 == arr3)   invalid operation: arr1 == arr3 (mismatched types [3]int and [2]int)
}

