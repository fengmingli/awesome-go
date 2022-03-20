/**
 * @Author LFM
 * @Date 2021/10/22 3:20 下午
 * @Since V1
 */

package dead_lock

import (
	"fmt"
	"github.com/sasha-s/go-deadlock"
	"math"
	"runtime"
	"sync"
)

func DeadLock() {
	var mu deadlock.Mutex
	// Use normally, it works exactly like sync.Mutex does.
	mu.Lock()
	defer mu.Unlock()
	// Or
	var rw deadlock.RWMutex
	rw.RLock()
	defer rw.RUnlock()
}

var wg = sync.WaitGroup{}

// 模拟执行业务的 goroutine
func doBusiness(ch chan bool, i int) {
	fmt.Println("go func:", i, "goroutine count:", runtime.NumGoroutine())
	<-ch
	wg.Done()
}

func GroutingNum() {
	maxCnt := math.MaxInt64
	// max_cnt := 10
	fmt.Println(maxCnt)
	ch := make(chan bool, 3)
	for i := 0; i < maxCnt; i++ {
		wg.Add(1)
		ch <- true
		go doBusiness(ch, i)
	}
	wg.Wait()
}
