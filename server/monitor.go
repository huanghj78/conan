package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	timeoutWeight  = 1
	statusWeight   = 2
	errorWeight    = 1
	electionWeight = 2
)

type Monitor struct {
	initialDuration   time.Duration
	durationThreshold time.Duration
	maxDelay          time.Duration
}

func (m *Monitor) start() {
	m.maxDelay = 0 * time.Second
	m.durationThreshold = 1 * time.Second
	logger.Infoln("Monitor start")
}

func (m *Monitor) initialCollect() {
	m.setupCluster()
	m.runWorkload(nil)
}

func (m *Monitor) setupCluster() {
	logger.Infof("Setup %s cluster\n", targetSystem)
	cmd := fmt.Sprintf("python3 %s/systems/%s/setup.sh", projectPath, targetSystem)
	logger.Infoln(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln(string(output))
		os.Exit(1)
	}
}

func (m *Monitor) runWorkload(seq *FaultSequence) {
	// isTimeout := 0
	isError := 0
	isElection := 0
	isStatus := 0

	cmd := fmt.Sprintf("%s/systems/%s/findLeader.sh", projectPath, targetSystem)
	logger.Infoln(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln("Find leader error")
		os.Exit(1)
	}
	leaderNum := string(output)
	logger.Infoln("Current leader num", leaderNum)

	logger.Infof("Run %s workload\n", targetSystem)
	cmd = fmt.Sprintf("python3 %s/systems/%s/workload/workload.py", projectPath, targetSystem)
	logger.Infoln(cmd)
	startTime := time.Now()
	output, err = exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		if seq != nil {
			logger.Infoln("System Error")
			// seq.score = ErrorScore
			isError = 1
		} else {
			logger.Errorln(string(output))
			os.Exit(1)
		}
	}

	cmd = fmt.Sprintf("%s/systems/%s/findLeader.sh", projectPath, targetSystem)
	output, err = exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logger.Errorln("========Find leader error======")
		// seq.opertor.update_UCB()
		isStatus = 1

	}
	newLeaderNum := string(output)
	logger.Infoln("Current leader num", newLeaderNum)
	if leaderNum != newLeaderNum {
		isElection = 1
	}

	elapsedTime := time.Since(startTime)
	if seq == nil {
		m.initialDuration = elapsedTime
	} else {
		delayTime := elapsedTime - m.initialDuration
		if delayTime > m.durationThreshold {
			// if delayTime
			logger.Infoln("System Timeout", delayTime)
			// isTimeout = 1
			// seq.opertor.update(1, 0)
		}
		timeoutValue := int(delayTime) / 1000000000
		seq.score = timeoutWeight*timeoutValue + errorWeight*isError + electionWeight*isElection + statusWeight*isStatus
		logger.Infoln("Fitness score:", seq.score)
		// seq.opertor.update_UCB(float64(seq.score))
		seq.opertor.update_TS(float64(seq.score))
	}

}
