package main

import (
	// E:\Go\src\github.com\gofunction\mylogger
	"github.com/gofunction/10mylogger"
)

func main() {
	log := mylogger.NewLogger("Error")
	for {
		log.Debug("debug 日志")
		log.Info("info 日志")
	}
}
