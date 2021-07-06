package main

import (
	"fmt"
	"strings"
)

func main() {
	// func LastIndex(str, substring string) int
	// Golang中的strings.LastIndex()
	// 函数返回给定字符串中子字符串的最后一个实例出现的起始索引。
	// 如果未找到子字符串，则返回-1。因此，此函数返回一个整数值。
	// 索引以零作为字符串的起始索引进行计数。

	// taking a string
	str := "GeeksforGeeks"
	substr := "Geeks"
	fmt.Println(strings.LastIndex(str, substr))

	// taking strings
	str2 := "My favorite sport is football"
	substr1 := "f"
	substr2 := "ll"
	substr3 := "SPORT"

	// using the function
	fmt.Println(strings.LastIndex(str2, substr1))
	fmt.Println(strings.LastIndex(str2, substr2))
	fmt.Println(strings.LastIndex(str2, substr3))
}
