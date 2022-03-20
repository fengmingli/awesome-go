package _chan

import (
	"fmt"
)

var NewSyncTaskHandler = &taskHandle{data: make(chan int)}

func NewAsyncTaskHandler(d chan int, e chan bool) *taskHandle {
	return &taskHandle{
		data: d,
		exit: e,
	}
}

type taskHandle struct {
	data chan int
	exit chan bool
}

func (h *taskHandle) SyncChan() {
	go func() {
		for d := range h.data { //通过range不断地处理data
			fmt.Println(d)
		}
	}()

	h.data <- 1
	h.data <- 2
	h.data <- 3
	close(h.data)
}

func (h *taskHandle) AsyncChanWrite() {
	for i := 0; i < 100; i++ {
		h.data <- i
	}
	close(h.data)
}

func (h *taskHandle) AsyncChanRead() {
	for d := range h.data { //如果data的缓冲区为空，这个协程会一直阻塞，除非被channel被close
		fmt.Println(d)
	}
	h.exit <- true
}
