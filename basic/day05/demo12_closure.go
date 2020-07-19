//File  : demo12_closure.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程：
			支持将一个函数作为另一个函数的参数，
			也支持将一个函数作为另一个函数的返回值。

		闭包(closure)：
			一个外层函数中，有内层函数，该内层函数中，
				1.会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)
				2.并且该外层函数的返回值就是这个内层函数。

			这个内层函数和外层函数的局部变量，统称为闭包结构。

			局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
			但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。


		还有一点需要注意：
			使用闭包的过程中，一旦外函数被调用一次返回了内函数的引用，虽然每次调用内函数，是开启一个函数执行过后消亡，
			但是闭包变量实际上只有一份，每次开启内函数都在使用同一份闭包变量
	*/

	//
	res := outter(100)  // 返回的是函数
	fmt.Println("闭包函数的结果为：",res())


	// 闭包变量实际上只有一份，每次开启内函数都在使用同一份闭包变量
	res2 := outter2()
	fmt.Println("第一次调用res2,结果为：",res2(100))  // 100+100
	fmt.Println("第二次调用res2,结果为：",res2(2200))  // 200+2200 在原来的基础上进行了计算

}

// 定义一个闭包函数
func outter(out_var int) func() int  {
	inner := func() int{// 匿名函数

		// 1.操作外层函数的变量, 在Python中，内层函数是无法操作外层函数变量的,除非给外层函数变量加 ’nonlocal‘ 关键字
		out_var += 100
		return out_var
	}
	// 2.返回值为内层函数
	return inner
}

// 定义一个闭包函数
func outter2() func(inn_var int) int  {
	out_var := 100
	inner := func(inn_var int) int{// 匿名函数

		// 1.操作外层函数的变量, 在Python中，内层函数是无法操作外层函数变量的,除非给外层函数变量加 ’nonlocal‘ 关键字
		out_var += inn_var
		return out_var
	}
	// 2.返回值为内层函数
	return inner
}

