package main

import (
	"log"
	"strings"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	Max_Seed_Use_Count = 10
	Seed_Candidate_Num = 10
)

type Fuzzer struct {
	FIPWeight float64
	opList    []Operator
}

func (fu *Fuzzer) start() {
	rand.Seed(uint64(time.Now().UnixNano()))
	//  Reproduce mode do not need fuzzer
	if mode == Reproduce {
		return
	}
	// Init mutation Op
	fu.opList = append(fu.opList, &InsertFTPOp{
		betaDist: distuv.Beta{Alpha: 10, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &InsertFIPOp{
		betaDist: distuv.Beta{Alpha: 10, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &DeleteFTPOp{
		betaDist: distuv.Beta{Alpha: 1, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &DeleteFIPOp{
		betaDist: distuv.Beta{Alpha: 1, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &ModifyFTPOp{
		betaDist: distuv.Beta{Alpha: 1, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &ModifyFIPOp{
		betaDist: distuv.Beta{Alpha: 10, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &InverseFTPOp{
		betaDist: distuv.Beta{Alpha: 1, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        rand.Float64(),
	})
	fu.opList = append(fu.opList, &InverseFIPOp{
		betaDist: distuv.Beta{Alpha: 1, Beta: 1, Src: rand.NewSource(rand.Uint64())},
		N:        1,
		Q:        0,
	})
	// Init seed
	fu.initSeed()
}

// 初始化种子序列，种子序列为一个FTP加若干个FIP
func (fu *Fuzzer) initSeed() {
	// FTP + FIPs
	TriggerList := []string{
		"Before Leader MsgApp",
		"Before Leader MsgAppResp",
		"Before Follower MsgApp",
		"Before Follower MsgAppResp",
	}
	for _, trigger := range TriggerList {
		newSeq := []FaultPoint{}
		// 选取FTP
		list := strings.Split(trigger, " ")
		actionList := []Action{}

		// 选取FIPs, 最后一个一定要为MessageFaultAction
		numElements := rand.Intn(len(faultActions))
		for i := 0; i < numElements; i++ {
			randomIndex := rand.Intn(len(faultActions))
			actionType := faultActions[randomIndex]
			var action Action
			if actionType == EnumRestartNode {
				action = &RestartNodeAction{}
			} else if actionType == EnumNetworkLoss {
				action = &NetworkLossAction{}
			} else if actionType == EnumNetworkDelay {
				action = &NetworkDelayAction{}
			} else if actionType == EnumCPUHog {
				action = &CPUHogAction{}
			}
			action.setArgs()
			actionList = append(actionList, action)
		}
		var action Action
		action = &MessageFaultAction{}
		action.setArgs()
		actionList = append(actionList, action)
		fp := FaultPoint{
			when:            list[0],
			who:             list[1],
			msgType:         list[2],
			faultActionList: actionList,
		}
		newSeq = append(newSeq, fp)
		newFaultSequence := &FaultSequence{
			seq:   newSeq,
			score: NormalScore,
			count: 0,
		}
		logger.Infoln("Init seed seq", translate(newFaultSequence.seq))
		insertIntoSeedBank(newFaultSequence)
	}
}

func (fu *Fuzzer) selectSeed() *FaultSequence {
	randomIndex := make([]int, Seed_Candidate_Num)
	for i := 0; i < Seed_Candidate_Num; i++ {
		randomIndex[i] = rand.Intn(len(seedSeqBank))
	}
	maxScore := 0
	seq := seedSeqBank[0]
	idx := 0
	for i := 0; i < Seed_Candidate_Num; i++ {
		if seedSeqBank[randomIndex[i]].score > maxScore {
			seq = seedSeqBank[randomIndex[i]]
			maxScore = seedSeqBank[randomIndex[i]].score
			idx = randomIndex[i]
		}
	}
	seq.count += 1
	if seq.count > Max_Seed_Use_Count {
		logger.Infoln("Remove seed sequnece", seq)
		seedSeqBank = append(seedSeqBank[:idx], seedSeqBank[idx+1:]...)
	}
	return seq
}

func (fu *Fuzzer) fuzzing() *FaultSequence {
	if len(seedSeqBank) == 0 {
		log.Println("seedSeqBank is empty")
		return nil
	}

	var seedSeq *FaultSequence
	for {
		seedSeq = fu.selectSeed()
		if seedSeq != nil {
			break
		}
	}
	logger.Infoln("select seed seq", translate(seedSeq.seq))

	var seq *FaultSequence
	for {
		var op Operator
		bestSample := 0.0
		for _, item := range fu.opList {
			sample := item.sample()
			// logger.Infoln("sample", sample)
			if sample > bestSample {
				bestSample = sample
				op = item
			}
		}
		logger.Infof("Select op %T", op)
		seq = op.run(seedSeq)
		if seq != nil {
			break
		}
	}
	return checkConstraints(seq)
}
