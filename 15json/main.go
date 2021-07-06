package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	key := "WeeklyReportSendStatus_pre"
	value := "2"
	kb, _ := json.Marshal(key)
	vb, _ := json.Marshal(value)

	fmt.Printf("kb:%#v\n", string(kb))
	fmt.Printf("vb:%#v\n", string(vb))
}
