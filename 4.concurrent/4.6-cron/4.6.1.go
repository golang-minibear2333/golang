package main

import (
	"fmt"
	"time"
)

func main(){
	for{
		fmt.Println(time.Now())
		time.Sleep(time.Second*1)
	}
}
