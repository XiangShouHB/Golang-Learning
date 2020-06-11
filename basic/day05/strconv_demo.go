//File  : strconv_demo.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本类型之前的转换
			string convert
		1.整数和字符串之间的转换:
			最常用：Atoi() Itoa()
			以下两个方法也有：
			ParseInt()
			FormatInt()
		2.ParseXXXX() 表示将字符串解析为 XXXX 类型
		3.FormatXXXX() 表示将 XXXX 类型格式化为字符串

		ParseInt(s string, base int, bitSize int)方法参数解析:

		参数 base 代表字符串按照给定的进制进行解释。一般的，base 的取值为 2~36，如果 base 的值为 0，
		则会根据字符串的前缀来确定 base 的值："0x" 表示 16 进制； "0" 表示 8 进制；否则就是 10 进制。

		参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。
		取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	*/

	// 1.整数与字符串之间的转换
	// Atoi() ASCII to int
	s1 := "10"
	num1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("String to int failed,error:", err)

	}
	fmt.Printf("字符串转换为整数；%d，类型：%T\n", num1, num1)

	// Itoa() int to ASCII
	s2 := strconv.Itoa(1314)
	fmt.Printf("整数转换为字符串；%s, 类型：%T\n", s2, s2)

	// ParseInt()  将字符串解析为整数
	num2, err := strconv.ParseInt("100", 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseInt failed, error:", err)
	}
	fmt.Printf("字符串解析为整数(10进制数)结果：%d，类型：%T\n", num2, num2)

	// 2.其他类型格式化为字符串
	// bool -> string
	flag_bool := true
	str_bool := strconv.FormatBool(flag_bool)
	fmt.Printf("bool 类型格式化为字符串结果：%s，类型：%T\n", str_bool, str_bool)

	// float -> string
	//fmt:格式
	//prec:精度
	//bitsize:位数  float64 或 32位
	f1 := 3.14
	str_float := strconv.FormatFloat(f1, 'E', -1, 64)
	fmt.Printf("float 类型格式化为字符串结果：%s，类型：%T\n", str_float, str_float)

}
