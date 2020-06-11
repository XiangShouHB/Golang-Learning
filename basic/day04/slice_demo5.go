//File  : slice_demo5.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		slice扩容机制分析：
		1.在一般情况，如果新切片需要扩容，那么新容量是原切片容量的2倍
		2.但是当原切片的长度大于或等于1024时，GO语言会以原容量的1.25倍作为新倍数来进行扩容
		3.如果一次追加的元素过多，无论原切片的长度是否大于1024，新切片的容量都会以新切片的长度为基准
		更详细参考：https://www.cnblogs.com/qcrao-2018/p/10631989.html
	 */
	// 2倍扩容机制
	fmt.Println("-----------------2倍扩容机制测试-----------------")
	a := [1023] int{}
	fmt.Printf("a的长度：%d,a的容量：%d\n",len(a),cap(a))
	b := a[:]
	fmt.Printf("b的长度：%d,b的容量：%d\n",len(b),cap(b))
	c := append(b,1)  // 原切片容量为1023(小于1024),扩容是2倍扩容
	fmt.Printf("c的长度：%d,c的容量：%d\n",len(c),cap(c))  // 容量已经变为2048

	d := append(c,b...)
	fmt.Printf("d的长度：%d,d的容量：%d\n",len(d),cap(d))  // 容量已经变为2048


	fmt.Println("-----------------1.25倍扩容机制测试，原容量已超1024-----------------")
	// 1.25扩容机制，不再以2倍增加，而是以1.25倍增加，即 2048 * 1.25 = 2560
	e := append(d, 1,1,1)  // 再加3个元素，超过容量(2048了)
	fmt.Printf("e的长度：%d,e的容量：%d\n",len(e),cap(e))  // 容量已经变为2560


	fmt.Println("-----------------超过2倍扩容机制测试-----------------")
	// 当一次追加的元素过多，以至于使新长度比原容量的2倍还要大，那么新容量会以新长度为基准
	f := [10]int{}
	g := f[:]

	test := a[:50]
	fmt.Printf("f的长度：%d,f的容量：%d\n",len(f),cap(f))
	fmt.Printf("test的长度：%d,f的容量：%d\n",len(test),cap(test))

	h := append(g,test...) // test的长度为50,远超20
	fmt.Printf("g的长度：%d,g的容量：%d\n",len(h),cap(h)) // 容量已经变为60,等于新长度


}

