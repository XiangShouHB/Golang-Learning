//File  : switch_demo.go
//Author: duanhaobin
//Date  : 2020/4/26

package main

import "fmt"

func main() {
	/*
		switch语句：
		语法结构：
			switch 变量名{
			case 数值1：分支1
			case 数值2：分支2
			case 数值3：分支3
			......
			default：
				最后一个分支
			}

		省略switch后的变量，相当于直接作用在true上
		switch{//true
		case true:
		case false:
		}
		注意事项：
			1.switch可以作用在其他类型上，case后的数值必须和switch作用的变量类型一致
			2.case是无序的,从上直下逐一测试，直到匹配为止。
	        switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break。
			3.case后的数值是唯一的(duplicate case 3 in switch)
			4.default语句是可选的操作
			5.Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
			而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
	 */
	//定义局部变量
	score := 85
	grade := ""

	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 50,60: //case 后面可以接多个数值
		grade = "D"
	default: // 可以没有default语句
		grade = "other"
	}
	fmt.Println("根据分数评定评级：",grade)

	switch  {
	case true:
		fmt.Println("Always ture")
	case false:
		fmt.Println("Or false")
	}

}

