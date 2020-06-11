//File  : demo08_struct_nesting.go
//Author: duanhaobin
//Date  : 2020/4/29

package main

import "fmt"

type Book struct {
	bookName string
	price    float64
}
type Student struct {
	name string
	age  int
	book Book
}

// book 字段为 结构体指针
type Student2 struct {
	name string
	age  int
	book *Book // 结构体指针，Book 的地址
}

func main() {
	/*
		结构体嵌套：一个结构体可能包含一个字段，而这个字段反过来就是一个结构体。这些结构被称为嵌套结构。
			has a 的关系
	*/

	// 声明初始化 book
	book1 := Book{
		bookName: "三国演义",
		price:    56.5,
	}

	// 声明初始化 sutdent
	student1 := Student{
		name: "Tom",
		age:  20,
		book: book1,
	}
	fmt.Println("student1 结构体内容：", student1)

	fmt.Printf("访问 book 中的字段，书名:%s，价格：%.2f\n", student1.book.bookName, student1.book.price)

	// 因为结构体是值类型，所以如果对 book1 的字段进行修改，是不会影响 student1 的字段值的，包括嵌套的结构体
	book1.bookName = "水浒传"
	fmt.Println("book1 结构体的内容：", book1)       // {水浒传 56.5}  name字段值改变了
	fmt.Println("student1 结构体的内容：", student1) // {Tom 20 {三国演义 56.5}} 未改变

	// 一般情况下， 都是希望值发生改变的，这时候，嵌套的结构体得修改为 结构体指针，如 Student2 结构体的定义
	book2 := Book{bookName: "呼啸山庄", price: 76.9}
	student2 := Student2{name: "Ruby", age: 18, book: &book2}
	fmt.Println("book2 结构体的内容：", book2)
	fmt.Println("student2 结构体的内容：", student2) //{Ruby 18 0xc0000044e0}  book 字段存储的是 book2 的地址

	// 二者的地址是一致的
	fmt.Printf("book2 的地址为：%p\n", &book2)
	fmt.Printf("student2.book 的地址为：%p\n", student2.book)

	// 那么在修改 book2 的属性时，student2 嵌套的 book 字段的内容也会跟着改变
	book2.bookName = "百年孤独"
	fmt.Println("book2 结构体的内容：", book2)             // {水浒传 56.5}  name字段值改变了
	fmt.Println("student2 结构体的内容：", *student2.book) // {Tom 20 {三国演义 56.5}} 未改变
}
