package main

import (
	"time"
)

type Coordinator struct {
	roundTicker        *time.Ticker
	isWorking          bool
	raftNotificationCh chan Notification
	curSeq             *FaultSequence
	fuzzer             Fuzzer
	injector           Injector
	checker            Checker
	monitor            Monitor
}

func NewCoordinator() *Coordinator {
	c := &Coordinator{
		roundTicker:        time.NewTicker(time.Duration(interval) * time.Second),
		isWorking:          false,
		raftNotificationCh: make(chan Notification, 500),
		fuzzer:             Fuzzer{},
		checker:            Checker{},
		injector: Injector{
			isWorking: false,
			cursor:    0,
			curSeq:    &FaultSequence{},
		},
		monitor: Monitor{},
	}

	go c.fuzzer.start()
	go c.injector.start()
	go c.checker.start()
	go c.monitor.start()

	if mode == Detect {
		// 收集初始信息
		c.monitor.initialCollect()
	}
	go c.start()
	return c
}

// Coordinator每隔30s进行一轮测试：结束injector的工作流，
// 判断上一轮测试是否完成，从fuzzer取出一条testSeq，
// 传递给injector，并启动它的工作流
func (rc *Coordinator) start() {
	for {
		select {
		case raftNotification := <-rc.raftNotificationCh:
			go rc.processNotification(raftNotification) // 故障注入的执行所带来的开销应该不能影响到其他正常动作到执行？
		case <-rc.roundTicker.C:
			go rc.processRoundTick()
		}
	}
}

// 结束上一轮的fuzzing测试并检查结果 or 开启新一轮的fuzzing测试
func (rc *Coordinator) processRoundTick() {
	// logger.Infoln("processRoundTick")
	if mode != Detect {
		return
	}
	if rc.isWorking {
		rc.isWorking = false
		rc.injector.stop()
		rc.checker.run(rc.curSeq)
		return
	} else {
		rc.isWorking = true
		testSeq := rc.fuzzer.fuzzing()
		// testSeq := rc.fuzzer.fuzzing_random()
		rc.curSeq = testSeq
		rc.monitor.setupCluster()
		rc.injector.run(rc.curSeq)
		rc.monitor.runWorkload(rc.curSeq)
	}
}

// 接收来自client的通知
func (rc *Coordinator) processNotification(notification Notification) {
	if notification.getNotificationType() == AppendEntriesType {
		v, ok := notification.(*AppendEntriesNotification)
		if ok {
			rc.injector.handleAppendEntries(v)
		}
	} else if notification.getNotificationType() == StartWorkingType {
		v, ok := notification.(*StartWorkingNotification)
		if ok {
			rc.injector.startReproduce(v)
			logger.Infoln("Start Reproduce", v)
		}
	}
}
