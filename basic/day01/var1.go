//File  : var1.go
//Author: duanhaobin
//Date  : 2020/4/24

package main

import (
	"fmt"
)

func main() {
	/*
	基本语法--变量
	 */

	// 1.先声明和赋值
	var num1 int
	num1 = 10
	fmt.Printf("num1为:%d\n", num1)

	// 2.声明加赋值,根据值自行判定变量类型
	var num2 = 20
	fmt.Printf("num2为:%d\n",num2)

	// 3. := 简短声明
	num3 := 30
	fmt.Printf("num3类型为:%T,值为:%d\n",num3,num3)

	// 多变量同时声明

	// 1.以逗号分隔，声明与赋值分开，若不赋值，则使用默认值
	var name1, name2, name3 string
	name1, name2, name3 = "张三", "李四", "王五"
	fmt.Printf("name1:%s,name2:%s,name3:%s\n",name1,name2,name3)

	// 2.直接赋值，下面的变量类型可以是不同的类型
	var a, b, c = 10, "20", 1.5
	fmt.Printf("a:%d,b:%s,c:%f\n",a,b,c)

	// 3.第三种，集合类型
	var(
		studentName = "马七"
		age = 20
		sex = "男"
	)
	fmt.Printf("学生姓名:%s,年龄:%d,性别:%s\n",studentName,age,sex)
}



