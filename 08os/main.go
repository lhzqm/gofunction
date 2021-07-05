package main

import (
	"fmt"
	"os"
)

func main() {
	ret := os.Getenv("NAMESPACE")
	fmt.Println("ret:", ret)
	println(os.Getenv("HOME"))
	println(os.Getenv("GOPATH"))
}
