package tmpA

import (
	"fmt"
	//_ "github.com/go100day/day11/tmpB"
	C "github.com/go100day/day11/tmpC"
)

func init() {
	fmt.Println("tool init tmp B")
}

func B() {
	C.C()
}
