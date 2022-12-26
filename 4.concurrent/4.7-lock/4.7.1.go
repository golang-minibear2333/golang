package main

import (
	"fmt"
	"time"
)

var s []int

func appendValue(i int) {
	s = append(s, i)
}

func main() {
	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValue(i)
	}
	time.Sleep(2)
	fmt.Println(len(s))
}
