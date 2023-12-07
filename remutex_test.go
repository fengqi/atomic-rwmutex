package main

import "testing"

func BenchmarkMutex(t *testing.B) {
	for i := 0; i < t.N; i++ {
		go MutexRead()
		go MutexWrite()
	}
}
