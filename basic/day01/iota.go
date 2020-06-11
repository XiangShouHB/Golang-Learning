//File  : iota.go
//Author: duanhaobin
//Date  : 2020/4/25

package main

import "fmt"

func main() {
	/*
		iota：特殊的常量，可以被编译器自动修改的常量
			1.每当定义一个const，iota的初始值为0
			2.每当定义一个常量，就会自动累加1
			3.直到下一个const出现，清零
			4.如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
			5.自增默认是int类型，可以自行进行显示指定类型
			6.iota 可以参与运算
	*/
	// 1.定义const枚举，首项为iota
	const (
		one   = iota //iota初始值为0
		two   = iota //自增1
		three = iota //再自增1
	)
	fmt.Println("one 的值为：", one)
	fmt.Println("two 的值为：", two)
	fmt.Println("three 的值为：", three)

	// 2.再次出现一个const，iota会清零
	const (
		D = iota
		E
	)
	fmt.Println("D的值为：", D)
	fmt.Println("E的值为：", E)
	fmt.Println("-------------------------")
	// 3.如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
	const (
		F = iota // 0
		G        // 1
		H        // 2
		I        // 3
		J = "断开" // iota终端 iota=4
		K = iota // 显示恢复iota,iota=5
		L        // iota=6
	)
	fmt.Println("F的值为：", F)
	fmt.Println("G的值为：", G)
	fmt.Println("H的值为：", H)
	fmt.Println("I的值为：", I)
	fmt.Println("J的值为：", J)
	fmt.Println("K的值为：", K)
	fmt.Println("L的值为：", L)

	// 4. iota 参与计算
	const (
		b  = 1 << (iota * 10) // iota初始值为0 ，所以 1 << (0 * 10)
		kb                    // 1 << (1 * 10)
		mb                    // 1 << (2 * 10)
		gb                    // 1 << (3 * 10)
		tb                    // 1 << (4 * 10)
		pb                    // 1 << (5 * 10)
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
