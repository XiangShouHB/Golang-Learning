package queue

import "fmt"

// 队列 Push 方法示例代码
func ExampleQueue_Push() {
	var q Queue
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println("队列为:", q)

	// Output:
	// 队列为: [1 2 3 4]
}

// 队列 Pop 方法示例代码
func ExampleQueue_Pop() {
	// 测试队列
	var q Queue
	q.Push(1)
	q.Push(2)
	fmt.Println("队列为:", q)

	q.Pop() // 移除第一个元素
	fmt.Println("移除元素后，队列为:", q)

	// Output:
	// 队列为: [1 2]
	// 移除元素后，队列为: [2]
}
func ExampleQueue_IsEmpty() {
	// 测试队列
	var q Queue
	q.Push(1)
	q.Push(2)
	fmt.Println("队列为:", q)

	fmt.Println("队列是否为空：", q.IsEmpty())
	q.Pop()
	q.Pop()
	fmt.Println("移除元素后，队列是否为空:", q.IsEmpty())

	// Output:
	// 队列为: [1 2]
	// 队列是否为空： false
	// 移除元素后，队列是否为空: true
}
