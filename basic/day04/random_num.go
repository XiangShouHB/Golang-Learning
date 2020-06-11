//File  : random_num.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		生成随机数random：
			伪随机数，根据一定的算法公式算出来的。
		math/rand
	*/
	num1 := rand.Int()
	fmt.Println(num1)

	//rand.Intn(n) 生成[0,n)范围内的随机数
	for i:=1; i < 5;i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	//使每次生成的随机数都不一样，设置随机种子sedd即可
	time:= time.Now()  // 获取当前时间
	fmt.Println("time:",time)
	timeStamp := time.Unix()  // // 获取当前时间戳，单位是秒
	rand.Seed(timeStamp) // 增加随机种子

	num2 := rand.Intn(10)
	fmt.Println("随机数num2：",num2)

	/*
	获取[10,20]之间的随机数
	rand.Intn()增加一个偏移量(	20-10=10)
	 */
	num3 := rand.Intn(11) + 10
	fmt.Println("10-20之间的随机数为：",num3)




}

