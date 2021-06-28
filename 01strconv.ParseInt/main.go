package main

import "strconv"

func main() {
	i, err := strconv.ParseInt("123", 10, 32)
	// golang strconv.ParseInt 是将字符串转换为数字的函数
	// 参数1 数字的字符串形式
	// 参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	// 参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	if err != nil {
		panic(err)
	}
	println(i)
}
