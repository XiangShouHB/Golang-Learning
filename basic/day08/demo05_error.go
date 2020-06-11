//File  : demo05_error.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		1.错误是什么?

			错误指的是可能出现问题的地方出现了问题。比如打开一个文件时失败，这种情况在人们的意料之中 。

			而异常指的是不应该出现问题的地方出现了问题。比如引用了空指针，这种情况在人们的意料之外。
			可见，错误是业务过程的一部分，而异常不是 。

			Go中的错误也是一种类型。错误用内置的error 类型表示。就像其他类型的，如int，float64，。
			错误值可以存储在变量中，从函数中返回，等等。

			如果一个函数或方法返回一个错误，那么按照惯例，它必须是函数返回的最后一个值

		2.如何处理错误？
			处理错误的惯用方法是将返回的错误与nil进行比较。nil值表示没有发生错误，而非nil值表示出现错误。

			我们检查错误是否为nil。如果它不是nil，我们只需打印错误并从主函数返回。

		3.错误类型源码：
			在 builtin.go 文件下，定义了错误类型：
				type error interface {
					Error() string
				}
			它包含一个 Error() 方法，返回值为 string。任何实现这个接口的类型都可以作为一个错误使用。

			Error这个方法提供了对错误的描述。

			当打印错误时，fmt.Println(err), fmt 包会自动调用 err.Error() 函数来打印字符串。

			这就是错误描述是如何在一行中打印出来的。

			从错误中提取更多信息的不同方法
		4.规范
			1).构造 error 的时候，要求传入的字符串首字母小写，结尾不带标点符号。因为接受error的时候，
			开发者也会增加一些日志内容，方便查看和检索
			2).error 通常是函数返回的最后一个参数

		参考：https://mp.weixin.qq.com/s/SL9ynJUgoPf0d9CWXYDCbQ
	*/
	// 初探 error
	f, err := os.Open("/test.txt")
	// 处理 error
	if err != nil {
		fmt.Println("open failed, err:", err) // 报错内容：The system cannot find the file specified
		return
	}
	fmt.Println("file is ：", f)
}
