package main

import "fmt"

func multipleLoop() {
	fmt.Println("解决多值发送死锁")
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			res, ok := <-chanInt
			if !ok {
				break
			}
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}

func main() {
	multipleLoop()
}
