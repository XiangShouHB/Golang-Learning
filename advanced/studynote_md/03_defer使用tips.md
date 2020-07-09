# defer学习及使用
**Go**提供了`defer`关键字，可以注册多个延迟调用，这些调用以先进后出(FILO)的顺序在函数返回前被执行。

常用于保证一些资源最终一定能够得到回收和释放

## 一、defer后面必须是函数或方法的调用，不能是语句

必须满足该条件，否则会报`expression in defer must be function call`错误

还有个小细节就是：`defer`后面的函数是匿名函数时，一定要记得在**结尾加**`()`，表示调用该匿名函数，有参数的记得传参。否则也会报以上错误

## 二、defer函数的参数传递为值拷贝

`defer`函数的实参在注册时通过值拷贝传递进去，下面的实例代码中，实参a的值在 `defer` 注册时通过值拷贝传递进去，
后续语句 `a++`并不会影响`defer`语句最后的输出结果：
```go
package main

import (
	"fmt"
)
// example1 defer函数的参数传递为值拷贝
func example() int {
	a := 0
	defer func(i int) {
		fmt.Println("defer i = ", i)
	}(a)
	// 不会影响 i 的输出
	a++
	return a
}

func main() {
	fmt.Println("结果为：", example())

}

输出：
defer i =  0
结果为： 1
```

## 三、defer语句必须先注册后执行
如果`defer`语句位于`return`之后，则`defer`因为没有注册，不会执行

```go
// example2 defer语句必须先注册后执行
func example2() {
	defer func() {
		fmt.Println("First.....")
	}()
	a := 0
	fmt.Println("a----", a)
	return
	defer func() {
        // 不会执行
		fmt.Println("Second.....")
	}()
}
func main() {
	fmt.Println("defer语句必须先注册后执行 :")
	example2()
}

输出：
defer语句必须先注册后执行 :
a---- 0
First.....
```

## 四、主动调用`os.Exit(int)`退出进程时，defer将不再执行
即使`defer`已经提前注册了，仍不执行。
```go
// example3 主动调用os.Exit(int)退出进程时，defer将不再执行
func example3() {
	defer func() {
		fmt.Println("defer不会被执行.....")
	}()
	fmt.Println("正常执行......")
	os.Exit(1)
}


func main() {
	fmt.Println("defer语句必须先注册后执行 :")
	example2()
}

输出：
主动调用os.Exit(int)退出进程时，defer将不再执行 :
正常执行......
exit status 1
```

## 五、多个需要关闭资源的场景，可以避免资源泄漏
`defer` 的好处是可以在一定程度上避免资源浪费，特别是在有很多`return`语句，有多个资源需要关闭的场景中，很容易漏掉资源的关闭操作，例如：
```go
func example4(dst, src string) (w int64, err error) {
	fileSrc, err := os.Open(src)
	if err != nil {
		return
	}
	fileDst, err := os.Open(dst)
	if err != nil {
		// fileSrc很容易忘记关闭
		fileSrc.Close()
		return
	}
	w, err = io.Copy(fileDst, fileSrc)
	fileDst.Close()
	fileSrc.Close()
	return
}

```
使用defer改写后，在打开资源的无报错好直接调用defer关闭资源
```go
func example4(dst, src string) (w int64, err error) {
	fileSrc, err := os.Open(src)
	if err != nil {
		return
	}
	fileDst, err := os.Open(dst)
	if err != nil {
		// fileSrc.Close()
		return
	}
	// 直接使用defer很方便，降低心智负担
	defer fileSrc.Close()
	w, err = io.Copy(fileDst, fileSrc)
	fileDst.Close()
	// fileSrc.Close()
	return
}
```

不过也要注意，`defer`语句的位置不当，有可能导致`panic`，一般`defer`语句放在错误检查语句之后。


## 六、副作用
`defer`也有明显的副作用：`defer`会推迟资源的释放，`defer`尽量不要放到循环语句里，将大函数内部的 `defer`语句单独拆分成一个小函数是一种很好的实践方式。

另外，`defer`相对于普通的函数调用需要间接的数据结构支持，相对于普通函数调用有一定的性能损耗

## 七、不要对有名返回值参数进行操作
`defer`中最要不要对有名返回值参数进行操作，否则会引发意想不到的结果

## 八、defer的执行顺序考察
```go
package main

import "fmt"

func GetFn() func() {
	fmt.Println("outside")
	return func() {
		fmt.Println("inside")
	}
}
func main() {

	defer GetFn()()
	fmt.Println("main...")
}
```
错误输出结果:
```
main...
outside
inside
```
正确输出结果：
```
outside
main...
inside
```
解析：

`GetFn()` 整体可以看做是一个变量名，那么 `GetFn()()`可以表示为 `F()`。

真正 defer 的是**F()函数**，对于F来讲，它只是一个字面量，也就是说，`GetFn()`还是正常执行，没有被defer。

因此在 main 函数里， 先执行`GetFn()`  然后`F()`被defer，所以正常执行main函数，打印 main...

此时 main 函数没有其他代码可执行了，因此就执行 defer 的代码 `F()`,就输出了 inside