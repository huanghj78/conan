package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"time"

	client "cpfi.client"
)

const (
	Normal string = "Normal"
	Omit   string = "Omit"
)

const (
	EnumNone         string = "None"
	EnumRestartNode  string = "RestartNode"
	EnumNetworkLoss  string = "EnumNetworkLoss"
	EnumNetworkDelay string = "EnumNetworkDelay"
	EnumCPUHog       string = "EnumCPUHog"
	EnumMessageFault string = "EnumMessageFault"
)

var faultActions = []string{
	EnumRestartNode,
	EnumNetworkLoss,
	EnumNetworkDelay,
	EnumCPUHog,
}

type Action interface {
	setArgs()
	run(ch chan string, msg *client.CPFI_msg)
	getActionType() string
	getActionDetail() string
	getArgs() string
}

type NoneAction struct {
}

func (a *NoneAction) setArgs() {}
func (a *NoneAction) run(ch chan string, msg *client.CPFI_msg) {
	ch <- Normal
}

func (a *NoneAction) getActionType() string {
	return EnumNone
}

func (a *NoneAction) getActionDetail() string {
	return EnumNone
}

func (a *NoneAction) getArgs() string {
	return ""
}

type MessageFaultAction struct {
	isDelay     bool
	isOmit      bool
	isModified  bool
	delayTime   time.Duration
	modifyDelta int
}

func (a *MessageFaultAction) setArgs() {
	// Five Fault Type:
	// 1. modify delay
	// 2. delay
	// 3. modify
	// 4. omit
	// 5. None
	flag := rand.Intn(5)
	if flag == 0 {
		a.isDelay = true
		a.isOmit = false
		a.isModified = true
	} else if flag == 1 {
		a.isDelay = true
		a.isOmit = false
		a.isModified = false
	} else if flag == 2 {
		a.isDelay = false
		a.isOmit = false
		a.isModified = true
	} else if flag == 3 {
		a.isDelay = false
		a.isOmit = true
		a.isModified = false
	} else if flag == 4 {
		a.isDelay = false
		a.isOmit = false
		a.isModified = false
	}
	if a.isDelay {
		a.delayTime = time.Duration(rand.Intn(2000)+100) * time.Millisecond
	}
	if a.isModified {
		a.modifyDelta = rand.Intn(10) - 5
	}
	// logger.Infoln("MessageFaultAction set args:", a.getArgs())
}

func (a *MessageFaultAction) run(ch chan string, msg *client.CPFI_msg) {
	if a.isDelay {
		time.Sleep(a.delayTime)
	}
	if a.isModified {
		if a.modifyDelta > 0 {
			msg.Term += uint64(a.modifyDelta)
		} else {
			msg.Term -= uint64(-a.modifyDelta)
		}
		logger.Infoln("New term is", msg.Term)
	}
	if a.isOmit {
		ch <- Omit
	} else {
		ch <- Normal
	}
}

func (a *MessageFaultAction) getActionType() string {
	return EnumMessageFault
}

func (a *MessageFaultAction) getActionDetail() string {
	delayStr := ""
	modifyStr := ""
	omitStr := ""
	if a.isDelay {
		delayStr = fmt.Sprintf(" Delay for %s ", a.delayTime)
	}
	if a.isModified {
		modifyStr = fmt.Sprintf(" Modify term %d ", a.modifyDelta)
	}
	if a.isOmit {
		omitStr = fmt.Sprintf("Omit")
	}
	return "EnumMessageFault " + delayStr + modifyStr + omitStr
}

func (a *MessageFaultAction) getArgs() string {
	delayStr := "false"
	modifyStr := "false"
	omitStr := "false"
	if a.isDelay {
		delayStr = "true"
	}
	if a.isModified {
		modifyStr = "true"
	}
	if a.isOmit {
		omitStr = "true"
	}
	// logger.Infoln(a.delayTime, a.modifyDelta)
	return fmt.Sprintf("%s %s %s %s %s", delayStr, omitStr, modifyStr, strconv.FormatInt(a.delayTime.Milliseconds(), 10), strconv.Itoa(a.modifyDelta))
}

type RestartNodeAction struct {
	nodeNum int
}

func (a *RestartNodeAction) setArgs() {
	a.nodeNum = rand.Intn(3) + 1
	// logger.Infoln("RestartNodeAction set args:", a.getArgs())
}

