package main

import "fmt"

func main() {
	// 死锁，放开注释体验
	//pendingForever()
	fmt.Println("死锁解决方法1")
	foo()
	foo2()
	fmt.Println("死锁解决方法2")
	bar()
}

func foo() {
	chanInt := make(chan int)
	go func() {
		chanInt <- 1
	}()
	res := <-chanInt
	fmt.Println(res)
}
func foo2() {
	chanInt := make(chan int)
	go func() {
		res := <-chanInt
		fmt.Println(res)
	}()
	chanInt <- 1
}

func bar() {
	chanInt := make(chan int, 1)
	chanInt <- 2
	res := <-chanInt
	fmt.Println(res)
}
