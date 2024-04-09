package main

import (
	"math"
	"math/rand"
	"strings"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	EnumInit               string = "Init"
	EnumInsertFTP          string = "InsertFTP"
	EnumInsertFIP          string = "InsertFIP"
	EnumDeleteFTP          string = "DeleteFTP"
	EnumDeleteFIP          string = "DeleteFIP"
	EnumModifyFTP          string = "ModifyFTP"
	EnumModifyFIP          string = "ModifyFIP"
	EnumInverseFTP         string = "InverseFTP"
	UCB_Const                     = 5
	FitnessScore_Threshold        = 5
)

var operatorList = []string{
	EnumInsertFTP,
	EnumInsertFIP,
	EnumDeleteFTP,
	EnumDeleteFIP,
	EnumModifyFTP,
	EnumModifyFIP,
	EnumInverseFTP,
}

var N int = 8

type Operator interface {
	run(seq *FaultSequence) *FaultSequence
	sample() float64
	update_TS(score float64)
	update_UCB(score float64)
	get_UCB() float64
}

type InsertFTPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *InsertFTPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *InsertFTPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *InsertFTPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *InsertFTPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

func (op *InsertFTPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	seqIndex := rand.Intn(len(seq.seq) + 1)
	var newFaultPoint FaultPoint
	index := rand.Intn(len(TriggerList))
	list := strings.Split(TriggerList[index], " ")

	actionList := []Action{}
	action := &MessageFaultAction{}
	action.setArgs()
	actionList = append(actionList, action)
	newFaultPoint = FaultPoint{
		when:            list[0],
		who:             list[1],
		msgType:         list[2],
		faultActionList: actionList,
	}
	logger.Infoln("Insert", newFaultPoint, "at", seqIndex)
	seq.seq = append(seq.seq[:seqIndex], append([]FaultPoint{newFaultPoint}, seq.seq[seqIndex:]...)...)
	return seq
}

type InsertFIPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *InsertFIPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	seqIndex := rand.Intn(len(seq.seq))
	curFaultPoint := seq.seq[seqIndex]
	// 不能插在最后(不用加一)
	fpIndex := rand.Intn(len(curFaultPoint.faultActionList))
	actionType := faultActions[rand.Intn(len(faultActions))]
	var action Action
	if actionType == EnumRestartNode {
		action = &RestartNodeAction{}
		action.setArgs()
	} else if actionType == EnumNetworkLoss {
		action = &NetworkLossAction{}
		action.setArgs()
	} else if actionType == EnumNetworkDelay {
		action = &NetworkDelayAction{}
		action.setArgs()
	} else if actionType == EnumCPUHog {
		action = &CPUHogAction{}
		action.setArgs()
	}
	seq.seq[seqIndex].faultActionList = append(seq.seq[seqIndex].faultActionList[:fpIndex], append([]Action{action}, seq.seq[seqIndex].faultActionList[fpIndex:]...)...)
	logger.Infoln("Insert FIP at", seqIndex, fpIndex)
	return seq
}

func (op *InsertFIPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *InsertFIPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *InsertFIPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *InsertFIPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

type DeleteFTPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *DeleteFTPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	if len(seq.seq) <= 1 {
		return seq
	}
	seqIndex := rand.Intn(len(seq.seq))
	seq.seq = append(seq.seq[:seqIndex], seq.seq[seqIndex+1:]...)
	logger.Infoln("Delete FTP at", seqIndex)
	return seq
}

func (op *DeleteFTPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *DeleteFTPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *DeleteFTPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *DeleteFTPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

type DeleteFIPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *DeleteFIPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	seqIndex := rand.Intn(len(seq.seq))
	if len(seq.seq[seqIndex].faultActionList) <= 1 {
		return nil
	}
	// 不能删除最后一个（要减一）
	fpIndex := rand.Intn(len(seq.seq[seqIndex].faultActionList) - 1)
	seq.seq[seqIndex].faultActionList = append(seq.seq[seqIndex].faultActionList[:fpIndex], seq.seq[seqIndex].faultActionList[fpIndex+1:]...)
	logger.Infoln("Delete FIP at", seqIndex, fpIndex)
	return seq
}

func (op *DeleteFIPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *DeleteFIPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *DeleteFIPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *DeleteFIPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

type ModifyFTPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *ModifyFTPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	seqIndex := rand.Intn(len(seq.seq))
	for _, trigger := range TriggerList {
		if !seq.seq[seqIndex].isMatch(trigger) {
			// logger.Infof("Modify FTP from %s %s %s to %s at %d\n", seq.seq[seqIndex].when, seq.seq[seqIndex].who, seq.seq[seqIndex].msgType, trigger, seqIndex)
			list := strings.Split(trigger, " ")
			if list[0] == "After" {
				if ac, ok := seq.seq[seqIndex].faultActionList[len(seq.seq[seqIndex].faultActionList)-1].(*MessageFaultAction); ok {
					if ac.isModified || ac.isOmit {
						continue
					}
				}
			}
			seq.seq[seqIndex].when = list[0]
			seq.seq[seqIndex].who = list[1]
			seq.seq[seqIndex].msgType = list[2]
			break
		}
	}
	logger.Infoln("Modify args at", seqIndex)
	return seq
}

