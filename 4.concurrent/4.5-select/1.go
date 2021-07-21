package main

import (
	"fmt"
)

func main() {
	chanInt1, chanInt2, done := make(chan int), make(chan int), make(chan struct{})
	defer close(chanInt1)
	defer close(chanInt2)
	go func() {
		for {
			select {
			case <-chanInt1:
			case <-chanInt2:
			case <-done:
				fmt.Println("bye")
			}
		}
	}()
	//chanInt1 <- 1
	//chanInt2 <- 2
	done <- struct{}{}
}
