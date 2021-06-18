package main

import "time"

func main() {
	t := make(chan bool)
	ch := make(chan int)
	defer func() {
		close(ch)
		close(t)
	}()
	go func() {
		time.Sleep(1e9) //等待1秒
		t <- true
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 123
	}()
	select {
	case <-ch: //从ch中读取数据

	case <-t: //如果1秒后没有从ch中读取到数据，那么从t中读取，并进行下一步操作
	}
}
