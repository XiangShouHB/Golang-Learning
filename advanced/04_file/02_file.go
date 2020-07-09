package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

/*
	文件操作：
	1.路径：
		相对路径：relative
			相对于当前工程

		绝对路径：absolute


		.当前目录
		..上一层目录
	2.创建文件夹，如果文件夹存在，创建失败
		os.MkDir()，创建一层
		os.MkDirAll()，可以创建多层

	3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
		os.Create()，创建文件

	4.打开文件：让当前的程序，和指定的文件之间建立一个连接
		os.Open(filename)
		os.OpenFile(filename,mode,perm)

	5.关闭文件：程序和文件之间的链接断开。
		file.Close()

	5.删除文件或目录：慎用，慎用，再慎用
		os.Remove()，删除文件和空目录
		os.RemoveAll()，删除所有
*/
func main() {
	// 1.路径
	fileName1 := "test.txt"
	fileName2 := "/home/duanhaobin/WorkSpace/go/src/imooc.com/Golang-Learning/advanced/04_file/test.txt"
	fmt.Println("是否为相对路径：", filepath.IsAbs(fileName1)) // false
	fmt.Println("是否为相对路径：", filepath.IsAbs(fileName2)) // true
	// 根据路径名/文件名，获取绝对路径
	fmt.Println(filepath.Abs(fileName1))
	// 获取父目录
	fmt.Println("获取父目录：", path.Join(fileName2, ".."))

	// 打开文件
	file, err := os.OpenFile(fileName2, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Println("打开的文件为：", file)

}
