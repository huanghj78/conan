package main

// #cgo CFLAGS: -I$JAVA_HOME/include
// #cgo CFLAGS: -I$JAVA_HOME/include/linux
// #include <jni.h>
import "C"
import (
	"fmt"
)

// GOARCH=amd64 GOOS=linux CGO_ENABLED= go build -o HelloWorld.so -buildmode=c-shared HelloWorld.go

//export Java_CPFIClient_BeforeSendReq
func Java_CPFIClient_BeforeSendReq() C.int {
	fmt.Println("Hello")
	// msg := &cpfi.CPFI_msg{
	// 	Type: cpfi.MsgApp,
	// 	Term: 1,
	// }
	// ok := cpfi.CPFI_hook(msg, "Before", "Leader")
	// // c_msg.term = C.ulonglong(msg.Term)
	// if ok {
	// 	return 1
	// }
	return 0
}

func main() {}
