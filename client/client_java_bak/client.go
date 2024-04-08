package main

/*
#include "client.h"
*/
import "C"
import (
	"cpfi"
)

// GOARCH=amd64 GOOS=linux CGO_ENABLED= go build -o cpfi.so -buildmode=c-shared client.go

//export Hello
func Hello() C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgApp,
		Term: 1,
	}
	cpfi.CPFI_hook(msg, "Before", "Leader")
	return 1
}

func main() {}
