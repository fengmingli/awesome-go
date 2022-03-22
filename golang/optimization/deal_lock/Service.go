/**
 * @Author LFM
 * @Date 2021/10/14 8:09 下午
 * @Since V1
 */

package deal_lock

import (
	"sync"
	"time"
)

type SafeMap struct {
	lock sync.Mutex
	data map[string]interface{}
}

var Sm = &SafeMap{data: map[string]interface{}{}}

func worker1() {
	for {
		Sm.lock.Lock()
		defer Sm.lock.Unlock()
		Sm.data["test"] = 1
		time.Sleep(10 * time.Second)
	}
}

func worker2() {
	for {
		Sm.lock.Lock()
		defer Sm.lock.Unlock()
		Sm.data["test"] = 2
		time.Sleep(10 * time.Second)
	}
}
