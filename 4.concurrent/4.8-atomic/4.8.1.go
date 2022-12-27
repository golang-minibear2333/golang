package main

import (
	"sync"
	"sync/atomic"
)

var (
	xInt32  int32
	xInt64  int64
	xuInt32 uint32
	xuInt64 uint64
)

func modify(delta int32) {
	atomic.AddInt32(&xInt32, delta)
	atomic.AddInt64(&xInt64, int64(delta))
	atomic.AddUint32(&xuInt32, uint32(delta))
	atomic.AddUint64(&xuInt64, uint64(delta))
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//xInt32++
			modify(1)
		}()
	}
	wg.Wait()
	print(xInt32)
}
