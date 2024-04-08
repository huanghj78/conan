package main

/*
#include "client.h"
*/
import "C"
import (
	"cpfi"
)

// go build -o libdcfi.so -buildmode=c-shared client.go

//export BeforeSendReq
func BeforeSendReq(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgApp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "Before", "Leader")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

//export AfterSendReq
func AfterSendReq(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgApp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "After", "Leader")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

//export BeforeSendAck
func BeforeSendAck(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgAppResp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "Before", "Follower")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

//export AfterSendAck
func AfterSendAck(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgAppResp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "After", "Follower")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

//export BeforeRecvAck
func BeforeRecvAck(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgAppResp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "Before", "Leader")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

//export BeforeRecvReq
func BeforeRecvReq(c_msg *C.CPFI_Message) C.int {
	msg := &cpfi.CPFI_msg{
		Type: cpfi.MsgApp,
		Term: uint64(c_msg.term),
	}
	ok := cpfi.CPFI_hook(msg, "Before", "Follower")
	c_msg.term = C.ulonglong(msg.Term)
	if ok {
		return 1
	}
	return 0
}

func main() {}
