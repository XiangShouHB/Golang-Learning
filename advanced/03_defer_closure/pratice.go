//File  : pratice.go
//Author: duanhaobin
//Date  : 2020/5/13

package main

import "fmt"

func GetFn() func() {
	fmt.Println("outside")
	return func() {
		fmt.Println("inside")
	}
}
func main() {
	/*
		错误输出结果:
			main...
			outside
			inside
		正确输出结果：
			outside
			main...
			inside
		解析：
			GetFn() 可以看做是一个变量名，那么 GetFn()()可以表示为 F()
			真正 defer 的是F()函数，对于F来讲，它只是一个字面量，也就是说，GetFn()还是正常执行，没有被defer
			因此在 main 函数里， 先执行GetFn()  然后F()被defer，所以就执行了打印 main...
			此时 main 函数没有其他代码可执行了，因此就执行 defer 的代码 F(),就输出了 inside
	*/
	defer GetFn()()
	fmt.Println("main...")
}
