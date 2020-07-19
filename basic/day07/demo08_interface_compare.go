//File  : demo08_interface_compare.go
//Author: duanhaobin
//Date  : 2020/5/3

package main

import (
	"errors"
	"fmt"
)

/*
什么是接口值

	接口值：即接口变量的值，由两个部分组成，一个具体的类型和那个类型的值。它们被称为接口的动态类型和动态值

	接口值的零值：动态类型type和对应的动态值value均为nil，如：var w io.Writer

	空接口值：当且仅当接口的动态类型type和对应的动态值value均为nil时，才为空接口值，此时它等于nil

	接口变量的赋值与调用过程：
	如w = os.Stdout，这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式的使用io.Writer(os.Stdout)是等价的。
	这个接口值w的动态类型被设为*os.Stdout指针的类型描述符，它的动态值持有os.Stdout的拷贝
	调用一个包含*os.File类型指针的接口值的Write方法，w.Write([]byte("hello")) ，使得(*os.File).Write方法被调用
	一个接口值可以持有任意大的动态值，不论动态值多大，接口值总是可以容下它

	接口值的可比较性：

	时刻记住：只能比较动态类型是可比较类型的接口值。
	如果接口值的动态类型是可比较的，那么它们之间就可以使用==和!=来进行比较：两个接口值相等仅当它们都是nil值或者它们的动态类型相同并且动态值也根据这个动态类型的==操作相等。
	如果接口值是可比较的，那么它们可以用在map的键或者作为switch语句的操作数

	非接口类型要么是安全的可比较类型（如基本类型和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），

	但是在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic。

	同样的风险也存在于使用接口作为map的键或者switch的操作数。

	注意：一个包含nil指针的接口不是nil接口（空接口），此时调用接口方法会发生panic错误。

	即一个接口值的动态类型type != nil，但动态值value == nil，此时的接口值 w != nil。（当把一个值为nil的非接口类型的变量转换为接口类型时，即出现这种情况）

	技巧：
	使用接口时，直接声明一个接口类型的变量，然后再对它赋值，之后使用该变量时，就可以直接把它和nil比较来判断是否为空接口
*/
func main() {
	var err1 = errors.New("1")
	var err2 = errors.New("1")
	fmt.Println(err1 == err2) // 为什么是 false   TODO

	fmt.Printf("err1 的地址为：%v\n", &err1)
	fmt.Printf("err2 的地址为：%v\n", &err2)
}

type Hello interface {
	sayHello(url string) string
}

type Greet struct {
	contents string
}

func (g *Greet)sayHello(url string)  string{
	return "Hello world"
}