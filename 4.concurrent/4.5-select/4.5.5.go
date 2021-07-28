package main

import (
	"fmt"
	"time"
)

func send(data []int) {
	fmt.Println(len(data))
}
func main() {
	chanInt, done := make(chan int), make(chan struct{})
	defer close(chanInt)
	defer close(done)
	go func() {
		timeout := time.Second
		t := time.NewTicker(timeout)
		defer t.Stop()
		buf := make([]int, 0, 5)
		for {
			select {
			case data := <-chanInt:
				t.Reset(timeout)
				if len(buf) < cap(buf) {
					buf = append(buf, data)
				} else {
					go send(buf)
					buf = make([]int, 0, cap(buf))
				}
			case <-t.C:
				if len(buf) > 0 {
					go send(buf)
					buf = make([]int, 0, cap(buf))
				}
			case <-done:
				return
			}
		}
	}()
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			time.Sleep(time.Second)
		}
		chanInt <- 1
	}
	done <- struct{}{}
}
