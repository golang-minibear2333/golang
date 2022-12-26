package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	var s []int
	var lock sync.Mutex

	appendValueSafe := func(i int) {
		lock.Lock()
		s = append(s, i)
		lock.Unlock()
	}

	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValueSafe(i)
	}
	time.Sleep(2)
	fmt.Println(len(s))
}
