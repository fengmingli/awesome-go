package _chan

import (
	"testing"
)

func TestSyncChan(t *testing.T) {
	NewSyncTaskHandler.SyncChan()
}
func TestTaskHandle_AsyncChan(t *testing.T) {
	data := make(chan int, 3)
	exit := make(chan bool)
	newChan := NewAsyncTaskHandler(data, exit)
	go newChan.AsyncChanWrite()
	go newChan.AsyncChanRead()
	<-newChan.exit //解除阻塞
}

func TestSelectChan(t *testing.T) {
	NewInput().MergeInput()
}

func TestTestQuit2Quit(t *testing.T) {
	TestQuit()
}

func TestName(t *testing.T) {
	NewTestPCB()
}
