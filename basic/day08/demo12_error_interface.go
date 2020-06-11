//File  : demo12_error_interface.go
//Author: duanhaobin
//Date  : 2020/5/3

package main

import "fmt"

func main() {
	/*
		demo11_error_customize_struct.go 中，处理错误时，主要是判断错误的类型是什么，比如 areaError 是计算面积错误
		当然，在某些情况下，这样做并不够用。
		例如，在一个网络请求中，需要调用者判断返回的错误类型，以此来决定是否重试。

		这种情况下，还有一种方法：
			不去判断错误的类型到底是什么，而是去判断错误是否具有某种行为，或者说实现了某个接口
	*/

	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)

}

/*
	当长度小于0时，lengthNegative() bool方法返回true;当宽度小于0时，widthNegative() bool方法返回true。
	这两种方法提供了更多关于误差的信息，在这种情况下，他们说面积计算是否失败，因为长度是负的，还是宽度为负的。
	因此，我们使用了struct错误类型的 方法 来提供更多关于错误的信息
*/
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError) Error() string {
	return e.err
}

// 长度不能为0
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

// 宽度不能为0
func (e *areaError) widthNegative() bool {
	return e.width < 0
}

/*
	计算面积函数
	rectArea函数检查长度或宽度是否小于0，如果它返回一个错误消息，则返回矩形的面积为nil。
*/
func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}
