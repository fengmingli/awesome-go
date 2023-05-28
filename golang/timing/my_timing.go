package timing

import (
	"context"
	"fmt"
	"time"
)

func MyTiming() {
	t5 := time.NewTimer(5 * time.Second)
	defer t5.Stop()

	for {
		<-t5.C
		fmt.Println("timer running........")
		t5.Reset(5 * time.Second)
	}
}

func MyTiming2() {
	t := time.NewTimer(time.Second * 2)
	ch := make(chan bool)
	go MyTime2(t, ch)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}

func MyTime2(t *time.Timer, ch chan bool) {
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("timer running....")
			// 需要重置Reset 使 t 重新开始计时
			t.Reset(time.Second * 2)
		case stop := <-ch:
			if stop {
				fmt.Println("timer Stop")
				return
			}
		}
	}
}

func MyTick() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	metricsCh := make(chan int64)
	volumeCh := make(chan int64)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done:")
				return
			case met := <-metricsCh:
				fmt.Println("metricsCh:", met)
			case vc := <-volumeCh:
				fmt.Println("volumeCh:", vc)
			}
		}
	}()

	ticker := time.Tick(10 * time.Second)
	count := 0
	count1 := 0
	for range ticker {
		count = count + 1
		count1 = count1 + 2
		metricsCh <- int64(count)
		volumeCh <- int64(count1)
		if count1 == 10 {
			return
		}
	}
}
