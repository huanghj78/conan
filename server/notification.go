package main

import (
	client "cpfi.client"
)

const (
	StartWorkingType  string = "StartWorkingType"
	AppendEntriesType string = "AppendEntriesType"
)

type Notification interface {
	getBlockingCh() chan string
	getNotificationType() string
}

type StartWorkingNotification struct {
	blockingCh chan string
	configPath string
}

func (n *StartWorkingNotification) getBlockingCh() chan string {
	return n.blockingCh
}

func (n *StartWorkingNotification) getNotificationType() string {
	return StartWorkingType
}

type AppendEntriesNotification struct {
	blockingCh chan string
	msg        *client.CPFI_msg
	who        string
	when       string
}

func (n *AppendEntriesNotification) getBlockingCh() chan string {
	return n.blockingCh
}

func (n *AppendEntriesNotification) getNotificationType() string {
	return AppendEntriesType
}
