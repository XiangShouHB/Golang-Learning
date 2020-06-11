//File  : practice01_memory_leak.go
//Author: duanhaobin
//Date  : 2020/5/1

package main

import "fmt"

var a []int

func main() {
	/*
		参考：https://mp.weixin.qq.com/s/pzJNcMc-P6X-xdC1N0r_mg
				https://blog.csdn.net/sunxianghuang/article/details/93869683
				https://segmentfault.com/a/1190000016412013
		Golang虽然是自带GC的语言，仍然存在内存泄漏的情况。
		记录其中的Slice内存泄漏的情况做分析，并介绍Slice实现和使用的一些关键逻辑。
	*/
	// 1.查看内存泄漏
	arr := []int{1, 1, 2, 3, 4, 5, 6, 7, 89}
	testMemoryLeak(arr)

}

// Golang是自带GC的，如果资源一直被占用，是不会被自动释放的，比如下面的代码，
// 如果传入的slice b是很大的，然后引用很小部分给全局变量a，那么b未被引用的部分就不会被释放，造成了所谓的内存泄漏。
// 新,旧slice指向的都是同一片内存地址，那么只要全局变量a在，b就不会被回收
func testMemoryLeak(b []int) {
	a = b[:1] // a = b[:1]是一个引用
	fmt.Println(a)
	return
}
