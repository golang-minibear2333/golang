package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func send(c chan<- int, wg *sync.WaitGroup) {
	c <- rand.Int()
	wg.Done()
}

func received(c <-chan int, wg *sync.WaitGroup) {
	for gotData := range c {
		fmt.Println(gotData)
	}
	wg.Done()
}

func main() {
	chanInt := make(chan int, 10)
	done := make(chan struct{})
	defer close(done)
	go func() {
		var wg sync.WaitGroup
		defer close(chanInt)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go send(chanInt, &wg)
		}
		wg.Wait()
	}()
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 8; i++ {
			wg.Add(1)
			go received(chanInt, &wg)
		}
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
}
