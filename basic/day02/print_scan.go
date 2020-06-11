//File  : print_scan.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
		输入和输出：
			fmt包：输入，输出

			输出：
				Print() //打印
				Printf() //格式化打印
				Println() //打印之后换行

			格式化打印占位符：
				%v,默认格式的值
				%T，打印类型
				%t,bool类型
				%s，字符串
				%f，浮点
				%d，10进制的整数
				%b，2进制的整数
				%o，8进制
				%x，%X，16进制
					%x：0-9，a-f
					%X：0-9，A-F
				%c，打印字符
				%p，打印地址
				。。。

			输入：
				Scanln()
					Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
				Scanf()

			bufio包
	 */

	// 1.输出函数
	a := 100
	b := 10.2
	c := "Hello"
	d := `world`
	e := true
	f := 'a'
	fmt.Printf("a：%d,类型为：%T\n",a,a)
	fmt.Printf("b：%f,类型为：%T\n",b,b)
	fmt.Printf("c：%s,类型为：%T\n",c,c)
	fmt.Printf("d：%s,类型为：%T\n",d,d)
	fmt.Printf("e：%t,类型为：%T\n",e,e)
	fmt.Printf("f：%v,类型为：%c\n",f,f)
	fmt.Printf("f：%d,类型为：%T,字符为：%c\n",f,f,f)

	// 2.输入函数
	var x int
	var y float64
	fmt.Println("请输出一个整数，一个浮点数")
	//fmt.Scanln(&x,&y)  // 以地址接受
	//fmt.Printf("x：%d,y：%f", x,y)

	fmt.Scanf("%d,%f",&x, &y)  // 格式化输入，中间要有 ','
	fmt.Printf("x：%d,y：%f", x,y)

}

