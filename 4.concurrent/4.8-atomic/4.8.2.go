package main

import (
	"fmt"
	"sync/atomic"
)

func foo1() {
	var xInt32 int32
	atomic.StoreInt32(&xInt32, 100)
	println(xInt32)
	v := atomic.LoadInt32(&xInt32)
	println(v)
}
func foo2() {
	var v atomic.Value
	v.Store([]int{})
	fmt.Println(v.Load().([]int))
}

func main() {
	foo1()
	foo2()
}
