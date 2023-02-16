package tools

import (
	"fmt"
	"runtime/debug"
)

func GoSafe(fn func()) {
	go runSafe(fn)
}

func runSafe(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			fmt.Printf("[runSafe] err:%v\n", err)
		}
	}()
	fn()
}
