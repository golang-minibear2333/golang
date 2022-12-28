package main

import (
	"fmt"
	"sync"
)

func unsafeMap() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()
	wg.Wait()
}

func safeMap() {
	var wg sync.WaitGroup
	var m sync.Map
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m.Store(i, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(m.Load(i))
		}
	}()
	wg.Wait()
}
func main() {
	//unsafeMap()
	safeMap()
}
