//File  : if_demo.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
		条件语句：if
		语法格式：
			if 条件表达式{
				//
			}

		if语句的其他写法：
		if 初始化语句; 条件{
			//注意变量的作用域问题
		}

	*/

	num := 6
	if num > 10 {
		fmt.Println("大于10")
	}else {
		fmt.Println("小于10")
	}
	// 上述代码可以合并
	if num := 6; num > 10{
		fmt.Println("大于10")
	}else {
		fmt.Println("小于10")
	}
	//需要注意的是，num的定义在if里，那么只能够在该if..else语句块中使用，否则编译器会报错的

}

