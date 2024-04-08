package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	Tolerant_Num = 1
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fail due to error: %v\n", err)
	}
}

func translate(seq []FaultPoint) string {
	res := ""
	for _, fp := range seq {
		tmp := ""
		tmp += fmt.Sprintf("[FTP %s %s %s]", fp.when, fp.who, fp.msgType)
		for _, ac := range fp.faultActionList {
			detail := ac.getActionDetail()
			tmp += fmt.Sprintf("[FIP %s]", detail)
		}
		res += tmp
	}
	return res
}

func findLeader() int {
	cmd := fmt.Sprintf("%s/systems/%s/findLeader.sh", projectPath, targetSystem)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Find leader error", err)
	}
	leaderNum := string(output)
	num, _ := strconv.Atoi(leaderNum)
	return num
}

func isMatch(no AppendEntriesNotification, fp FaultPoint) bool {
	if no.who == fp.who && no.when == fp.when && no.msg.Type == fp.msgType {
		return true
	}
	return false
}

func checkConstraints(seq *FaultSequence) *FaultSequence {
	for _, fp := range seq.seq {
		for idx, ac := range fp.faultActionList {
			if ac.getActionType() == EnumMessageFault {
				if idx != len(fp.faultActionList)-1 {
					fp.faultActionList = append(fp.faultActionList[:idx], fp.faultActionList[idx+1:]...)
				}
			}
		}
		ac := fp.faultActionList[len(fp.faultActionList)-1]
		if ac.getActionType() != EnumMessageFault {
			var action Action
			action = &MessageFaultAction{}
			action.setArgs()
			fp.faultActionList = append(fp.faultActionList, action)
		}
		// 2
		if ac, ok := fp.faultActionList[len(fp.faultActionList)-1].(*MessageFaultAction); ok {
			args := strings.Split(ac.getArgs(), " ")
			if args[1] == "true" && args[2] == "true" {
				ac.setArgs()
			}
		}

		// 3
		if fp.isMatch("After Leader Send") || fp.isMatch("After Leader Receive") || fp.isMatch("After Follower Send") || fp.isMatch("After Follower Receive") {
			if ac, ok := fp.faultActionList[len(fp.faultActionList)-1].(*MessageFaultAction); ok {
				for {
					args := strings.Split(ac.getArgs(), " ")
					if args[1] == "false" && args[2] == "false" {
						break
					}
					ac.setArgs()
				}
			}

		}
		// 1, 4
		lastActionType := EnumNone
		for idx, ac := range fp.faultActionList {
			curActionType := ac.getActionType()
			if curActionType == lastActionType {
				if curActionType == EnumRestartNode {
					fp.faultActionList = append(fp.faultActionList[:idx], fp.faultActionList[idx+1:]...)
				} else if curActionType == EnumNetworkDelay {
					Lastid := ""
					Curid := ""
					if LastAc, ok := fp.faultActionList[idx-1].(*NetworkDelayAction); ok {
						Lastid = strings.Split(LastAc.getArgs(), " ")[1]
					}
					if CurAc, ok := fp.faultActionList[idx-1].(*NetworkDelayAction); ok {
						Curid = strings.Split(CurAc.getArgs(), " ")[1]
					}
					if Lastid == Curid {
						fp.faultActionList = append(fp.faultActionList[:idx], fp.faultActionList[idx+1:]...)
					}
				} else if curActionType == EnumNetworkLoss {
					Lastid := ""
					Curid := ""
					if LastAc, ok := fp.faultActionList[idx-1].(*NetworkLossAction); ok {
						Lastid = strings.Split(LastAc.getArgs(), " ")[1]
					}
					if CurAc, ok := fp.faultActionList[idx-1].(*NetworkLossAction); ok {
						Curid = strings.Split(CurAc.getArgs(), " ")[1]
					}
					if Lastid == Curid {
						fp.faultActionList = append(fp.faultActionList[:idx], fp.faultActionList[idx+1:]...)
					}
				} else if curActionType == EnumCPUHog {
					Lastid := ""
					Curid := ""
					if LastAc, ok := fp.faultActionList[idx-1].(*CPUHogAction); ok {
						Lastid = strings.Split(LastAc.getArgs(), " ")[1]
					}
					if CurAc, ok := fp.faultActionList[idx-1].(*CPUHogAction); ok {
						Curid = strings.Split(CurAc.getArgs(), " ")[1]
					}
					if Lastid == Curid {
						fp.faultActionList = append(fp.faultActionList[:idx], fp.faultActionList[idx+1:]...)
					}
				}
			}
			lastActionType = curActionType
		}

	}
	return seq
}
