//File  : demo06_error_create.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	/*

		error：是内置的数据类型，内置的接口
				定义方法：Error() string

		创建error,使用go语言提供好的包：
			1.errors 包下的函数：
				errors.New()，创建一个error对象
			2.fmt包下的 Errorf() 函数：
				func Errorf(format string, a ...interface{}) error
				可以格式化的返回错误字符串
	*/
	// 1.errors.New() 创建一个 error
	err1 := errors.New("这是 errors.New() 创建的错误")
	fmt.Printf("err1 错误类型：%T，错误为：%v\n", err1, err1)

	// 2.fmt.Errorf()
	err2 := fmt.Errorf("这个 fmt.Errorf() 创建的错误,错误编码为：%d", 404)
	fmt.Printf("err2 错误类型：%T，错误为：%v\n", err2, err2)

	// 3. go 1.13 新增加的错误处理特性  %w
	err3 := fmt.Errorf("err3: %w", err2)
	fmt.Printf("err3 错误类型：%T，错误为：%v\n", err3, err3)

	fmt.Println("-------------------------------------")
	// 3.调用 checkAge
	if err3 := checkAge(-1); err3 != nil {
		log.Fatal(err3) // 调用日志库函数，返回错误 -> 2020/05/02 16:48:41 输出的年龄有误：-1，不合法
		return
	}
	fmt.Println("main is over")
}

//定义一个返回错误的函数：验证年龄是否合法，如果为负数，就返回一个error
func checkAge(age int) error {
	if age < 0 {
		return fmt.Errorf("输出的年龄有误：%d，不合法", age)
	}
	return nil
}
