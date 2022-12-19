package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("同Sleep的方式1 time.After")
	fmt.Println(time.Now())
	<-time.After(1 * time.Second)
	fmt.Println(time.Now())

	fmt.Println("方式2 Timer")
	done := make(chan struct{})
	timer := time.NewTimer(1 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println(time.Now())
				timer.Reset(1 * time.Second)
			case <-done:
				return
			}
		}
	}()
	<-time.After(5*time.Second + time.Millisecond*100)
	done <- struct{}{}
}
