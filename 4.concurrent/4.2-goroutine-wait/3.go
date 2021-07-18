package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("hello world")
		}
		done <- true
	}()

	<-done
	fmt.Println("over!")
}
