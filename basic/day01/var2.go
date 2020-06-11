// File  : var2.go
// Author: duanhaobin
// Date  : 2020/4/24

package main

import "fmt"
var global_num = 1000
// global_string := "global string"  语法错误
func main() {
	/*
		变量学习注意事项：
		1.变量必须先定义才能使用
		2.变量的类型和赋值必须一致
		3.同一个作用域内，变量名不能冲突
		4.简短定义方式，左边的变量至少有一个是新的
		5.简短定义方式，不能定义全局变量
		6.变量的零值，就是默认值
			整型：默认值是0
			浮点类型：默认是0
			字符串：默认值""
			切片：[]
			bool ：false
		7.变量定义了就要使用，否则无法通过编译,错误信息：num declared and not used
	*/
	num := 1
	fmt.Printf("num值为：%d，地址为：%p\n",num,&num)

	// num := 2  变量名不能冲突
	num = 2
	fmt.Printf("num值为：%d，地址为：%p\n",num,&num)  // 地址未变化

	var name = "张三"

	// name = 100  类型和赋值必须一致，否则编译会报错：cannot use 100 (type int) as type string in assignment

	// 使用简短声明时，确保最少要有一个新的变量被定义
	// name := "李四"  name已经声明过了，因此报错no new variables on left side of :=

	name, age,sex := "李四", 30,"男"  //  age,sex为新变量，故能使用简短声明
	fmt.Printf("姓名值为：%s,年龄：%d,性别：%s\n",name,age,sex)

	fmt.Println("默认值————————————————————————————————")
	// 默认值
	var a int
	fmt.Println(a)  // 整数数字0

	var b float64
	fmt.Println(b)  // 浮点数默认值为0，有点意外，Java,Python中0.0

	var s string
	fmt.Println(s)  // 空字符串 ""

	var array []int
	fmt.Println(array)  // 切片 []

	// 简短声明不能定义全局变量
	fmt.Println("全局变量：",global_num)

	// var no_use = 1  该变量不使用会报错：no_use declared and not used

}
