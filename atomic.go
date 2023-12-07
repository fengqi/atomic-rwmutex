package main

import (
	"sync/atomic"
	"time"
)

var (
	poolAtomic = atomic.Value{}
)

func AtomicWrite() {
	poolAtomic.Store(time.Now().Unix())
}

func AtomicRead() {
	_, ok := poolAtomic.Load().(int64)
	if ok {
		//
	}
}
