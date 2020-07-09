package main

import (
	"fmt"
	"io"
	"os"
)

// example1 defer函数的参数传递为值拷贝
func example1() int {
	a := 0
	defer func(i int) {
		fmt.Println("defer i = ", i)
	}(a)
	// 不会影响 i 的输出
	a++
	return a
}

// example2 defer语句必须先注册后执行
func example2() {
	defer func() {
		fmt.Println("First.....")
	}()
	a := 0
	fmt.Println("a----", a)
	return
	defer func() {
		fmt.Println("Second.....")
	}()
}

// example3 主动调用os.Exit(int)退出进程时，defer将不再执行
func example3() {
	defer func() {
		fmt.Println("defer不会被执行.....")
	}()
	fmt.Println("正常执行......")
	os.Exit(1)
}

// example4 多个需要关闭资源的场景，可以避免资源泄漏
func example4(dst, src string) (w int64, err error) {
	fileSrc, err := os.Open(src)
	if err != nil {
		return
	}
	fileDst, err := os.Open(dst)
	if err != nil {
		// fileSrc很容易忘记关闭
		// fileSrc.Close()
		return
	}
	// 直接使用defer很方便，降低心智负担
	defer fileSrc.Close()
	w, err = io.Copy(fileDst, fileSrc)
	fileDst.Close()
	// fileSrc.Close()
	return
}
func main() {
	fmt.Println("defer函数的参数传递为值拷贝结果为 :", example1())
	fmt.Println("defer语句必须先注册后执行 :")
	example2()
	fmt.Println("主动调用os.Exit(int)退出进程时，defer将不再执行 :")
	example3()
}
