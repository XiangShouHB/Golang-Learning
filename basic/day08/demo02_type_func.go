//File  : demo02_type_func.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Go语言支持函数式编程，可以使用高阶编程语法。
		一个函数可以作为另一个函数的参数，也可以作为另一个函数的返回值，
		那么在定义这个高阶函数的时候，如果函数的类型比较复杂，我们可以使用type来定义这个函数的类型

		注意：
			如果返回值是函数类型，自身函数的局部变量和返回值里的局部变量是各自独立的，不能共享数据
	*/

	my_fun_test := int2Str(1, 9)
	fmt.Println("my_fun_test 函数执行后..........", my_fun_test(2, 8))
}

// 定义 函数类型
type my_fun func(int, int) string

// 和写匿名函数一致
func int2Str(a, b int) my_fun { // int2Str 函数的局部变量a,b
	//resFunc := func(a, b) string {
	// 这么写会报错，因为在函数内定义的变量为局部变量，现在int2Str的返回值为 函数类型 my_fun
	// my_fun 函数类型也有自己的参数
	resFunc := func(a, b int) string { // 返回值类型 my_fun 的局部变量a,b
		fmt.Printf("匿名函数的局部变量,a = %d,b = %d\n", a, b)
		return strconv.Itoa(a) + strconv.Itoa(b)
	}
	fmt.Printf("int2Str 函数的局部变量,a = %d,b = %d\n", a, b)
	return resFunc

}
