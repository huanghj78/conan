package cpfi

import (
	"fmt"
)

func CPFI_hook(msg *CPFI_msg, when, who string) bool {
	if err := initRPCClient(); err != nil {
		fmt.Println(err)
		return true
	}
	request := &CPFI_append_request{
		Msg:  *msg,
		When: when,
		Who:  who,
	}
	var response CPFI_append_response
	if err := rpcClient.Call("Server.AppendEntries", request, &response); err != nil {
		printRPCError(err)
		return true
	}
	// 修改任期
	msg.Term = response.Msg.Term
	return response.Ok
}

func StartWorking(path string) {
	if err := initRPCClient(); err != nil {
		fmt.Println(err)
		return
	}
	request := &CPFI_request{
		Text: path,
	}
	var response CPFI_response
	if err := rpcClient.Call("Server.StartWorking", request, &response); err != nil {
		printRPCError(err)
		return
	}
	return
}
