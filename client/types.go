package cpfi

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
