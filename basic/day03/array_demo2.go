//File  : array_demo2.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
		一维数组：存储的多个数据是数值本身
			a1 :=[3]int{1,2,3}

		二维数组：存储的是一维的一维
			a2 := [3][4]int{{},{},{}}

			该二维数组的长度，就是3。
			存储的元素是一维数组，一维数组的元素是数值，每个一维数组长度为4。

		多维数组：。。。

		数组遍历
	 */
	arrInt := [...]int{1,2,3,4,5}
	// 1.遍历一维数组  传统方式
	for i:=0; i < len(arrInt); i++{
		fmt.Println(arrInt[i])
	}

	// 2.range方式遍历 下标与值都可以输出
	for i,v := range arrInt{
		fmt.Printf("下标：%d,元素值：%d\n",i,v)
	}
	// 如果只需要值并希望忽略索引，那么可以通过使用_ blank标识符替换索引来实现这一点。
	for _,v := range arrInt{
		fmt.Printf("元素值：%d\n",v)
	}

	// 3.多维数组
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Printf("二维数组的地址：%p\n", &a2)
	fmt.Printf("二维数组的长度：%d\n", len(a2))

	//遍历二维数组
	for i:=0;i<len(a2);i++{
		for j:=0;j<len(a2[i]);j++{
			fmt.Print(a2[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------")
	//for range 遍历二维数组
	for _,arr := range a2{
		for _,val := range arr{
			fmt.Print(val,"\t")
		}
		fmt.Println()
	}


}

