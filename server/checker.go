package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Checker struct {
	curSeq *FaultSequence
}

func (c *Checker) start() {
	logger.Infoln("Checker Start")
}

func (c *Checker) run(seq *FaultSequence) {
	// 阈值为0
	// insertSortedIntoSeedBank(seq)
	insertIntoSeedBank(seq)
	c.curSeq = seq
	logger.Infoln("Checker run")
	cmd := fmt.Sprintf("python3 %s/systems/%s/workload/check.py", projectPath, targetSystem)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln(string(output))
		logger.Errorln("=====================CHECK ERROR=====================")
		c.storeSeq()
		return
	}
	logger.Infoln("=====================CHECK PASS=====================")
}

func (c *Checker) storeSeq() {
	seq := []FaultPointToJson{}
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("无法获取当前工作目录:", err)
	}
	jsonFileName := fmt.Sprintf("%d.json", timestamp)
	jsonFilePath := filepath.Join(currentDir, "../sequences", jsonFileName)
	logger.Infoln("Store sequence to file", jsonFileName)
	for _, fp := range c.curSeq.seq {
		var list []string
		for _, ac := range fp.faultActionList {
			list = append(list, fmt.Sprintf("%s %s", ac.getActionType(), ac.getArgs()))
		}

		seq = append(seq, FaultPointToJson{
			Who:             fp.who,
			When:            fp.when,
			MsgType:         string(fp.msgType),
			FaultActionList: strings.Join(list, ","), // []
		})
	}
	file, err := os.Create(jsonFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 创建一个 JSON 编码器
	encoder := json.NewEncoder(file)
	err = encoder.Encode(seq)

	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}
}
