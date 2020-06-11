//File  : map_demo2.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		map的遍历：
			传统：for(i=0,i < len(map); i++){}
			使用：for range 推荐使用

				数组，切片：index，value
				map：key，value
	*/
	map1 := make(map[int] string)
	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"
	map1[4] = "four"
	map1[5] = "five"

	fmt.Println("map1 为：",map1)

	// 1. 传统for循环  这种限制比较大，通常适用于key为连续的整数，所以不推荐使用
	for i:=1; i <= len(map1);i++{
		fmt.Println("key = ",i, "value = ",map1[i])
	}
	fmt.Println("--------------------------------------")
	// 2. range 循环
	for k,v := range map1{
		fmt.Println("key = ",k, "value = ",v)  // 结果是乱序的
	}

	for _,v := range map1{  // 同理 k,_ := range map1 即可输出key
		fmt.Println("value = ", v)
	}
}

