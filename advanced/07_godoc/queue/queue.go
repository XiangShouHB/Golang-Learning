package queue

// 定义一个数据类型为int的队列
type Queue []int

// Push 往队列插入一个元素
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop 从队列里移除第一个元素
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
