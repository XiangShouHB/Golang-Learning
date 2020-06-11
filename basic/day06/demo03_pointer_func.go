//File  : demo03_pointer_func.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import (
	"fmt"
)

func main() {
	/*
		函数指针：一个指针，指向了一个函数的指针。
			因为go语言中，function，默认看作一个指针，没有*。
			slice,map,function

		指针函数：一个函数，该函数的返回值是一个指针。

	*/
	// 1.函数指针
	f := func1
	fmt.Printf("函数的类型：%T,函数的地址：%p\n",f,f)

	// 2.指针函数
	var arrPtr *[3] int  // 定义一个指针，指向数组
	arrPtr = func2()
	fmt.Printf("arrPtr 的类型：%T,地址：%p,值：%v\n",arrPtr,&arrPtr,arrPtr)
	fmt.Printf("arrPtr 指针中，存储的数组地址为：%p\n",arrPtr)  // 和 func2 函数中 arr 的内存地址是一致的
}

func func1()  {
}

func func2() *[3] int  {
	arr := [3] int{1,2,3}
	fmt.Printf("func2 函数中 arr 的地址：%p\n",&arr)
	return &arr
}