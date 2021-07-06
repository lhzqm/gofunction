package main

import "fmt"

func main() {
	s := []int{}
	// s = append(s, 0)
	if s == nil {
		fmt.Println("空切片")
	}
	fmt.Println("s:", s)
	fmt.Printf("s:%v\n", s)
	fmt.Println("s&:", &s)

	s1 := make([]int, 3)
	if s1 == nil {
		fmt.Println("空切片")
	}
	fmt.Println("s1:", s1)
	fmt.Printf("s1:%v\n", s1)
	fmt.Println("s1&:", &s1)

	var s2 []int
	if s2 == nil {
		fmt.Println("空切片")
	}
	fmt.Println("s2:", s2)
	fmt.Printf("s2:%v\n", s2)
	fmt.Println("s2&:", &s2)

	// s2[0] = 1
	// s[1] = 2
	// s[2] = 3

	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("空切片")
	}

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	// printSlice(numbers1)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
