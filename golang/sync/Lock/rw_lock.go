package Lock

import (
	"sync"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/5/29 10:48 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}
