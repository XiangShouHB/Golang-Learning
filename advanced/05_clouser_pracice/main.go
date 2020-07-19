package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
	闭包应用：
	1. 实现斐波那契数列
	2. 函数式接口
 */
func main() {
	// 第一种方式：需要不断调用实现的闭包函数
	f :=fibonacci()
	fmt.Println(f())
	fmt.Println("-----------------------------")

	//	第二种方式：通过函数式接口，让调用者更加友好
	f2:=fibonacci2()
	printFileContent(f2)
}


// 1. 通过闭包实现一个斐波那契数列
// 实例：1,1,2,3,5,8,13,21,34......
//	    a,b
//        b,a+b
func fibonacci()  func() int{
	a, b := 0,1
	return func() int {
		a, b = b, a+ b // 往后移一位
		return a
	}
}

// 2.函数式接口实现
// 通过实现 Read 接口，类似读文件的思想，让其不断读取斐波那契值，直到读取到EOF退出，返回结果

type iFibonacci func() int

// 改造下生成斐波那契值的函数
func fibonacci2()  iFibonacci{
	a, b := 0,1
	return func() int {
		a, b = b, a+ b // 往后移一位
		return a
	}
}


func (i iFibonacci) Read(p []byte) (n int, err error) {
	next:= i()
	// 输出2000以下的斐波那契值即可
	if next > 2000{
		return  0,io.EOF
	}
	// 将next值转为格式化后的字符串，方便Rread，一行行的读
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		
	}
}
