//File  : data_type.go
//Author: duanhaobin
//Date  : 2020/4/25

package main

import "fmt"

func main() {
	/*
	基本语法--数据类型：
	1.基本数据类型：
			布尔类型：bool
				取值：true，false
			数值类型：
				整数：int
					有符号：最高位表示符号位，0正数，1负数，其余位表示数值
						int8:(-128 到 127)
						int16:(-32768 到 32767)
						int32:(-2147483648 到 2147483647)
						int64:(-9223372036854775808 到 9223372036854775807)
					无符号：所有的位表示数值
						uint8： (0 到 255)
						uint16：(0 到 65535)
						uint32：(0 到 4294967295)
						uint64： (0 到 18446744073709551615)

					int, uint

					byte:uint8
					rune:int32
				浮点：生活中的小数
					float32,float64
				复数：complex，
			字符串：string
	 2.复合数据类型
		array，slice，map，function，pointer，struct，interface，channel。。。
	 */
	// 1.布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T,%t\n",b1,b1)
	b2 :=false
	fmt.Printf("%T,%t\n",b2,b2)

	// 2.整数
	var i1 int8
	i1 = 100
	fmt.Println(i1)
	var i2 uint8
	i2 = 200
	fmt.Println(i2)

	var i3 int
	i3 = 1000
	fmt.Printf("i3的类型为；%T,值为：%d\n",i3,i3)

	// 注意：语法角度：int，int64不认为是同一种类型
	//var i4 int64
	//i3 == i4  invalid operation: i3 == i4 (mismatched types int and int64)

	var i5 byte  // byte相当于unit8
	i5 = 255  // 范围为0-255
	fmt.Printf("i5的类型：%T,值为：%d\n",i5,i5)

	var i6 = 100  // 变量声明时赋值100，Golang推断的数据类型为int
	fmt.Printf("i6的类型为：%T,值为：%d\n",i6,i6)

	// 浮点
	var f1 float32
	f1 = 3.111111111111
	fmt.Printf("f1 的类型为：%T,值为：%.3f\n",f1,f1)

	var f2 float64
	f2 = 4.222222222222222   // 64位系统下，精度更大
	fmt.Printf("f2 的类型为：%T,值为：%.5f\n",f2,f2)

	var f3 = 5.333 // 变量声明时赋值100，Golang推断的数据类型为float
	fmt.Printf("f3的类型为：%T,值为：%f\n",f3,f3)  // 格式化输出时，浮点数默认保留小数点后6位，若没有值以0补充

}

