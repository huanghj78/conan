package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

type Injector struct {
	isWorking    bool
	cursor       int
	curSeq       *FaultSequence // 当前要执行的序列 从testSeqBank拿
	reproduceSeq *[]FaultPointToJson
}

func (i *Injector) start() {
	logger.Infoln("Injector start")
	if mode == Reproduce {
		return
	}
}

func (i *Injector) run(testSeq *FaultSequence) bool {
	if i.isWorking {
		logger.Warnln("Injector have been running")
		return false
	}
	logger.Infoln("Injector run, receive testSeq", translate(testSeq.seq))
	i.isWorking = true
	i.curSeq = testSeq
	i.cursor = 0
	return true
}

func (i *Injector) stop() {
	if !i.isWorking {
		logger.Warnln("Injector has been stopped")
		return
	}
	logger.Infoln("Injector stop")
	i.isWorking = false
}

func (i *Injector) startReproduce(notification *StartWorkingNotification) {
	ch := notification.getBlockingCh()
	configPath := notification.configPath
	i.isWorking = true
	i.cursor = 0
	i.reproduceSeq = &[]FaultPointToJson{}
	file, err := os.Open(configPath)
	if err != nil {
		logger.Errorln("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// 创建JSON解码器
	decoder := json.NewDecoder(file)

	// 解码JSON数据并存储到切片中
	err = decoder.Decode(i.reproduceSeq)
	if err != nil {
		logger.Errorln("Error decoding JSON:", err)
		return
	}
	ch <- Normal
	logger.Infoln("reproduce seq:", i.reproduceSeq)
	// reproduceSeq -> curSeq
	var fpSeq []FaultPoint
	for _, fp := range *i.reproduceSeq {
		point := FaultPoint{
			msgType: fp.MsgType,
			who:     fp.Who,
			when:    fp.When,
		}
		faultActionList := []Action{}
		list := strings.Split(fp.FaultActionList, ",")
		for _, ac := range list {
			acList := strings.Split(ac, " ")
			actionType := acList[0]
			if actionType == EnumCPUHog {
				percent, _ := strconv.Atoi(acList[1])
				nodeNum, _ := strconv.Atoi(acList[2])
				duration, _ := strconv.Atoi(acList[3])
				faultActionList = append(faultActionList, &CPUHogAction{
					percent:  percent,
					nodeNum:  nodeNum,
					duration: duration,
				})
			} else if actionType == EnumNetworkDelay {
				delayTime, _ := strconv.Atoi(acList[1])
				nodeNum, _ := strconv.Atoi(acList[2])
				duration, _ := strconv.Atoi(acList[3])
				faultActionList = append(faultActionList, &NetworkDelayAction{
					delayTime: delayTime,
					nodeNum:   nodeNum,
					duration:  duration,
				})
			} else if actionType == EnumNetworkLoss {
				percent, _ := strconv.Atoi(acList[1])
				nodeNum, _ := strconv.Atoi(acList[2])
				duration, _ := strconv.Atoi(acList[3])
				faultActionList = append(faultActionList, &NetworkLossAction{
					percent:  percent,
					nodeNum:  nodeNum,
					duration: duration,
				})
			} else if actionType == EnumRestartNode {
				nodeNum, _ := strconv.Atoi(acList[1])
				faultActionList = append(faultActionList, &RestartNodeAction{
					nodeNum: nodeNum,
				})
			} else if actionType == EnumMessageFault {
				isDelay := true
				isOmit := true
				isModified := true
				if acList[1] == "false" {
					isDelay = false
				}
				if acList[2] == "false" {
					isOmit = false
				}
				if acList[3] == "false" {
					isModified = false
				}
				delayTime, _ := strconv.Atoi(acList[4])
				modifyDelta, _ := strconv.Atoi(acList[5])
				faultActionList = append(faultActionList, &MessageFaultAction{
					isDelay:     isDelay,
					isOmit:      isOmit,
					isModified:  isModified,
					delayTime:   time.Duration(delayTime) * time.Millisecond,
					modifyDelta: modifyDelta,
				})
			}
		}
		point.faultActionList = faultActionList
		fpSeq = append(fpSeq, point)
	}
	i.curSeq = &FaultSequence{
		seq: fpSeq,
	}
	logger.Infoln("Reproduce seq", translate(i.curSeq.seq))

}

func (i *Injector) handleAppendEntries(notification *AppendEntriesNotification) {
	ch := notification.getBlockingCh()
	ok := Normal
	if !i.isWorking {
		ch <- ok
		return
	}
	msg := notification.msg
	if i.cursor >= len(i.curSeq.seq) {
		ch <- ok
		return
	}

	// 当前的FaultPoint
	fp := i.curSeq.seq[i.cursor]
	if !isMatch(*notification, fp) {
		logger.Infoln("notification: ", notification.when, notification.who, notification.msg.Type)
		logger.Infoln("FaultPoint: ", fp.when, fp.who, fp.msgType)
		ch <- ok
		return
	} else {
		logger.Infoln("Inject Fault!")
		fp.run(ch, msg)
		i.cursor += 1
	}
}
