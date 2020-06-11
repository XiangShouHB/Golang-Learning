//File  : slice_demo4.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是数据本身。
			值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct


		浅拷贝：拷贝的是数据地址。
			导致多个变量指向同一块内存
			引用类型的数据，默认都是浅拷贝：slice，map，

			因为切片是引用类型的数据，直接拷贝的是地址。

		func copy(dst, src []Type) int
			其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice）
			目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致
			copy() 函数的返回值表示实际发生复制的元素个数。

	*/
	// 1.深拷贝
	arr := [5]int{1,2,3,4,5}
	arr2 := arr
	fmt.Println("arr 为:",arr)
	fmt.Println("arr2 为:",arr2)
	arr2[1] = 20 // arr2元素变化和arr1没有关系
	fmt.Println("arr 为:",arr)
	fmt.Println("arr2 改变后为:",arr2)

	fmt.Println("--------------------------------")
	// 2.浅拷贝  slice为例，只是拷贝的引用地址
	s1 := make([]int, 0)  // 长度0，容量默认为0
	s2 := []int{6,7,8,9}
	fmt.Println("s1 为:",s1)
	fmt.Println("s2 为:",s2)
	s1 = s2
	fmt.Println("s1拷贝后 为:",s1)

	// 改变s1，s2也会改变
	s1[0] = 60
	fmt.Println("s1 改变后为:",s1)
	fmt.Println("s2 为:",s2)  // 也发生了改变

	// copy()函数
	s3 := make([]int,5,5)
	copy(s3, s2)  // 将s2深拷贝给s3
	fmt.Println("s3 copy后为:",s3)
	fmt.Println("s2 为:",s2)

	// 改变s2，s3不会改变
	s2[1] = 70
	fmt.Println("s3 为:",s3) // 未发生改变
	fmt.Println("s2 改变后为:",s2)  // 第二个元素变为70

	// 注意目标切片必须分配过空间且足够承载复制的元素个数
	s4 := make([]int,0)  // 创建一个空切片
	copy(s4, s3) // 将s3 copy()给s4
	fmt.Println("s4 copy后为:",s4)  // s4定义的切片长度是0，因此不会有任何值可以被拷贝到s4中，因此输出空数组
	fmt.Println("s3 为:",s3)
}

