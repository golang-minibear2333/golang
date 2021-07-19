package main

import "fmt"

// 关闭后的通道还是可以正常接收
func main() {
	var chanInt chan int = make(chan int, 10)
	go func() {
		defer fmt.Println("chanInt is closed")
		defer close(chanInt)
		chanInt <- 1
	}()
	res := <-chanInt
	fmt.Println(res)
}
