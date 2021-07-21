package main

import "fmt"

func multipleDeathLock2() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		chanInt <- 1
		chanInt <- 2
	}()
	for {
		res, ok := <-chanInt
		if !ok {
			break
		}
		fmt.Println(res)
	}
}

func main() {

}
