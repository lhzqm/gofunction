package main

import (
	"fmt"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentPath(path string) string {
	return substr(path, 0, strings.LastIndex(path, "/")+1)
}
func main() {
	res := GetParentPath("-test.v=true")
	fmt.Println("res:", res)
}
