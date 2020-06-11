//File  : goto_demo.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
	goto:可以无条件地转移到过程中指定的行。
	语法：
		goto label;
		..
		..
		label: statement;
	 */

	/*
		以下代码的输出结果是如下所示吗:
		i的值为：1
		i的值为：2
		i的值为：4
		i的值为：5

		执行后，会发现是
			i的值为：1
			i的值为：2
		的无限循环,为什么呢？
		关键点在for循环语句，i :=1，当执行到goto时，并不是还是接着执行之前的for循环语句，而是重新开始执行LOOP下的代码
		所以，i又被初始化为1，而循环条件为 i < 5
		因此，会出现无限输出以下内容：
			i的值为：1
			i的值为：2
	 */
	LOOP:
	for i := 1;i < 5;i++ {
		if i == 3{
			i += 1
			goto LOOP  // 以为i为3的结果不会输出
		}
		fmt.Printf("i的值为：%d\n", i)
	}



}

