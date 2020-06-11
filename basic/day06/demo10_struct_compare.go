//File  : demo10_struct_compare.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

func main() {
	/*
		结构体比较：
			1.结构体是值类型，如果每个字段具有可比性，则是可比较的。如果它们对应的字段相等，则认为两个结构体变量是相等的。
			2.如果结构变量包含的字段是不可比较的，那么结构变量是不可比较的
	*/
	// 1.结构体的字段是可比较的
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// 2.结构体的字段是不可比较的

	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}

	//if image1 == image2 { // 编译报错：invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
	//	fmt.Println("image1 and image2 are equal")
	//}

	fmt.Println("image1 ->", image1)
	fmt.Println("image2 ->", image2)
}
