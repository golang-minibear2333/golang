package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say2("hello", &wg)
	go say2("world", &wg)
	fmt.Println("over!")
	wg.Wait()
}

func say2(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
