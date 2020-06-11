//File  : demo07_error_assertion1.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		从错误中提取更多信息的不同方法
			1.断言底层结构类型并从 结构字段 获取更多信息
				demo05_error.go 文件中，为什么会输出那样的错误信息？查看了下源码，核心如下：
				type PathError struct {
					Op   string
					Path string
					Err  error
				}

				func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
				从上面的代码中，您可以理解 PathError 通过声明 Error() string 方法实现了错误接口。
				该方法拼接 操作、路径 和 实际错误 并返回它。这样我们就得到了错误信息，

			2.断言底层类型，并通过调用struct类型的 方法 获取更多信息。 实战见：demo08_error_assertion2.go
			3.直接比较  实战见：demo09_error_assertion3.go

	*/
	// 1.断言底层结构类型并从 结构字段 获取更多信息
	// PathError结构的路径字段包含导致错误的文件的路径。让我们修改 demo05_error.go，查看错误是如何拼接的
	f, err := os.Open("test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Printf("err.Op -> %s \n", err.Op)
		fmt.Printf("err.Path -> %s\n", err.Path)
		fmt.Printf("err.Err -> %v\n", err.Err)
		return
	}
	fmt.Println(f.Name(), "打开成功")

}
