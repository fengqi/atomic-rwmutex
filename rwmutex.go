package main

import (
	"sync"
	"time"
)

var (
	pool  = make([]int64, 0)
	mutex = &sync.RWMutex{}
)

func MutexWrite() {
	mutex.Lock()
	defer mutex.Unlock()
	pool = append(pool, time.Now().Unix())
}

func MutexRead() {
	mutex.RLock()
	defer mutex.RUnlock()
	if len(pool) > 0 {
		_ = pool[len(pool)-1]
	}
}
