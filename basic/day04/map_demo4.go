//File  : map_demo4.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
	一：数据类型：
			基本数据类型：int，float，string，bool
			复合数据类型：array，slice，map，function，pointer，struct。。。

				array：[size]数据类型
				slice：[]数据类型
				map：map[key的类型] value的类型
			可以嵌套其他类型
	二：存储特点：
		值类型：int，float，string，bool，array，struct
		引用类型：slice，map
			由make()创建的数据类型都为引用类型，如slice，map，chan
	*/
	// 1.嵌套map
	map_student := make(map[string] map[string]string)
	map_bob := make(map[string] string)
	map_mike := make(map[string] string)
	map_bob["name"] ="Bob"
	map_bob["age"] ="18"
	map_bob["class"] ="359班"
	map_mike["name"] ="mike"
	map_mike["age"] ="19"
	map_mike["class"] ="361班"

	map_student["Bob"] = map_bob
	map_student["mike"] = map_mike

	fmt.Println("map_student ：",map_student)

	// 2. map 是引用类型
	map_bob_brother := make(map[string] string)
	map_bob_brother = map_bob

	fmt.Println("map_bob_brother ：",map_bob_brother)
	// 修改map_bob_brother，map_bob 也会改变
	map_bob_brother["age"] = "20"
	fmt.Println("map_bob_brother 改变：",map_bob_brother)
	fmt.Println("map_bob ：",map_bob)


}

