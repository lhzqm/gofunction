package main

import "fmt"

type image struct {
	data map[int]int
}

func main() {
	// 1.按照顺序提供初始化值
	type person struct {
		name string
		age  int
	}
	P := person{"Tom", 25}
	fmt.Println("P:", P)
	// 2.通过field:value的方式初始化，这样可以任意顺序

	// 3.new方式,未设置初始值的，会赋予类型的默认初始值
	p := new(person)
	fmt.Printf("p:%T\n", p)
	fmt.Println("p:", p)
	fmt.Println("p.name:", p.name)
	fmt.Println("p.age:", p.age)

	p.age = 24
}
