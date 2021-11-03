package tmpA

import (
	"fmt"
	//_ "github.com/go100day/day11/tmpB"
	B "github.com/go100day/day11/tmpB"
)

func init() {
	fmt.Println("tool init tmp A")
}

func A() {
	fmt.Println("a")
	B.B()
}
