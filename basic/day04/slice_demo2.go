//File  : slice_demo2.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		切片Slice：
			1.每一个切片引用了一个底层数组
			2.切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
			3.当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长 乘2)
			4.切片一旦扩容，就是重新指向一个新的底层数组

		s1:3--->6--->12--->24

		s2:4--->8--->16--->32....

	*/
	s1 := [] int{1,2,3}
	fmt.Printf("原切片s1的长度：%d,容量：%d\n",len(s1),cap(s1))
	fmt.Printf("%p\n", s1)  // 查看s1的地址

	s1 = append(s1,4,5)  // 容量扩容到6
	fmt.Printf("追加4,5,6元素后，s1的长度：%d,容量：%d\n",len(s1),cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 6,7,8)  // 容量扩容到12
	fmt.Printf("追加6,7,8元素后，s1的长度：%d,容量：%d\n",len(s1),cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1,s1...)  // 容量扩容到24
	fmt.Printf("追加6,7,8元素后，s1的长度：%d,容量：%d\n",len(s1),cap(s1))
	fmt.Printf("%p\n", s1)
	fmt.Println("s1:", s1)


}

