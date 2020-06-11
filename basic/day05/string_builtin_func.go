//File  : string_builtin_func.go
//Author: duanhaobin
//Date  : 2020/4/28

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/*
		strings包下的关于字符串的函数
	*/
	str := "hello world,Hello Golang"

	// 1.func Compare(a, b string) int， 字符串比较，比较ASCII码值
	fmt.Println("a < b:", strings.Contains("a", "b")) // 小于，a<b,返回-1
	fmt.Println("a < a:", strings.Contains("a", "a")) // 相等，a=a,返回0
	fmt.Println("c > b:", strings.Contains("c", "b")) // 大于，c>b,返回1

	// 2.func Contains(s, substr string) bool 是否包含指定的字符串
	fmt.Println("str 包含 Go :", strings.Contains(str, "Go"))

	// 3.func ContainsAny(s, chars string) bool 是否包含chars中任意的一个字符即可
	fmt.Println("str 包含'test'字符串中的人一个字符  :", strings.Contains(str, "test"))

	// 4.func Count(s, substr string) int  统计substr在s中出现的次数
	fmt.Println("统计 h 在 str 中出现的次数 :", strings.Count(str, "h"))

	// 5.func HasPrefix(s, prefix string) bool 字符串s是否以 prefix 为前缀开头
	fmt.Println("字符串 str 是否以 he 为前缀开头 :", strings.HasPrefix(str, "he"))

	// 6.func HasSuffix(s, suffix string) bool 字符串s是否以 suffix 为后缀结尾
	fmt.Println("字符串s是否以 .txt 为后缀结尾:", strings.HasPrefix(str, ".txt"))

	// 7.func Index(s, substr string) int 查找 substr第一个实例的 在 s 中的索引，如果不存在就返回-1
	fmt.Println("查找 Go 在 str 中的位置 :", strings.Index(str, "Go")) // 返回 G 的索引

	// 8.func IndexAny(s, chars string) int  查找 chars 中任意的一个字符，出现在 s 中的位置,只要找到立刻返回
	fmt.Println("查找 Test 中任意的一个字符，首次出现在 str 中的位置:", strings.IndexAny(str, "Test"))

	// 9.func LastIndex(s, substr string) int 返回 substr 的最后一个实例的索引，如果substr 不在s中，则返回-1。
	fmt.Println("返回 substr 的最后一个实例的索引:", strings.LastIndex(str, "Test")) // Test 整体

	// 10.func LastIndexAny(s, chars string) int 返回 substr 中任意字符的最后一个实例的索引，如果都不在s中，则返回-1。
	// 如果多个字符都在 s 中，从查出的索引中选择最靠后的索引
	fmt.Println("返回 Teh 的任意一个字符的最后一个实例在 str 中的索引:", strings.LastIndexAny(str, "Teh")) // Teh 任意一个字符

	// 11.func Join(elems []string, sep string) string  拼接字符串，以 'sep' 隔开每个字符，返回字符串
	str2 := []string{"hello", "world", "hello", "golang"}
	fmt.Println("拼接字符字符串，以 - 隔开：", strings.Join(str2, "-"))

	// 12.func Split(s, sep string) []string  分割字符串，以 'sep'分割， 返回字符串切片
	fmt.Println(" xyz   以空字符串分割", strings.Split(" xyz ", ""))

	// 13.func Repeat(s string, count int) string  返回一个由字符串的计数副本组成的 新字符串。
	// 如果计数为负数或 (len(S)*count) 溢出，则会引起恐慌
	fmt.Println("重复字符串：", strings.Repeat("Hello", 4))

	// 14.func Replace(s, old, new string, n int) string 字符串替换，将旧字符(串)替换为新字符(串)，n表示替换个数，-1则全替换
	str3 := "*Hello*world*hello*golang"
	fmt.Println("替换字符串：", strings.Replace(str3, "*", "-", -1))

	// 15.func ToLower(s string) string  字符串全部转小写
	fmt.Println("字符串全部转小写：", strings.ToLower("DHB LOVE CR"))

	// 16.func ToUpper(s string) string  字符串全部转大写
	fmt.Println("字符串全部转大写：", strings.ToUpper("hello"))

	// 17.func Fields(s string) []string  以空格符 分割整个字符串 返回字符串切片
	fmt.Println("Fields are: ", strings.Fields("  foo bar  baz   "))

	// 18.func FieldsFunc(s string, f func(rune) bool) []string  分割字符串，分割方式以 func函数实现为准，返回字符串切片
	f := func(r rune) bool{
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		//unicode.IsLetter() 报告此符文是否为字母
		//unicode.IsNumber() 报告此符文是否为字母
		//rune 是int 32的别名，在所有方面都等同于int 32。按照惯例，它用于区分字符值和整数值。
	}
	fmt.Println("FieldsFunc are: ", strings.FieldsFunc("21@@#hello*@;23shel#shi**_+_2", f))  // 提取所有字母 数字

	// 19.func Trim(s string, cutset string) string  返回将 s 前后端所有 cutset 包含的 utf-8 码值都去掉的字符串。
	fmt.Println(strings.Trim("@@#Hello@, #Gophers@@#", "@#"))  // 只去除首尾部分的指定字符串

	// 20.func TrimSpace(s string) string 返回字符串的一个片段，删除由Unicode定义的所有前导和尾随空格。
	fmt.Println(strings.TrimSpace(" hel lo, wo  r l d  "))  // 只去除首尾部分的空格
	fmt.Println(strings.Replace(" hel lo, wo  r l d  "," ","",-1))  // 比较Replace去掉所有空格

	// 21. 字符串截取 切片操作
	fmt.Println(str)
	str4 := str[:5]  // 截取前5个字符
	str5 := str[5:10]  // 截取[5,10)的字符串

	fmt.Println("截取前5个字符:",str4)
	fmt.Println("截取[5,10)的字符串:",str5)
}