func (a *RestartNodeAction) run(ch chan string, msg *client.CPFI_msg) {
	go func() {
		cmd := fmt.Sprintf("%s/systems/%s/fault/restart_node.sh %d", projectPath, targetSystem, a.nodeNum)
		output, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			logger.Errorln("执行命令时发生错误:", err)
		}
		logger.Infoln("命令输出:\n", string(output))
		leader := findLeader()
		logger.Infof("Current Leader is Node%d\n", leader)
		logger.Infof("Restart Node%d\n", a.nodeNum)
	}()
}

func (a *RestartNodeAction) getActionType() string {
	return EnumRestartNode
}

func (a *RestartNodeAction) getActionDetail() string {
	return fmt.Sprintf("EnumRestartNode Node%d", a.nodeNum)
}

func (a *RestartNodeAction) getArgs() string {
	return strconv.Itoa(a.nodeNum)
}

type NetworkLossAction struct {
	percent  int
	nodeNum  int
	duration int //s
}

func (a *NetworkLossAction) setArgs() {
	a.percent = rand.Intn(100) + 1
	a.nodeNum = rand.Intn(3) + 1
	a.duration = rand.Intn(5) + 1
	// logger.Infoln("NetworkLossAction set args:", a.getArgs())
}

func (a *NetworkLossAction) run(ch chan string, msg *client.CPFI_msg) {
	cmd := fmt.Sprintf("%s/systems/%s/fault/network_loss.sh %d %d %d", projectPath, targetSystem, a.percent, a.nodeNum, a.duration)
	logger.Infoln(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln(string(output))
	}
}

func (a *NetworkLossAction) getActionType() string {
	return EnumNetworkLoss
}

func (a *NetworkLossAction) getActionDetail() string {
	return fmt.Sprintf("EnumNetworkLoss %d at node%d for %ds", a.percent, a.nodeNum, a.duration)
}

func (a *NetworkLossAction) getArgs() string {
	return strconv.Itoa(a.percent) + " " + strconv.Itoa(a.nodeNum) + " " + strconv.Itoa(a.duration)

}

type CPUHogAction struct {
	percent  int
	nodeNum  int
	duration int //s
}

func (a *CPUHogAction) setArgs() {
	a.percent = rand.Intn(100) + 1
	a.nodeNum = rand.Intn(3) + 1
	a.duration = rand.Intn(5) + 1
	// logger.Infoln("CPUHogAction set args:", a.getArgs())
}

func (a *CPUHogAction) run(ch chan string, msg *client.CPFI_msg) {
	cmd := fmt.Sprintf("%s/systems/%s/fault/cpu_hog.sh %d %d %d", projectPath, targetSystem, a.percent, a.nodeNum, a.duration)
	logger.Infoln(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln(string(output))
	}
}

func (a *CPUHogAction) getActionType() string {
	return EnumCPUHog
}

func (a *CPUHogAction) getActionDetail() string {
	return fmt.Sprintf("EnumCPUHog %d for %ds at node%d", a.percent, a.duration, a.nodeNum)
}

func (a *CPUHogAction) getArgs() string {
	return strconv.Itoa(a.percent) + " " + strconv.Itoa(a.nodeNum) + " " + strconv.Itoa(a.duration)
}

type NetworkDelayAction struct {
	delayTime int //ms
	nodeNum   int
	duration  int //s
}

func (a *NetworkDelayAction) setArgs() {
	a.delayTime = rand.Intn(1000) + 500
	a.nodeNum = rand.Intn(3) + 1
	a.duration = rand.Intn(5) + 1
	// logger.Infoln("NetworkDelayAction set args:", a.getArgs())
}

func (a *NetworkDelayAction) run(ch chan string, msg *client.CPFI_msg) {
	cmd := fmt.Sprintf("%s/systems/%s/fault/network_delay.sh %d %d %d", projectPath, targetSystem, a.delayTime, a.nodeNum, a.duration)
	logger.Infoln(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln(string(output))
	}
}

func (a *NetworkDelayAction) getActionType() string {
	return EnumNetworkDelay
}

func (a NetworkDelayAction) getActionDetail() string {
	return fmt.Sprintf("EnumNetworkDelay %dms for %ds at node%d", a.delayTime, a.duration, a.nodeNum)
}

func (a *NetworkDelayAction) getArgs() string {
	return strconv.Itoa(a.delayTime) + " " + strconv.Itoa(a.nodeNum) + " " + strconv.Itoa(a.duration)
}
