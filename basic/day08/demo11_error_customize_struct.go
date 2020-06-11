//File  : demo11_error_customize_struct.go
//Author: duanhaobin
//Date  : 2020/5/3

package main

import (
	"fmt"
	"math"
)

// 第一步 创建一个struct类型来表示错误。错误类型的命名约定是，名称应该以文本Error结束。
type areaeError struct {
	err    string //错误字段存储了实际的错误消息。
	radius float64
}

// 第二步 是实现 error 接口
// 使用一个指针接收器区域错误来实现错误接口的Error() string方法。这个方法打印出半径和错误描述
func (err *areaeError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", err.radius, err.err)
}

// 重构计算面积的方法
func cicleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaeError{"半径为负数", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	/*
		使用struct 类型和字段提供关于错误的更多信息:

		将错误接口实现为错误的strut类型，这给我们提供了更多的错误处理的灵活性。

		在我们的示例中，如果我们想要访问导致错误的半径，
		那么现在唯一的方法是解析错误描述区域计算失败，半径-20.00小于零。
		这不是一种正确的方法，因为如果描述发生了变化，我们的代码就会中断。

		我们将使用在前面的教程中解释的标准库的策略，在“断言底层结构类型并从struct字段获取更多信息”，
		并使用struct字段来提供对导致错误的半径的访问。
		我们将创建一个实现错误接口的struct类型，并使用它的字段来提供关于错误的更多信息。
	*/
	radius := -10.0
	area, err := cicleArea(radius)
	if err != nil {
		if err, ok := err.(*areaeError); ok {
			fmt.Println(err)
		}
		return
	}
	fmt.Printf("半径为%.f的圆的面积为：%.2f", radius, area)

}
