package main

import "fmt"

func goroutineLeak() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			//不使用ok会goroutine泄漏
			res := <-chanInt
			//res,ok := <-chanInt
			//if !ok {
			//     break
			//}
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}
func goroutineLeakNoClosed() {
	chanInt := make(chan int)
	go func() {
		for {
			res := <-chanInt
			fmt.Println(res)
		}
	}()
}
func goroutineLeakNoClosed2() {
	chanInt := make(chan int)
	go func() {
		for {
			chanInt <- 1
		}
	}()
}
func main() {
	goroutineLeak()
	goroutineLeakNoClosed()
	goroutineLeakNoClosed2()
}
