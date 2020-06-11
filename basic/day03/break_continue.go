//File  : break_continue.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
	1.break:跳出当前循环体。break语句用于在结束其正常执行之前突然终止for循环
	2.continue:跳过当前循环,continue后面的语句不会被执行，直接开始执行下一次循环
	 */
	// break
	for i:=1; i < 3;i++{
		for j:=1; j < 3;j++ {
			if i*j == 4{
				fmt.Println(i,j)
				break
			}
			fmt.Println("内存循环，i:%d,j:%d",i,j)  // 当i=j=2时，内层循环执行到break时跳出循环体,这段代码不会被打印
		}
		fmt.Println("外层循环")
	}
	fmt.Println("---------------------------")
	// continue
	for i:=1; i < 3;i++{
		for j:=1; j < 3;j++ {
			if i*j == 4{
				fmt.Println(i,j)
				continue
			}
			fmt.Println("内存循环，i:%d,j:%d",i,j)  // 当i=j=2时，内层循环执行到continue时跳出本次循环,这段代码不会被打印
		}
		fmt.Println("外层循环")
	}

	/*
	 经过实践，这两段代码的数据是一致的。因为嵌套循环的存在
	 */

	for i := 1; i < 3;i++{
		fmt.Println("break 实例",i)  // break 后，不会再进入循环体
		break
	}
	fmt.Println("++++++++++++++++++-------------")

	for i := 1; i < 3;i++{
		fmt.Println("continue 实例",i)  // continue 后，跳过当前训话，执行一次循环
		continue
	}

}

