//File  : demo03_type_non_local.go
//Author: duanhaobin
//Date  : 2020/5/2

package main

import (
	"time"
)

// 自定义 持续时间 类型
type myDuration = time.Duration // time 包下的类型，持续时间类型->代表两个瞬间之间经过的时间

// 为自定义类型 myDuration 添加方法
func (duration myDuration) EasySet(str string) {
	// 报错：cannot define new methods on non-local type time.Duration

}

// 修改错误：
// 将类型别名改为类型定义：
//type myDuration time.Duration

func main() {
	/*
		能够随意地为各种类型起名字，是否意味着可以在自己包里为这些类型任意添加方法？

		非本地类型不能定义方法:
			不能在一个非本地的类型 time.Duration 上定义新方法。
			非本地方法指的就是使用 time.Duration 的代码所在的包，也就是 main 包。

			因为 time.Duration 是在 time 包中定义的，在 main 包中使用。
			time.Duration 包与 main 包不在同一个包中，因此不能为不在一个包中的类型定义方法。

			解决这个问题有下面两种方法：

			1.将类型别名改为类型定义： type MyDuration time.Duration，也就是将 MyDuration 从别名改为类型。
			2.将 MyDuration 的别名定义放在 time 包中。
	*/
}
