//File  : slice_demo1.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		数组array：
			存储一组相同数据类型的数据结构。
				特点：定长

		切片slice：
			同数组类似，也叫做变长数组或者动态数组。
				特点：变长

			是一个引用类型的容器，指向了一个底层数组。
		make()
			func make(t Type, size ...IntegerType) Type

			第一个参数：类型
				slice，map，chan
			第二个参数：长度len
				实际存储元素的数量
			第三个参数：容量cap
				最多能够存储的元素的数量


		append()，专门用于向切片的尾部追加元素
			slice = append(slice, elem1, elem2)
			slice = append(slice, anotherSlice...)
	*/

	// 声明数组
	arr := [5] int{1,2,3,4,5}  //必须有数组长度，或者[...]
	fmt.Println("数组arr:",arr)

	// 声明切片，不写数组长度的为切片
	var s1 [] int  // 空切片
	fmt.Println("切片s1为：",s1)

	s1 = arr[1:4]
	fmt.Println("切片s1为：",s1)
	fmt.Println("未改变之前的数组arr：",arr)

	// 对现有数组的引用，对slice所做的任何修改都将反映在底层数组中
	s1[0] = 10
	s1[1] = 20
	fmt.Println("改变后的切片s1为：",s1)
	fmt.Println("数组也被改变了，arr为：",arr)

	s2:= [] int{1,2,3}
	fmt.Println("切片s2为：",s2)

	// 查看类型
	fmt.Printf("数组类型为：%T,切片类型：%T\n",arr,s1)

	// 使用make函数定义切片 func make(t Type, size ...IntegerType) Type
	s3 := make([]int, 3, 10) // 数据类型为 int，长度为3，容量为10的切片
	s3[0] = 1
	fmt.Println("切片s3为：",s3)


	// append()，专门用于向切片的尾部追加元素
	s3 = append(s3, 2,3,4)
	fmt.Println("切片s3为：",s3)

	// 追加一个数组或切片
	s3 = append(s3, s2...)
	fmt.Println("切片s3为：",s3)


	// 遍历切片，和遍历数组一样
	fmt.Println("------------------------------------")

	for i:=0; i < len(s2);i++{
		fmt.Printf("下标->%d,元素值->%d\n",i,s2[i])
	}

	fmt.Println("------------------------------------")
	for i,v := range s2{
		fmt.Printf("下标->%d,元素值->%d\n",i,v)

	}
}

