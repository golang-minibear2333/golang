package main

import "sync/atomic"

func main() {
	var xInt32 int32
	for {
		v := atomic.LoadInt32(&xInt32)
		if atomic.CompareAndSwapInt32(&xInt32, v, v+100) {
			break
		}
	}
	print(xInt32)
}
