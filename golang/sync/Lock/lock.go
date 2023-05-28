package Lock

import (
	"sync"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/5/29 10:47 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}
