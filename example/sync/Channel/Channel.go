/**
 * @Author LFM
 * @Date 2021/10/13 11:00 下午
 * @Since V1
 */

package Channel

import (
	"fmt"
	"time"
)

var num = 100

func channelExample() {
	msg := make(chan int)
	go goRouting1(msg)
	go goRouting2(msg)
	time.Sleep(1 * time.Second)
}

func goRouting1(ch chan int) {
	for i := 1; i <= num; i++ {
		ch <- 1
		if i%2 == 1 {
			fmt.Println("goRouting1 :", i)
		}
	}
}

func goRouting2(ch chan int) {
	for i := 1; i <= num; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("goRouting2 :", i)
		}
	}
}
