//File  : string_type.go
//Author: duanhaobin
//Date  : 2020/4/25

package main

import "fmt"

func main() {
	/*
		字符串：
		1.概念：多个byte的集合，理解为一个字符序列
		2.语法：使用双引号
			"abc","hello","A"
				也可以使用``
		3.编码问题
				计算机本质只识别0和1
				A：65，B:66,C:67...
				a:97,b:98...
			ASCII(美国标准信息交换码)

			中国的编码表：gbk，兼容ASCII
				中
				家
			Unicode编码表：号称统一了全世界
				UTF-8，UTF-16,UTF-32...

		4.转义字符：\
			A：有一些字符，有特殊的作用，可以转义为普通的字符
				\',\'
			B：有一些字符，就是一个普通的字符，转义后有特殊的作用
				\n,换行
				\t,制表符
	*/
	// 1.定义字符串
	var s1 string
	s1 = "Hello"
	fmt.Printf("s1的数据类型：%T,值为：%s\n",s1,s1)

	s2 := `world`
	fmt.Printf("s2的数据类型：%T,值为：%s\n",s2,s2)

	// 2.注意单引号 '' 并不能表示字符串，Python中可以表示
	var a = 'A'
	fmt.Printf("a的数据类型：%T,值为：%d\n",a,a)  // 在ASCII编码中，A对应编码为65

	// 3.中文
	var ch = '中'
	// %c 查看汉字'中'的编码，,
	// %d 双(单)引号围绕的字符串，由 Go 语法安全地转义
	fmt.Printf("ch的数据类型：%T,Unicode编码值为：%d，汉字为：%c，原字符串：%\nq",ch,ch,ch,ch)

	// 4.转义字符:\
	fmt.Println("\"转义字符实践\"")

	// 5.字符串的嵌套表示， ""和``
	s3 := `Hello" world"`
	fmt.Println("s3为：",s3)  // 将内层的双引号也输出

	s4 := "Hello`world`"
	fmt.Println("s3为：",s4)  // 同理




}

