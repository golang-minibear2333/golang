package main

import (
	"fmt"
)

func multipleDeathLock() {
	fmt.Println("多值未匹配成功的死锁")
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		res := <-chanInt
		fmt.Println(res)
	}()
	chanInt <- 1
	chanInt <- 1
}

func main() {
	multipleDeathLock()
}
