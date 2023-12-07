package main

import "testing"

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go AtomicRead()
		go AtomicWrite()
	}
}
