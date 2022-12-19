package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C
			fmt.Println(time.Now())
		}
	}()
	<-time.After(5*time.Second + time.Millisecond*100)
	ticker.Stop()
}
