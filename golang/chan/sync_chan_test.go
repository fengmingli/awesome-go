package _chan

import (
	"context"
	"fmt"
	"testing"
	"time"
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
	//NewTestPCB()

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop: // 收到了停滞信号
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func TestName2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
