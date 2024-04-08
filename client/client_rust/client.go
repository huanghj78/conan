package main

/*
#include "client.h"
*/
import "C"
import (
	"cpfi"
	"log"
	"net/rpc"
)

// go build -o libclient.so -buildmode=c-shared client.go
const (
	MsgApp     string = "MsgApp"
	MsgAppResp string = "MsgAppResp"
	MsgStart   string = "MsgStart"
)

var CPFI_msg_name = map[int32]string{
	0: MsgApp,
	1: MsgAppResp,
	2: MsgStart,
}

type CPFI_msg struct {
	Type string
	Term uint64
}

type CPFI_append_request struct {
	Msg  CPFI_msg
	When string
	Who  string
}

type CPFI_append_response struct {
	Msg CPFI_msg
	Ok  bool
}

type CPFI_request struct {
	Text string
}

type CPFI_response struct {
	Ok bool
}

const SERVER_ADDR = "127.0.0.1:8080"

var rpcClient *rpc.Client = nil

func initRPCClient() error {
	if rpcClient != nil {
		return nil
	}
	hostPort := SERVER_ADDR
	var err error
	rpcClient, err = rpc.Dial("tcp", hostPort)
	if err != nil {
		log.Printf("error in setting up connection to %s due to %v\n", hostPort, err)
		return err
	}
	return nil
}

func printRPCError(err error) {
	log.Printf("Sieve client RPC error: %v\n", err)
}

//export BeforeSendReq
func BeforeSendReq(c_msg *C.CPFIMessage) C.int {
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
func AfterSendReq(c_msg *C.CPFIMessage) C.int {
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
func BeforeSendAck(c_msg *C.CPFIMessage) C.int {
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
func AfterSendAck(c_msg *C.CPFIMessage) C.int {
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
func BeforeRecvAck(c_msg *C.CPFIMessage) C.int {
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
func BeforeRecvReq(c_msg *C.CPFIMessage) C.int {
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
