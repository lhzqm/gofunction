package tmpA

import (
	"fmt"
	//_ "github.com/go100day/day11/tmpB"
	A "github.com/go100day/day11/tmpA"
)

func init() {
	fmt.Println("tool init tmp C")
}

func C() {
	fmt.Println("c")
	A.A()
}
