package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("ok")

	// var num = int(1)
	// time.Sleep(num * time.Second)

	var num = int(1)
	// time.Duration(num) 做类型转换
	time.Sleep(time.Duration(num) * time.Second)
}
