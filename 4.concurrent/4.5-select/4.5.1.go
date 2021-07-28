package main

import (
	"fmt"
	"time"
)

func baz() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		select {
		case data, ok := <-chanInt:
			if ok {
				fmt.Println(data)
			}
		}
	}()
	time.Sleep(time.Second)
	chanInt <- 1
}
func foo() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		select {
		case data, ok := <-chanInt:
			if ok {
				fmt.Println(data)
			}
		default:
			fmt.Println("全部阻塞")
		}
	}()
	time.Sleep(time.Second)
	chanInt <- 1
}
func bar() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			select {
			case data, ok := <-chanInt:
				if ok {
					fmt.Println(data)
				}
			default:
				fmt.Println("全部阻塞")
			}
		}
	}()
	chanInt <- 1
	time.Sleep(time.Second)
}
func main() {
	//foo()
	bar()
	baz()
}
