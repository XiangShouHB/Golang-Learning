package main

import (
	"fmt"
	"os"
)

/*
	FileInfo接口,其定义了File信息相关的方法。
		type FileInfo interface {
			Name() string       // base name of the file 文件名.扩展名 aa.txt
			Size() int64        // 文件大小，字节数 12540
			Mode() FileMode     // 文件权限 -rw-rw-rw-
			ModTime() time.Time // 修改时间 2018-04-13 16:30:53 +0800 CST
			IsDir() bool        // 是否文件夹
			Sys() interface{}   // 基础数据源接口(can return nil)
		}
*/

func main() {
	// 获取当前文件的目录
	dir, _ := os.Getwd()
	fileinfo, err := os.Stat(dir + "/test.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	// 1.获取文件名
	fmt.Println("文件名为：", fileinfo.Name())

	// 2.获取文件大小
	fmt.Println("文件大小为：", fileinfo.Size())

	// 3.判断是否为目录
	fmt.Println("是否为目录：", fileinfo.IsDir())

	// 4.修改时间
	fmt.Println("修改时间：", fileinfo.ModTime())

	// 5.文件权限
	fmt.Println("文件权限：", fileinfo.Mode())
}
