package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var m []byte

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	cpu()
	heap()
	goroutine()
	mutex()
	select {}
}
func mutex() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 1000; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
			time.Sleep(200 * time.Millisecond)
		}(i)
	}

}
func goroutine() {
	c := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() {
			r := <-c
			fmt.Println("goroutine recv:", r)
		}()
	}
	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
			time.Sleep(time.Second)
		}
	}()
}

func heap() {

	m = make([]byte, 0)
	appendMem(m)
}
func appendMem(b []byte) {
	GB := 1024 * 1024 * 1024
	go func() {
		for len(b) < GB {
			b = append(b, make([]byte, 1024*1024, 1024*1024)...)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
func cpu() {
	var workerNum int = 2
	c := make(chan int, workerNum)
	go func() { produceTask(c) }()
	for i := 0; i < workerNum; i++ {
		go func() { worker(c) }()
	}
}
func worker(c chan int) {
	for {
		<-c
		//do hard work
		fmt.Println("begin work at ", time.Now().Format("2006-01-02 15:04:05.000"))
		var j int64
		for i := 0; i < 10000000; i++ {
			j += int64(i)
		}
		fmt.Println("end work at ", time.Now().Format("2006-01-02 15:04:05.000"))

	}
}
func produceTask(c chan int) {
	for {
		fmt.Println("produce task")
		c <- 1
	}
}
