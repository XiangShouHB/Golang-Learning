//File  : demo02_pointer_arr.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

func main() {
	/*
			数组指针：
				首先，它是一个指针，但是数据类型为数组，或者说指向数组
				语法：
					var arrPtr *[size]int  创建一个指针arrPtr,指向一个int类型的数组，大小为size

				*号可以不写问题：
					var arr = [4] int {1,2,3,4}
					var arrPtr *[4] int
					arrPtr = &arr
					fmt.Println(* arrPtr[0])
					以上代码执行有问题吗？

					在go语言中*寻址运算符和[]中括号运算符的优先级是不同的！
					[]中括号是初等运算符
					*寻址运算符是单目运算符

		            初等运算符的优先级是大于单目运算符的，因此先参与计算的是 arrPtr [0];
					arrPtr[0]其实就是数组的第一个元素，就是数字1
		            数字1必然是int类型，而不是一个地址，因此针对数字1使用*寻址运算符自然也就发生了错误。
		        	解决问题的办法很简单，就是添加一个小括号就可以了
					即：
		                (*p)[0]
					不过因为*在go语言中，建立了 p:=&arr 这种类似地址关系后，*允许不写。
					所以，访问时候可以直接写成p[0]。事实上在工作开发过程中，这种写法反而更加常见。
					ps：
		            仅对访问下标时，*寻址运算符允许不写！


			指针数组：
				首先，它是一个数组，数组中元素类型都是指针，本质上就是，元素都为地址的数组
				语法：
					var ptrArr [size]*int  创建一个数组ptrArr,大小为size,数据类型为 *int

			*[5]float64,指针，一个存储了5个浮点类型数据的数组的指针
			*[3]string，指针，数组的指针，存储了3个字符串
			[3]*string，数组，存储了3个字符串的指针地址的数组
			[5]*float64，数组，存储了5个浮点数据的地址的数组
			*[5]*float64，指针，一个数组的指针，存储了5个float类型的数据的指针地址的数组的指针
			*[3]*string，指针，存储了3个字符串的指针地址的数组的指针
			**[4]string，指针，存储了4个字符串数据的数组的指针的指针
			**[4]*string，指针，存储了4个字符串的指针地址的数组，的指针的指针

	*/

	// 1.数组指针
	var arrPtr *[4]int // 创建一个指针，其数据类型为数组
	fmt.Println("数组指针 arrPtr 为：", arrPtr)

	var arr = [4]int{1, 2, 3, 4}
	arrPtr = &arr
	fmt.Println("将 arr 的内存地址赋值给数组指针 arrPtr,   arrPtr=", arrPtr) // &[1 2 3 4]  不应该是地址吗？

	fmt.Printf("arr 数组的地址为：%p\n", &arr)      // 0xc00009e140
	fmt.Printf("arrPtr 存储的地址为：%p\n", arrPtr) // 0xc00009e140  存储的

	fmt.Printf("arrPtr 指针自己的地址为：%p\n", &arrPtr) // 0xc0000ca018  指针ptr自己的地址值

	// 2. * 可以不写实践
	//fmt.Println("通过指针访问数组的第一个元素：",*arrPtr[0])  // 编译错误，操作符先后顺序引起的问题 [] > *
	fmt.Println("通过指针访问数组的第一个元素：", (*arrPtr)[0])
	fmt.Println("通过指针访问数组的第一个元素：", arrPtr[0])

	// 3.指针数组
	var ptrArr [4]*int

	a, b, c, d := 1, 2, 3, 4
	arr2 := [4]int{a, b, c, d}
	fmt.Println("数组 arr2 :", arr2)

	ptrArr = [4]*int{&a, &b, &c, &d} // 存的都是内存地址
	fmt.Println("指针数组 ptrArr :", ptrArr)

	// 4.操作数据，查看变化
	// 4.1 数组变化
	arr2[0] = 100                          // arr2的第一个元素改变, a会不会变化，ptrArr会不会变化？
	fmt.Println("a 的值为；", a)               // 1  arr2是值拷贝，变化都和a无关
	fmt.Println("ptrArr 的值为；", *ptrArr[0]) // 1  ptrArr[0]指向的是 a 的内存地址，a没变，那么 *ptrArr也不会变化

	fmt.Println("ptrArr[0] 的值：", ptrArr[0])
	fmt.Printf("a 的内存地址为：%p\n", &a)

	// 4.2 指针数组变化，本质上是 变量a变化
	*ptrArr[0] = 1000              // 指针数组的第一个元素地址指向的值发生改变， a会不会变化？arr2会不会变化
	fmt.Println("a 的值为：", a)       // 1000 变化，道理同上
	fmt.Println("arr2 的值为：", arr2) // [100 2 3 4]不变  道理同上

	//4.3 变量b变化
	b = 2000                               // b发生变化， arr2 会不会变化？ ptrArr 会不会变化？
	fmt.Println("arr2 的值为：", arr2)         // [100 2 3 4] 不变 道理同上
	fmt.Println("ptrArr 的值为：", *ptrArr[1]) // 2000  变化 道理同上

}
