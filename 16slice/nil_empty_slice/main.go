package main

import "fmt"

func main() {
	var slice1 []int // nil slice
	if slice1 == nil {
		fmt.Println("slice1: nil")
	}
	fmt.Printf("slice:%p\n", &slice1)

	// slice1[0] = 1
	// slice1 = append(slice1, 1)

	var slice2 = []int{} // empty slice
	if slice2 == nil {
		fmt.Println("slice2: nil")
	} else {
		fmt.Println("slice2: empty")
	}
	fmt.Printf("slice2:%p\n", &slice2)

	var slice3 = make([]int, 0)
	if slice3 == nil {
		fmt.Println("slice3: nil")
	} else {
		fmt.Println("slice3: empty")
	}
	fmt.Printf("slice3:%p\n", &slice3)
}
