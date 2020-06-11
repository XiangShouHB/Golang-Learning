//File  : type_convert.go
//Author: duanhaobin
//Date  : 2020/4/25

package main

import "fmt"

func main() {
	/*
		数据类型转换：Type Convert
		go语言是静态语言，定义，赋值，运算必须类型一致

		语法格式：Type(Value)

		注意点：兼容类型可以转换

		常数：在有需要的时候，自动转型
		变量：需要手动转型
	*/

	// 1.普通转换,同类型转换
	var a int8
	a = 10
	var b int16
	b = int16(a)
	fmt.Printf("int8类型a转为int16类型，并赋值给b。a:%d,b:%d\n",a,b)

	// 2.不同类型转换（兼容类型可以转换）
	var c float64
	c = 1.45
	b = int16(c)  // 取c的整数部分赋值给b
	fmt.Printf("float64类型c转为int16类型，并赋值给b。c:%f,b:%d\n",c,b)

	// 3.不兼容类型无法转换
	//var d = "hello"
	//b = int16(d)  //cannot convert d (type string) to type int16

	// 4.常数转换：在有需要的时候，会自动转型
	var e = c + 10
	fmt.Printf("常数会自动转型,e的数据类型为：%T，值为：%f\n",e,e)



}

