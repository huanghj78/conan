package main

import (
	"strings"

	client "cpfi.client"
)

const (
	Detect     string = "Detect"
	Reproduce  string = "Reproduce"
	Evaluation string = "Evaluation"
	Leader     string = "Leader"
	Follower   string = "Follower"
	FTP        string = "FTP"
	FIP        string = "FIP"
)

var seedSeqBank []*FaultSequence
var testSeqBank []*FaultSequence

const (
	NormalScore  int = 10
	TimeoutScore int = 15
	ErrorScore   int = 13
)

// 控制不同概率
var TriggerList = []string{
	"Before Leader MsgApp",
	"Before Leader MsgApp",
	"Before Leader MsgApp",
	"Before Leader MsgAppResp",
	"Before Leader MsgAppResp",
	"After Leader MsgApp",
	"After Follower MsgAppResp",
	"Before Follower MsgApp",
	"Before Follower MsgApp",
	"Before Follower MsgApp",
	"Before Follower MsgAppResp",
	"Before Follower MsgAppResp",
}

type FaultPoint struct {
	msgType         string
	who             string
	when            string
	faultActionList []Action
}

func (fp FaultPoint) isMatch(f string) bool {
	list := strings.Split(f, " ")
	if fp.when == list[0] && fp.who == list[1] && fp.msgType == list[2] {
		return true
	}
	return false
}

func (fp FaultPoint) run(ch chan string, msg *client.CPFI_msg) {
	for _, action := range fp.faultActionList {
		go action.run(ch, msg)
	}
}

type FaultSequence struct {
	seq     []FaultPoint
	score   int
	opertor Operator
	count   int
}

func insertIntoSeedBank(sequence *FaultSequence) {
	if sequence.score > FitnessScore_Threshold {
		seedSeqBank = append(seedSeqBank, sequence)
	}
}

type FaultPointToJson struct {
	MsgType         string `json:"msgType"`
	Who             string `json:"who"`
	When            string `json:"when"`
	FaultActionList string `json:"faultActionList"` // [faultType args, faultType args]
}
