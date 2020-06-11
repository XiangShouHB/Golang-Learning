//File  : slice_demo3.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		从已有的数组上，直接创建切片：
		slice := arr[startIndex:endIndex]
		 	切片中的数据：[startIndex,endIndex),
			将arr中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片（前闭后开），长度为 endIndex-startIndex

			arr[:endIndex],从头到end
			arr[startIndex:]从start到末尾

			slice 实际存储的是什么？
				slice 存储的是数组arr[startIndex:endIndex]的地址值，可以%p来查看地址值。存储了底层数组的引用
				&slice 才是slice自己在内存中开辟的地址值

		该切片的底层数组就是当前的数组。
		 	长度是从 startIndex 到 endIndex 切割的数据量。
			但是容量从 startIndex 到数组的末尾。

		切片追加元素，变化如下：
			1.追加后的切片容量小于原数组容量，那么切片的引用还是指向原数组的地址，只是切片增加了元素
			2.追加后的切片容量大于原数组容量，那么此时会重新创建一个新数组，切片的引用也指向了新数组的地址。
	          不仅增加了元素，还扩容，引用地址也发生改变
	 */
	arr := [10] int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("arr->",arr)
	fmt.Println("----------1.已有数组直接创建切片--------------------")
	s1 := arr[:5] //[1,2,3,4,5]
	s2 := arr[2:5] //[3,4,5]
	s3 := arr[5:] //[6,7,8,9,10]
	s4 := arr[:]
	//s5 := arr[:-1]  Go不支持负下标操作，Python支持
	fmt.Println("s1->",s1)
	fmt.Println("s2->",s2)
	fmt.Println("s3->",s3)
	fmt.Println("s4->",s4)

	// arr和s1,s2,s3..的内存地址是一致的，切片本身没有任何数据。它们只是对现有数组的引用
	fmt.Printf("arr 的内存地址：%p\n",&arr)
	fmt.Printf("s1 的引用地址：%p\n",s1)  // 为什么与arr内存地址一致？切片操作可以理解为对arr数组进行了分割，
	                                           //而数组的内存地址是 '首地址'，也就是第一份 "元素"的地址
	                                           //此处的元素指 这些切片
	fmt.Printf("s2 的引用地址：%p\n",s2)
	fmt.Printf("s3 的引用地址：%p\n",s3)
	fmt.Printf("s4 的引用地址：%p\n",s4)  // 与arr内存地址一致，因为本身[:]就代表原数组
	// 可以查看一下s3的内存地址
	fmt.Printf("s3 的内存地址：%p\n",&s3)



	fmt.Println("----------2.长度和容量--------------------")
	fmt.Printf("s1	len:%d,cap:%d\n",len(s1),cap(s1)) //s1	len:5,cap:10
	fmt.Printf("s2	len:%d,cap:%d\n",len(s2),cap(s2)) //s2	len:3,cap:8
	fmt.Printf("s3	len:%d,cap:%d\n",len(s3),cap(s3)) //s3	len:5,cap:5
	fmt.Printf("s4	len:%d,cap:%d\n",len(s4),cap(s4)) //s4	len:10,cap:10

	fmt.Println("----------3.更改数组的内容,切片也会发生改变--------------------")
	arr[4] = 100
	fmt.Println("arr ->",arr)  // 原数组发生改变
	fmt.Println("s1 ->",s1)   // s1发生改变
	fmt.Println("s2 ->",s2)   // s2发生改变
	fmt.Println("s3 ->",s3)   // s3未改变，因为s3取的是后五个元素 ->[5:]

	fmt.Println("----------4.更改切片的内容，原数组也会发生改变--------------------")
	s2[1] = 200
	fmt.Println("arr ->",arr)  // 原数组发生改变
	fmt.Println("s1 ->",s1)   // s1发生改变
	fmt.Println("s2 ->",s2)   // s2发生改变
	fmt.Println("s3 ->",s3)   // s3未改变，因为s3取的是后五个元素 ->[5:]

	fmt.Println("----------5.切片添加元素(不扩容)，原数组也会发生改变--------------------")
	s2 = append(s2, 1,1,1,1) // 追加了4个元素， 原来s2长度为3，容量为8，(3+4=7没有启动扩容)
	fmt.Println("arr ->",arr)  // 原数组发生改变
	fmt.Println("s1 ->",s1)   // s1未改变，因为s2截取[2:5]，追加的元素在arr的后半段范围呢，而s1去的前半段范围
	fmt.Println("s2 ->",s2)   // s2发生改变
	fmt.Println("s3 ->",s3)   // s3发生改变

	fmt.Printf("arr 的内存地址：%p\n",&arr)
	fmt.Printf("s1 的引用地址：%p\n",s1)  // s1地址未发生变化
	fmt.Printf("s2 的引用地址：%p\n",s2)  // s2地址未发生变化
	fmt.Printf("s3 的引用地址：%p\n",s3)  // s3地址未发生变化



	fmt.Println("----------5.切片添加元素(扩容)，原数组也会发生改变,切片也发生了大变化----------")
	s2 = append(s2, 2,2,3,3)  // 再追加4个元素，超过了最大容量8，启动扩容机制

	// 首先查看s2和原数组的容量
	fmt.Printf("arr	len:%d,cap:%d\n",len(arr),cap(arr)) //arr	len:10,cap:10, 原数组容量未变化
	fmt.Printf("s2	len:%d,cap:%d\n",len(s2),cap(s2)) //s2	len:11,cap:16，容量扩大了1倍，

	fmt.Println("arr ->",arr)  // 原数组并未追加 2,2,3,3 元素
	fmt.Println("s2 ->",s2)   // s2发生改变

	// s2的引用发生了变化，不再和之前引用的保持一致。
	// 因为扩容，append()函数返回了一个新的数组(容量扩大2倍),然后将引用赋值给了s2

	fmt.Printf("s2 的引用地址变化了->%p\n",s2)   // s2发生改变

	fmt.Printf("arr 的内存地址->%p\n",&arr)  // arr还是原来的数组，未变化

}

