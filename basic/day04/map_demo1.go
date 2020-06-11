//File  : map_demo1.go
//Author: duanhaobin
//Date  : 2020/4/27

package main

import "fmt"

func main() {
	/*
		map：映射，是一种专门用于存储键值对的集合。属于引用类型

		存储特点：
			A：存储的是无序的键值对
			B：键不能重复，并且和value值一一对应的。
					map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错。

		语法结构：
			1.创建map
				var map1 map[key类型] value类型
					nil map，无法直接使用

				var map2 = make(map[key类型] value类型)

				var map3 = map[key类型] value类型{key:value,key:value,key:value...}

			2.添加/修改
				map[key]=value
					如果key不存在，就是添加数据
					如果key存在，就是修改数据

			3.获取
				map[key]-->value

				value,ok := map[key]
					根据key获取对应的value
						如果key存在，value就是对应的数据，ok为true
						如果key不存在，value就是值类型的默认值，ok为false

			4.删除数据：
				delete(map，key)
					如果key存在，就可以直接删除
					如果key不存在，删除失败
					删除函数不返回任何值
			5.长度：
				len(),返回map拥有的key的数量



		每种数据类型默认值：
			int：0
			float：0.0-->0
			string:""
			array:[00000]

			slice：nil
			map：nil

	*/
	// 1.创建map,赋值
	var map1 map[int] string   // 只有声明未初始化，等同于nil
	fmt.Println("map1 为：",map1)
	//map1[1] = "one"  // nill map不能用来存放键值对，否则报错：panic: assignment to entry in nil map
	if map1 == nil {
		map1 = make(map[int]string)  // 简单初始化
		fmt.Println("初始化后，map1 == nil：",map1 == nil)
	}

	var map2 = make(map[int] string)
	map2[1] = "one"
	map2[2] = "two"
	fmt.Println("map2 为：",map2)

	var map3 = map[string] int{"语文":90,"数学":95,"物理": 91 }
	fmt.Println("map3 为：", map3)

	map4 := map[int] string{}
	map4[0] = "zero"
	map4[1] = "one"
	fmt.Println("map4 为：",map4)

	// 2.获取数据，根据key获取对应的value值
	// 根据key获取对应的value，如果key存在，获取数值，如果key不存在，获取的是value值类型的零值
	fmt.Println("获取 map2[1] 的值：", map2[1])
	fmt.Println("获取 map3[1] 的值：", map3["test"])  // 返回int的默认值 0

	//但是当key如果不存在的时候，我们会得到该value值类型的默认值，比如string类型得到空字符串，int类型得到0。但是程序不会报错。
	//所以我们可以使用ok-idiom获取值，可知道key/value是否存在
	val, ok := map2[2]
	if ok{
		fmt.Println("map2[2] 的值为：",val)
	}else {
		fmt.Println("map2[2] 没有值")
	}
	// 3.修改数据
	fmt.Println("map2 修改数据前",map2)
	map2[3] = "three" // key 不存在，则新增
	map2[2] = "二"    // key 存在，则修改
	fmt.Println("map2 修改数据后",map2)

	// 4.删除数据
	// 如果key存在，就可以直接删除
	// 如果key不存在，删除失败
	// 删除函数不返回任何值
	delete(map2,3)
	fmt.Println("map2 删除key=3的数据后：",map2)
	delete(map2,30)
	fmt.Println("map2 删除key=30的数据后：",map2)

	// 5.长度
	fmt.Println("map3 的长度",len(map3))

}

