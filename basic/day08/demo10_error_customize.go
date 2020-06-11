//File  : demo10_error_customize.go
//Author: duanhaobin
//Date  : 2020/5/3

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	/*
		自定义错误:
			创建自定义错误可以使用 errors 包下的 New() 函数，以及 fmt 包下的：Errorf() 函数
			fmt包的 Errorf 函数的用武之地。这个函数根据一个格式说明器格式化错误，并返回一个字符串作为值来满足错误。
	*/
	// 1.errors.New() 创建自定义错误

	area, err := circleArea(10)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("圆的面积为；", area)

	// 2.打印出导致错误的实际半径，

	area2, err2 := circleArea2(-10)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	fmt.Println("圆的面积为；", area2)

	// 3.
}

// 计算圆的面积函数, errors.New()
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("参数错误，半径不能小于0")
	}
	return math.Pi * radius * radius, nil
}

// 计算圆的面积函数, fmt.Errorf()
func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("参数错误，半径 %.f 不能小于0", radius)
	}
	return math.Pi * radius * radius, nil
}