func (op *ModifyFTPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *ModifyFTPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *ModifyFTPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *ModifyFTPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

type ModifyFIPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

// ToDo: 不要随机修改，有策略地修改
func (op *ModifyFIPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	seqIndex := rand.Intn(len(seq.seq))
	fpIndex := rand.Intn(len(seq.seq[seqIndex].faultActionList))
	curFaultPoint := seq.seq[seqIndex]
	action := curFaultPoint.faultActionList[fpIndex]

	if action.getActionType() == EnumRestartNode {
		a, _ := action.(*RestartNodeAction)
		newNum := 1
		if a.nodeNum == 1 {
			newNum = 2
		} else if a.nodeNum == 2 {
			newNum = 3
		}
		// logger.Infof("Modify restart node num from %d to %d\n", a.nodeNum, newNum)
		seq.seq[seqIndex].faultActionList[fpIndex] = &RestartNodeAction{
			nodeNum: newNum,
		}
	} else if action.getActionType() == EnumNetworkLoss {
		seq.seq[seqIndex].faultActionList[fpIndex].setArgs()
	} else if action.getActionType() == EnumNetworkDelay {
		seq.seq[seqIndex].faultActionList[fpIndex].setArgs()
	} else if action.getActionType() == EnumCPUHog {
		seq.seq[seqIndex].faultActionList[fpIndex].setArgs()
	} else if action.getActionType() == EnumMessageFault {
		seq.seq[seqIndex].faultActionList[fpIndex].setArgs()
	} else {
		return nil
	}
	logger.Infoln("Modify args at", seqIndex, fpIndex)
	return seq
}

func (op *ModifyFIPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *ModifyFIPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *ModifyFIPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *ModifyFIPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

type InverseFTPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *InverseFTPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	if len(seq.seq) <= 1 {
		return nil
	}
	// 确保不交换第一个和最后一个point
	seqIndex1 := rand.Intn(len(seq.seq))
	seqIndex2 := rand.Intn(len(seq.seq))
	for seqIndex2 == seqIndex1 { // 确保生成的两个索引不同
		seqIndex2 = rand.Intn(len(seq.seq))
	}
	seq.seq[seqIndex1], seq.seq[seqIndex2] = seq.seq[seqIndex2], seq.seq[seqIndex1]
	// logger.Infof("Inverse %d and %d\n", seqIndex1, seqIndex2)
	return seq
}

func (op *InverseFTPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *InverseFTPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *InverseFTPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *InverseFTPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}

// ///////////
type InverseFIPOp struct {
	betaDist distuv.Beta
	N        int
	Q        float64
}

func (op *InverseFIPOp) run(seedSeq *FaultSequence) *FaultSequence {
	seq := &FaultSequence{
		opertor: op,
		score:   NormalScore,
		seq:     make([]FaultPoint, 0),
		count:   0,
	}
	seq.seq = append(seq.seq, seedSeq.seq...)
	if len(seq.seq) <= 1 {
		return nil
	}
	// 确保不交换第一个和最后一个point
	seqIndex1 := rand.Intn(len(seq.seq))
	seqIndex2 := rand.Intn(len(seq.seq))
	for seqIndex2 == seqIndex1 { // 确保生成的两个索引不同
		seqIndex2 = rand.Intn(len(seq.seq))
	}
	seq1fpIndex := rand.Intn(len(seq.seq[seqIndex1].faultActionList))
	seq2fpIndex := rand.Intn(len(seq.seq[seqIndex2].faultActionList))
	ac1 := seq.seq[seqIndex1].faultActionList[seq1fpIndex]
	ac2 := seq.seq[seqIndex2].faultActionList[seq2fpIndex]
	seq.seq[seqIndex1].faultActionList[seq1fpIndex] = ac2
	seq.seq[seqIndex2].faultActionList[seq2fpIndex] = ac1
	// logger.Infof("Inverse %d and %d\n", seqIndex1, seqIndex2)
	return seq
}

func (op *InverseFIPOp) sample() float64 {
	return op.betaDist.Rand()
}

func (op *InverseFIPOp) update_TS(score float64) {
	if score > FitnessScore_Threshold {
		op.betaDist.Alpha += 1
	} else {
		op.betaDist.Beta += 1
	}
}

func (op *InverseFIPOp) update_UCB(score float64) {
	op.N += 1
	op.Q += score
}

func (op *InverseFIPOp) get_UCB() float64 {
	return op.Q/float64(op.N) + UCB_Const*math.Sqrt(math.Log(float64(N))/float64(op.N))
}
