package main

import (
	"fmt"
	"sync"
)

func main(){
	var loadOnce sync.Once
	var x int
	for i:=0;i<10;i++{
		loadOnce.Do(func() {
			x++
		})
	}
	fmt.Println(x)
}
