package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwlock sync.RWMutex
	lock   sync.Mutex
	wg     sync.WaitGroup
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	//lock.Unlock()
}
func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	//lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
