package main

import (
	"fmt"
	"sync"
)

func errFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _, d := range sList {
		go func() {
			defer wg.Done()
			fmt.Println(d)
		}()
	}
	wg.Wait()
}

func correctFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _, d := range sList {
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
		}(d)
	}
	wg.Wait()
}
func main() {
	fmt.Println("error function")
	errFunc()
	fmt.Println("correct function")
	correctFunc()
}
