package main

import "fmt"

// 数字求和函数 不定参数
func AddNumbers(numbers ...int) (sum int) {
	fmt.Println("numbers:", numbers)
	fmt.Println("len:",len(numbers))
	for _, v := range numbers {
		sum += v
	}

	// 具名函数可以不用显式的返回返回值变量
	return

}

func main() {
	fmt.Println(AddNumbers(1, 2, 3, 4, 5, 6))
}
