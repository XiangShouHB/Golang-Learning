//File  : demo09_error_assertion3.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	/*
		3.直接比较

		获得更多关于错误的详细信息的第三种方法是直接与类型错误的变量进行比较。让我们通过一个例子来理解这个问题。

		filepath包的Glob函数用于返回与模式匹配的所有文件的名称。当模式出现错误时，该函数将返回一个错误ErrBadPattern。

		在filepath包中定义了ErrBadPattern，如下所述：
			var ErrBadPattern = errors.New("syntax error in pattern")

		errors.New()用于创建新的错误。

		当模式出现错误时，由Glob函数返回ErrBadPattern。
	*/

	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println("error:", error)
		return
	}
	fmt.Println("matched files:", files)

	/*
		不要忽略错误

		永远不要忽略一个错误。忽视错误会招致麻烦。

		重新编写一个示例，该示例列出了与模式匹配的所有文件的名称，而忽略了错误处理代码。
	*/

	files2, _ := filepath.Glob("[")
	fmt.Println("matched files2", files2)
	//由于我们忽略了这个错误，输出看起来好像没有文件匹配这个模式，但是实际上这个模式本身是畸形的。所以不要忽略错误

	// 使用 Contains 字符串匹配
	err := openFile("./test.txt")
	if strings.Contains(error.Error(), "not found") {
		// handle error
		fmt.Println(err)
	}
}

func openFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file, err:", err)
	}
	return nil
}
